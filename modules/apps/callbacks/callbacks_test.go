package ibccallbacks_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/stretchr/testify/suite"

	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/ibc-go/v7/modules/apps/callbacks/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v7/testing"
)

// CallbacksTestSuite defines the needed instances and methods to test callbacks
type CallbacksTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain

	path *ibctesting.Path
}

// setupChains sets up a coordinator with 2 test chains.
func (suite *CallbacksTestSuite) setupChains() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(2))
}

// SetupTransferTest sets up a transfer channel between chainA and chainB
func (suite *CallbacksTestSuite) SetupTransferTest() {
	suite.setupChains()

	suite.path = ibctesting.NewPath(suite.chainA, suite.chainB)
	suite.path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	suite.path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	suite.path.EndpointA.ChannelConfig.Version = transfertypes.Version
	suite.path.EndpointB.ChannelConfig.Version = transfertypes.Version

	suite.coordinator.Setup(suite.path)
}

// SetupICATest sets up an interchain accounts channel between chainA (controller) and chainB (host).
// It funds and returns the interchain account address owned by chainA's SenderAccount.
func (suite *CallbacksTestSuite) SetupICATest() string {
	suite.setupChains()

	suite.path = ibctesting.NewPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(suite.path)

	icaOwner := suite.chainA.SenderAccount.GetAddress().String()
	// ICAVersion defines a interchain accounts version string
	ICAVersion := icatypes.NewDefaultMetadataString(suite.path.EndpointA.ConnectionID, suite.path.EndpointB.ConnectionID)
	ICAControllerPortID, err := icatypes.NewControllerPortID(icaOwner)
	suite.Require().NoError(err)

	suite.path.SetChannelOrdered()
	suite.path.EndpointA.ChannelConfig.PortID = ICAControllerPortID
	suite.path.EndpointB.ChannelConfig.PortID = icatypes.HostPortID
	suite.path.EndpointA.ChannelConfig.Version = ICAVersion
	suite.path.EndpointB.ChannelConfig.Version = ICAVersion

	suite.RegisterInterchainAccount(icaOwner)
	// open chan init must be skipped. So we cannot use .CreateChannels()
	suite.path.EndpointB.ChanOpenTry()
	suite.path.EndpointA.ChanOpenAck()
	suite.path.EndpointB.ChanOpenConfirm()

	interchainAccountAddr, found := suite.chainB.GetSimApp().ICAHostKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), suite.path.EndpointA.ConnectionID, suite.path.EndpointA.ChannelConfig.PortID)
	suite.Require().True(found)

	// fund the interchain account on chainB
	msgBankSend := &banktypes.MsgSend{
		FromAddress: suite.chainB.SenderAccount.GetAddress().String(),
		ToAddress:   interchainAccountAddr,
		Amount:      sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(100000))),
	}
	res, err := suite.chainB.SendMsgs(msgBankSend)
	suite.Require().NotEmpty(res)
	suite.Require().NoError(err)

	return interchainAccountAddr
}

// RegisterInterchainAccount invokes the the InterchainAccounts entrypoint, routes a new MsgChannelOpenInit to the appropriate handler,
// commits state changes and updates the testing endpoint accordingly on chainA.
func (suite *CallbacksTestSuite) RegisterInterchainAccount(owner string) {
	portID, err := icatypes.NewControllerPortID(owner)
	suite.Require().NoError(err)

	channelSequence := suite.chainA.GetSimApp().GetIBCKeeper().ChannelKeeper.GetNextChannelSequence(suite.chainA.GetContext())

	err = suite.chainA.GetSimApp().ICAControllerKeeper.RegisterInterchainAccount(suite.chainA.GetContext(), suite.path.EndpointA.ConnectionID, owner, suite.path.EndpointA.ChannelConfig.Version)
	suite.Require().NoError(err)

	// commit state changes for proof verification
	suite.chainA.NextBlock()

	// update port/channel ids
	suite.path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(channelSequence)
	suite.path.EndpointA.ChannelConfig.PortID = portID
}

// AssertHasExecutedExpectedCallback checks if the only the expected type of callback has been executed.
// It assumes that the source chain is chainA and the destination chain is chainB.
//
// The callbackType can be one of the following:
//   - types.CallbackTypeAcknowledgement
//   - types.CallbackTypeReceivePacket
//   - types.CallbackTypeTimeout
//   - "none" (no callback should be executed)
func (suite *CallbacksTestSuite) AssertHasExecutedExpectedCallback(callbackType string, isSuccessful bool) {
	successCount := uint64(0)
	if isSuccessful {
		successCount = 1
	}
	switch callbackType {
	case types.CallbackTypeAcknowledgement:
		suite.Require().Equal(successCount, suite.chainA.GetSimApp().MockKeeper.AckCallbackCounter.Success)
		suite.Require().Equal(1-successCount, suite.chainA.GetSimApp().MockKeeper.AckCallbackCounter.Failure)
		suite.Require().True(suite.chainA.GetSimApp().MockKeeper.TimeoutCallbackCounter.IsZero())
		suite.Require().True(suite.chainB.GetSimApp().MockKeeper.RecvPacketCallbackCounter.IsZero())
	case types.CallbackTypeReceivePacket:
		suite.Require().Equal(successCount, suite.chainB.GetSimApp().MockKeeper.RecvPacketCallbackCounter.Success)
		suite.Require().Equal(1-successCount, suite.chainB.GetSimApp().MockKeeper.RecvPacketCallbackCounter.Failure)
		suite.Require().True(suite.chainA.GetSimApp().MockKeeper.TimeoutCallbackCounter.IsZero())
		suite.Require().True(suite.chainB.GetSimApp().MockKeeper.AckCallbackCounter.IsZero())
	case types.CallbackTypeTimeoutPacket:
		suite.Require().Equal(successCount, suite.chainA.GetSimApp().MockKeeper.TimeoutCallbackCounter.Success)
		suite.Require().Equal(1-successCount, suite.chainA.GetSimApp().MockKeeper.TimeoutCallbackCounter.Failure)
		suite.Require().True(suite.chainA.GetSimApp().MockKeeper.AckCallbackCounter.IsZero())
		suite.Require().True(suite.chainB.GetSimApp().MockKeeper.RecvPacketCallbackCounter.IsZero())
	case "none":
		suite.Require().True(suite.chainA.GetSimApp().MockKeeper.AckCallbackCounter.IsZero())
		suite.Require().True(suite.chainA.GetSimApp().MockKeeper.TimeoutCallbackCounter.IsZero())
		suite.Require().True(suite.chainB.GetSimApp().MockKeeper.RecvPacketCallbackCounter.IsZero())
	default:
		suite.FailNow("invalid callback type")
	}
	suite.Require().True(suite.chainB.GetSimApp().MockKeeper.AckCallbackCounter.IsZero())
	suite.Require().True(suite.chainB.GetSimApp().MockKeeper.TimeoutCallbackCounter.IsZero())
	suite.Require().True(suite.chainA.GetSimApp().MockKeeper.RecvPacketCallbackCounter.IsZero())
}

func TestIBCCallbacksTestSuite(t *testing.T) {
	suite.Run(t, new(CallbacksTestSuite))
}