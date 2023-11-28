"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[6888],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>h});var i=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,i,a=function(e,t){if(null==e)return{};var n,i,a={},r=Object.keys(e);for(i=0;i<r.length;i++)n=r[i],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(i=0;i<r.length;i++)n=r[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=i.createContext({}),c=function(e){var t=i.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=c(e.components);return i.createElement(l.Provider,{value:t},e.children)},d="mdxType",u={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},m=i.forwardRef((function(e,t){var n=e.components,a=e.mdxType,r=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),d=c(n),m=a,h=d["".concat(l,".").concat(m)]||d[m]||u[m]||r;return n?i.createElement(h,o(o({ref:t},p),{},{components:n})):i.createElement(h,o({ref:t},p))}));function h(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var r=n.length,o=new Array(r);o[0]=m;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s[d]="string"==typeof e?e:a,o[1]=s;for(var c=2;c<r;c++)o[c]=n[c];return i.createElement.apply(null,o)}return i.createElement.apply(null,n)}m.displayName="MDXCreateElement"},1012:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>u,frontMatter:()=>r,metadata:()=>s,toc:()=>c});var i=n(87462),a=(n(67294),n(3905));const r={title:"Consensus State interface",sidebar_label:"Consensus State interface",sidebar_position:3,slug:"/ibc/light-clients/consensus-state"},o="Implementing the ConsensusState interface",s={unversionedId:"light-clients/developer-guide/consensus-state",id:"version-v8.0.x/light-clients/developer-guide/consensus-state",title:"Consensus State interface",description:"A ConsensusState is the snapshot of the counterparty chain, that an IBC client uses to verify proofs (e.g. a block).",source:"@site/versioned_docs/version-v8.0.x/03-light-clients/01-developer-guide/03-consensus-state.md",sourceDirName:"03-light-clients/01-developer-guide",slug:"/ibc/light-clients/consensus-state",permalink:"/v8/ibc/light-clients/consensus-state",draft:!1,tags:[],version:"v8.0.x",sidebarPosition:3,frontMatter:{title:"Consensus State interface",sidebar_label:"Consensus State interface",sidebar_position:3,slug:"/ibc/light-clients/consensus-state"},sidebar:"defaultSidebar",previous:{title:"Client State interface",permalink:"/v8/ibc/light-clients/client-state"},next:{title:"Handling Updates and Misbehaviour",permalink:"/v8/ibc/light-clients/updates-and-misbehaviour"}},l={},c=[{value:"<code>ClientType</code> method",id:"clienttype-method",level:2},{value:"<code>GetTimestamp</code> method",id:"gettimestamp-method",level:2},{value:"<code>ValidateBasic</code> method",id:"validatebasic-method",level:2}],p={toc:c},d="wrapper";function u(e){let{components:t,...n}=e;return(0,a.kt)(d,(0,i.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"implementing-the-consensusstate-interface"},"Implementing the ",(0,a.kt)("inlineCode",{parentName:"h1"},"ConsensusState")," interface"),(0,a.kt)("p",null,"A ",(0,a.kt)("inlineCode",{parentName:"p"},"ConsensusState")," is the snapshot of the counterparty chain, that an IBC client uses to verify proofs (e.g. a block)."),(0,a.kt)("p",null,"The further development of multiple types of IBC light clients and the difficulties presented by this generalization problem (see ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/main/docs/architecture/adr-006-02-client-refactor.md"},"ADR-006")," for more information about this historical context) led to the design decision of each client keeping track of and set its own ",(0,a.kt)("inlineCode",{parentName:"p"},"ClientState")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"ConsensusState"),", as well as the simplification of client ",(0,a.kt)("inlineCode",{parentName:"p"},"ConsensusState")," updates through the generalized ",(0,a.kt)("inlineCode",{parentName:"p"},"ClientMessage")," interface."),(0,a.kt)("p",null,"The below ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/core/exported/client.go#L133"},(0,a.kt)("inlineCode",{parentName:"a"},"ConsensusState"))," interface is a generalized interface for the types of information a ",(0,a.kt)("inlineCode",{parentName:"p"},"ConsensusState")," could contain. For a reference ",(0,a.kt)("inlineCode",{parentName:"p"},"ConsensusState")," implementation, please see the ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/light-clients/07-tendermint/consensus_state.go"},"Tendermint light client ",(0,a.kt)("inlineCode",{parentName:"a"},"ConsensusState")),"."),(0,a.kt)("h2",{id:"clienttype-method"},(0,a.kt)("inlineCode",{parentName:"h2"},"ClientType")," method"),(0,a.kt)("p",null,"This is the type of client consensus. It should be the same as the ",(0,a.kt)("inlineCode",{parentName:"p"},"ClientType")," return value for the ",(0,a.kt)("a",{parentName:"p",href:"/v8/ibc/light-clients/client-state"},"corresponding ",(0,a.kt)("inlineCode",{parentName:"a"},"ClientState")," implementation"),"."),(0,a.kt)("h2",{id:"gettimestamp-method"},(0,a.kt)("inlineCode",{parentName:"h2"},"GetTimestamp")," method"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"GetTimestamp")," should return the timestamp (in nanoseconds) of the consensus state snapshot."),(0,a.kt)("h2",{id:"validatebasic-method"},(0,a.kt)("inlineCode",{parentName:"h2"},"ValidateBasic")," method"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"ValidateBasic")," should validate every consensus state field and should return an error if any value is invalid. The light client implementer is in charge of determining which checks are required."))}u.isMDXComponent=!0}}]);