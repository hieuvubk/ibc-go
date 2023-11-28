"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[9569],{3905:(e,n,t)=>{t.d(n,{Zo:()=>p,kt:()=>d});var a=t(67294);function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function r(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);n&&(a=a.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,a)}return t}function o(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?r(Object(t),!0).forEach((function(n){i(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):r(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function s(e,n){if(null==e)return{};var t,a,i=function(e,n){if(null==e)return{};var t,a,i={},r=Object.keys(e);for(a=0;a<r.length;a++)t=r[a],n.indexOf(t)>=0||(i[t]=e[t]);return i}(e,n);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(a=0;a<r.length;a++)t=r[a],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var c=a.createContext({}),l=function(e){var n=a.useContext(c),t=n;return e&&(t="function"==typeof e?e(n):o(o({},n),e)),t},p=function(e){var n=l(e.components);return a.createElement(c.Provider,{value:n},e.children)},h="mdxType",u={inlineCode:"code",wrapper:function(e){var n=e.children;return a.createElement(a.Fragment,{},n)}},m=a.forwardRef((function(e,n){var t=e.components,i=e.mdxType,r=e.originalType,c=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),h=l(t),m=i,d=h["".concat(c,".").concat(m)]||h[m]||u[m]||r;return t?a.createElement(d,o(o({ref:n},p),{},{components:t})):a.createElement(d,o({ref:n},p))}));function d(e,n){var t=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var r=t.length,o=new Array(r);o[0]=m;var s={};for(var c in n)hasOwnProperty.call(n,c)&&(s[c]=n[c]);s.originalType=e,s[h]="string"==typeof e?e:i,o[1]=s;for(var l=2;l<r;l++)o[l]=t[l];return a.createElement.apply(null,o)}return a.createElement.apply(null,t)}m.displayName="MDXCreateElement"},61674:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>o,default:()=>u,frontMatter:()=>r,metadata:()=>s,toc:()=>l});var a=t(87462),i=(t(67294),t(3905));const r={title:"Active Channels",sidebar_label:"Active Channels",sidebar_position:9,slug:"/apps/interchain-accounts/active-channels"},o="Understanding Active Channels",s={unversionedId:"apps/interchain-accounts/active-channels",id:"version-v7.3.x/apps/interchain-accounts/active-channels",title:"Active Channels",description:"The Interchain Accounts module uses ORDERED channels to maintain the order of transactions when sending packets from a controller to a host chain. A limitation when using ORDERED channels is that when a packet times out the channel will be closed.",source:"@site/versioned_docs/version-v7.3.x/02-apps/02-interchain-accounts/09-active-channels.md",sourceDirName:"02-apps/02-interchain-accounts",slug:"/apps/interchain-accounts/active-channels",permalink:"/v7/apps/interchain-accounts/active-channels",draft:!1,tags:[],version:"v7.3.x",sidebarPosition:9,frontMatter:{title:"Active Channels",sidebar_label:"Active Channels",sidebar_position:9,slug:"/apps/interchain-accounts/active-channels"},sidebar:"defaultSidebar",previous:{title:"Client",permalink:"/v7/apps/interchain-accounts/client"},next:{title:"Authentication Modules",permalink:"/v7/apps/interchain-accounts/legacy/auth-modules"}},c={},l=[{value:"Future improvements",id:"future-improvements",level:2}],p={toc:l},h="wrapper";function u(e){let{components:n,...t}=e;return(0,i.kt)(h,(0,a.Z)({},p,t,{components:n,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"understanding-active-channels"},"Understanding Active Channels"),(0,i.kt)("p",null,"The Interchain Accounts module uses ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc/tree/master/spec/core/ics-004-channel-and-packet-semantics#ordering"},"ORDERED channels")," to maintain the order of transactions when sending packets from a controller to a host chain. A limitation when using ORDERED channels is that when a packet times out the channel will be closed."),(0,i.kt)("p",null,"In the case of a channel closing, a controller chain needs to be able to regain access to the interchain account registered on this channel. ",(0,i.kt)("inlineCode",{parentName:"p"},"Active Channels")," enable this functionality."),(0,i.kt)("p",null,"When an Interchain Account is registered using ",(0,i.kt)("inlineCode",{parentName:"p"},"MsgRegisterInterchainAccount"),", a new channel is created on a particular port. During the ",(0,i.kt)("inlineCode",{parentName:"p"},"OnChanOpenAck")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"OnChanOpenConfirm")," steps (on controller & host chain respectively) the ",(0,i.kt)("inlineCode",{parentName:"p"},"Active Channel")," for this interchain account is stored in state."),(0,i.kt)("p",null,"It is possible to create a new channel using the same controller chain portID if the previously set ",(0,i.kt)("inlineCode",{parentName:"p"},"Active Channel")," is now in a ",(0,i.kt)("inlineCode",{parentName:"p"},"CLOSED")," state. This channel creation can be initialized programatically by sending a new ",(0,i.kt)("inlineCode",{parentName:"p"},"MsgChannelOpenInit")," message like so:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},"msg := channeltypes.NewMsgChannelOpenInit(portID, string(versionBytes), channeltypes.ORDERED, []string{connectionID}, icatypes.HostPortID, authtypes.NewModuleAddress(icatypes.ModuleName).String())\nhandler := keeper.msgRouter.Handler(msg)\nres, err := handler(ctx, msg)\nif err != nil {\n  return err\n}\n")),(0,i.kt)("p",null,"Alternatively, any relayer operator may initiate a new channel handshake for this interchain account once the previously set ",(0,i.kt)("inlineCode",{parentName:"p"},"Active Channel")," is in a ",(0,i.kt)("inlineCode",{parentName:"p"},"CLOSED")," state. This is done by initiating the channel handshake on the controller chain using the same portID associated with the interchain account in question.  "),(0,i.kt)("p",null,"It is important to note that once a channel has been opened for a given interchain account, new channels can not be opened for this account until the currently set ",(0,i.kt)("inlineCode",{parentName:"p"},"Active Channel")," is set to ",(0,i.kt)("inlineCode",{parentName:"p"},"CLOSED"),"."),(0,i.kt)("h2",{id:"future-improvements"},"Future improvements"),(0,i.kt)("p",null,"Future versions of the ICS-27 protocol and the Interchain Accounts module will likely use a new channel type that provides ordering of packets without the channel closing in the event of a packet timing out, thus removing the need for ",(0,i.kt)("inlineCode",{parentName:"p"},"Active Channels")," entirely.\nThe following is a list of issues which will provide the infrastructure to make this possible:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/ibc-go/issues/1599"},"IBC Channel Upgrades")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/ibc-go/issues/1661"},"Implement ORDERED_ALLOW_TIMEOUT logic in 04-channel")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/ibc-go/issues/1662"},"Add ORDERED_ALLOW_TIMEOUT as supported ordering in 03-connection")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"https://github.com/cosmos/ibc-go/issues/1663"},"Allow ICA channels to be opened as ORDERED_ALLOW_TIMEOUT"))))}u.isMDXComponent=!0}}]);