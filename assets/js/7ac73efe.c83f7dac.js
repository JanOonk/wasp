(self.webpackChunkdoc_ops=self.webpackChunkdoc_ops||[]).push([[5016],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return l},kt:function(){return f}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function c(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?c(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):c(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},c=Object.keys(e);for(r=0;r<c.length;r++)n=c[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(r=0;r<c.length;r++)n=c[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),u=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},l=function(e){var t=u(e.components);return r.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,c=e.originalType,s=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),d=u(n),f=o,m=d["".concat(s,".").concat(f)]||d[f]||p[f]||c;return n?r.createElement(m,a(a({ref:t},l),{},{components:n})):r.createElement(m,a({ref:t},l))}));function f(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var c=n.length,a=new Array(c);a[0]=d;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:o,a[1]=i;for(var u=2;u<c;u++)a[u]=n[u];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},8215:function(e,t,n){"use strict";var r=n(7294);t.Z=function(e){var t=e.children,n=e.hidden,o=e.className;return r.createElement("div",{role:"tabpanel",hidden:n,className:o},t)}},5064:function(e,t,n){"use strict";n.d(t,{Z:function(){return p}});var r=n(7294),o=n(9443);var c=function(){var e=(0,r.useContext)(o.Z);if(null==e)throw new Error('"useUserPreferencesContext" is used outside of "Layout" component.');return e},a=n(6010),i="tabItem_1uMI",s="tabItemActive_2DSg";var u=37,l=39;var p=function(e){var t=e.lazy,n=e.block,o=e.defaultValue,p=e.values,d=e.groupId,f=e.className,m=c(),v=m.tabGroupChoices,b=m.setTabGroupChoices,h=(0,r.useState)(o),y=h[0],g=h[1],w=r.Children.toArray(e.children),O=[];if(null!=d){var k=v[d];null!=k&&k!==y&&p.some((function(e){return e.value===k}))&&g(k)}var N=function(e){var t=e.currentTarget,n=O.indexOf(t),r=p[n].value;g(r),null!=d&&(b(d,r),setTimeout((function(){var e,n,r,o,c,a,i,u;(e=t.getBoundingClientRect(),n=e.top,r=e.left,o=e.bottom,c=e.right,a=window,i=a.innerHeight,u=a.innerWidth,n>=0&&c<=u&&o<=i&&r>=0)||(t.scrollIntoView({block:"center",behavior:"smooth"}),t.classList.add(s),setTimeout((function(){return t.classList.remove(s)}),2e3))}),150))},T=function(e){var t,n;switch(e.keyCode){case l:var r=O.indexOf(e.target)+1;n=O[r]||O[0];break;case u:var o=O.indexOf(e.target)-1;n=O[o]||O[O.length-1]}null==(t=n)||t.focus()};return r.createElement("div",{className:"tabs-container"},r.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,a.Z)("tabs",{"tabs--block":n},f)},p.map((function(e){var t=e.value,n=e.label;return r.createElement("li",{role:"tab",tabIndex:y===t?0:-1,"aria-selected":y===t,className:(0,a.Z)("tabs__item",i,{"tabs__item--active":y===t}),key:t,ref:function(e){return O.push(e)},onKeyDown:T,onFocus:N,onClick:N},n)}))),t?(0,r.cloneElement)(w.filter((function(e){return e.props.value===y}))[0],{className:"margin-vert--md"}):r.createElement("div",{className:"margin-vert--md"},w.map((function(e,t){return(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==y})}))))}},9443:function(e,t,n){"use strict";var r=(0,n(7294).createContext)(void 0);t.Z=r},5657:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return u},contentTitle:function(){return l},metadata:function(){return p},toc:function(){return d},default:function(){return m}});var r=n(2122),o=n(9756),c=(n(7294),n(3905)),a=n(5064),i=n(8215),s=["components"],u={},l="How to deposit to a chain",p={unversionedId:"guide/core_concepts/accounts/accounts-deposit",id:"guide/core_concepts/accounts/accounts-deposit",isDocsHomePage:!1,title:"How to deposit to a chain",description:"The deposit entrypoint credits the transfered IOTAS into your on-chain account.",source:"@site/docs/guide/core_concepts/accounts/accounts-deposit.mdx",sourceDirName:"guide/core_concepts/accounts",slug:"/guide/core_concepts/accounts/accounts-deposit",permalink:"/docs/guide/core_concepts/accounts/accounts-deposit",editUrl:"https://github.com/iotaledger/chronicle.rs/tree/main/docs/docs/guide/core_concepts/accounts/accounts-deposit.mdx",version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"How accounts work",permalink:"/docs/guide/core_concepts/accounts/accounts-how"},next:{title:"How to withdraw from a chain",permalink:"/docs/guide/core_concepts/accounts/accounts-withdraw"}},d=[],f={toc:d};function m(e){var t=e.components,n=(0,o.Z)(e,s);return(0,c.kt)("wrapper",(0,r.Z)({},f,n,{components:t,mdxType:"MDXLayout"}),(0,c.kt)("h1",{id:"how-to-deposit-to-a-chain"},"How to deposit to a chain"),(0,c.kt)("p",null,"The ",(0,c.kt)("inlineCode",{parentName:"p"},"deposit")," entrypoint credits the transfered IOTAS into your on-chain account."),(0,c.kt)(a.Z,{defaultValue:"solo",values:[{label:"Solo",value:"solo"},{label:"Rust (ISCP)",value:"rust"}],mdxType:"Tabs"},(0,c.kt)(i.Z,{value:"solo",mdxType:"TabItem"},(0,c.kt)("pre",null,(0,c.kt)("code",{parentName:"pre",className:"language-go"},"_, err := chain.PostRequestSync(\n    solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(N),\n    wallet,\n)\n"))),(0,c.kt)(i.Z,{value:"rust",mdxType:"TabItem"},(0,c.kt)("pre",null,(0,c.kt)("code",{parentName:"pre",className:"language-rust"},'// deposits N iotas from current SC into another chain\nlet transfer = ScTransfers::iotas(N);\nctx.post(chainID, ScHname::new("accounts"), ScHname::new("deposit") , None, transfer,0);\n')))))}m.isMDXComponent=!0},6010:function(e,t,n){"use strict";function r(e){var t,n,o="";if("string"==typeof e||"number"==typeof e)o+=e;else if("object"==typeof e)if(Array.isArray(e))for(t=0;t<e.length;t++)e[t]&&(n=r(e[t]))&&(o&&(o+=" "),o+=n);else for(t in e)e[t]&&(o&&(o+=" "),o+=t);return o}function o(){for(var e,t,n=0,o="";n<arguments.length;)(e=arguments[n++])&&(t=r(e))&&(o&&(o+=" "),o+=t);return o}n.d(t,{Z:function(){return o}})}}]);