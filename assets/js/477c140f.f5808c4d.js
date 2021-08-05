(self.webpackChunkdoc_ops=self.webpackChunkdoc_ops||[]).push([[7603],{3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return l},kt:function(){return p}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function c(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var s=r.createContext({}),u=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):c(c({},t),e)),n},l=function(e){var t=u(e.components);return r.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},f=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,s=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),f=u(n),p=o,m=f["".concat(s,".").concat(p)]||f[p]||d[p]||a;return n?r.createElement(m,c(c({ref:t},l),{},{components:n})):r.createElement(m,c({ref:t},l))}));function p(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,c=new Array(a);c[0]=f;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:o,c[1]=i;for(var u=2;u<a;u++)c[u]=n[u];return r.createElement.apply(null,c)}return r.createElement.apply(null,n)}f.displayName="MDXCreateElement"},8215:function(e,t,n){"use strict";var r=n(7294);t.Z=function(e){var t=e.children,n=e.hidden,o=e.className;return r.createElement("div",{role:"tabpanel",hidden:n,className:o},t)}},5064:function(e,t,n){"use strict";n.d(t,{Z:function(){return d}});var r=n(7294),o=n(9443);var a=function(){var e=(0,r.useContext)(o.Z);if(null==e)throw new Error('"useUserPreferencesContext" is used outside of "Layout" component.');return e},c=n(6010),i="tabItem_1uMI",s="tabItemActive_2DSg";var u=37,l=39;var d=function(e){var t=e.lazy,n=e.block,o=e.defaultValue,d=e.values,f=e.groupId,p=e.className,m=a(),h=m.tabGroupChoices,v=m.setTabGroupChoices,w=(0,r.useState)(o),b=w[0],y=w[1],g=r.Children.toArray(e.children),O=[];if(null!=f){var k=h[f];null!=k&&k!==b&&d.some((function(e){return e.value===k}))&&y(k)}var x=function(e){var t=e.currentTarget,n=O.indexOf(t),r=d[n].value;y(r),null!=f&&(v(f,r),setTimeout((function(){var e,n,r,o,a,c,i,u;(e=t.getBoundingClientRect(),n=e.top,r=e.left,o=e.bottom,a=e.right,c=window,i=c.innerHeight,u=c.innerWidth,n>=0&&a<=u&&o<=i&&r>=0)||(t.scrollIntoView({block:"center",behavior:"smooth"}),t.classList.add(s),setTimeout((function(){return t.classList.remove(s)}),2e3))}),150))},N=function(e){var t,n;switch(e.keyCode){case l:var r=O.indexOf(e.target)+1;n=O[r]||O[0];break;case u:var o=O.indexOf(e.target)-1;n=O[o]||O[O.length-1]}null==(t=n)||t.focus()};return r.createElement("div",{className:"tabs-container"},r.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,c.Z)("tabs",{"tabs--block":n},p)},d.map((function(e){var t=e.value,n=e.label;return r.createElement("li",{role:"tab",tabIndex:b===t?0:-1,"aria-selected":b===t,className:(0,c.Z)("tabs__item",i,{"tabs__item--active":b===t}),key:t,ref:function(e){return O.push(e)},onKeyDown:N,onFocus:x,onClick:x},n)}))),t?(0,r.cloneElement)(g.filter((function(e){return e.props.value===b}))[0],{className:"margin-vert--md"}):r.createElement("div",{className:"margin-vert--md"},g.map((function(e,t){return(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==b})}))))}},9443:function(e,t,n){"use strict";var r=(0,n(7294).createContext)(void 0);t.Z=r},208:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return u},contentTitle:function(){return l},metadata:function(){return d},toc:function(){return f},default:function(){return m}});var r=n(2122),o=n(9756),a=(n(7294),n(3905)),c=n(5064),i=n(8215),s=["components"],u={},l="How to withdraw from a chain",d={unversionedId:"guide/core_concepts/accounts/accounts-withdraw",id:"guide/core_concepts/accounts/accounts-withdraw",isDocsHomePage:!1,title:"How to withdraw from a chain",description:"The withdraw endpoint sends funds owned by the callee to their L1 address.",source:"@site/docs/guide/core_concepts/accounts/accounts-withdraw.mdx",sourceDirName:"guide/core_concepts/accounts",slug:"/guide/core_concepts/accounts/accounts-withdraw",permalink:"/docs/guide/core_concepts/accounts/accounts-withdraw",editUrl:"https://github.com/iotaledger/chronicle.rs/tree/main/docs/docs/guide/core_concepts/accounts/accounts-withdraw.mdx",version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"How to deposit to a chain",permalink:"/docs/guide/core_concepts/accounts/accounts-deposit"},next:{title:"View account balances",permalink:"/docs/guide/core_concepts/accounts/accounts-views"}},f=[],p={toc:f};function m(e){var t=e.components,n=(0,o.Z)(e,s);return(0,a.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"how-to-withdraw-from-a-chain"},"How to withdraw from a chain"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"withdraw")," endpoint sends funds owned by the callee to their L1 address."),(0,a.kt)(c.Z,{defaultValue:"solo",values:[{label:"Solo",value:"solo"},{label:"Rust (ISCP)",value:"rust"}],mdxType:"Tabs"},(0,a.kt)(i.Z,{value:"solo",mdxType:"TabItem"},(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"_, err := chain.PostRequestSync(\n    solo.NewCallParams(accounts.Contract.Name, accounts.FuncDeposit.Name).WithIotas(N),\n    wallet,\n)\n"))),(0,a.kt)(i.Z,{value:"rust",mdxType:"TabItem"},(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-rust"},'// withdrawals from a chain (funds go to the common account, as there is no way to specify hname for withdrawals)\nlet transfer = ScTransfers::iotas(1);\nctx.post(chainID, ScHname::new("accounts"), ScHname::new("withdraw"), None, transfer, 0);\n')))))}m.isMDXComponent=!0},6010:function(e,t,n){"use strict";function r(e){var t,n,o="";if("string"==typeof e||"number"==typeof e)o+=e;else if("object"==typeof e)if(Array.isArray(e))for(t=0;t<e.length;t++)e[t]&&(n=r(e[t]))&&(o&&(o+=" "),o+=n);else for(t in e)e[t]&&(o&&(o+=" "),o+=t);return o}function o(){for(var e,t,n=0,o="";n<arguments.length;)(e=arguments[n++])&&(t=r(e))&&(o&&(o+=" "),o+=t);return o}n.d(t,{Z:function(){return o}})}}]);