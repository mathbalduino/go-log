"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[464],{3905:function(e,t,n){n.d(t,{Zo:function(){return c},kt:function(){return h}});var a=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,i=function(e,t){if(null==e)return{};var n,a,i={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var s=a.createContext({}),p=function(e){var t=a.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},c=function(e){var t=p(e.components);return a.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),m=p(n),h=i,u=m["".concat(s,".").concat(h)]||m[h]||d[h]||o;return n?a.createElement(u,r(r({ref:t},c),{},{components:n})):a.createElement(u,r({ref:t},c))}));function h(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,r=new Array(o);r[0]=m;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:i,r[1]=l;for(var p=2;p<o;p++)r[p]=n[p];return a.createElement.apply(null,r)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},4101:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return s},metadata:function(){return p},toc:function(){return c},default:function(){return m}});var a=n(7462),i=n(3366),o=(n(7294),n(3905)),r=["components"],l={sidebar_position:8},s="Life cycle",p={unversionedId:"basic-concepts/life_cycle",id:"basic-concepts/life_cycle",isDocsHomePage:!1,title:"Life cycle",description:"Any created log will pass between four phases:",source:"@site/docs/basic-concepts/life_cycle.md",sourceDirName:"basic-concepts",slug:"/basic-concepts/life_cycle",permalink:"/go-log/docs/basic-concepts/life_cycle",editUrl:"https://github.com/mathbalduino/go-log/edit/main/docs/docs/basic-concepts/life_cycle.md",tags:[],version:"current",sidebarPosition:8,frontMatter:{sidebar_position:8},sidebar:"tutorialSidebar",previous:{title:"Outputs",permalink:"/go-log/docs/basic-concepts/outputs"},next:{title:"Log levels",permalink:"/go-log/docs/basic-concepts/log_levels"}},c=[{value:"Sync Phase 1: Creation",id:"sync-phase-1-creation",children:[]},{value:"Sync Phase 2: Pre handling",id:"sync-phase-2-pre-handling",children:[]},{value:"(A)Sync Phase 3: Post handling",id:"async-phase-3-post-handling",children:[]},{value:"(A)Sync Phase 4: Output",id:"async-phase-4-output",children:[]}],d={toc:c};function m(e){var t=e.components,n=(0,i.Z)(e,r);return(0,o.kt)("wrapper",(0,a.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"life-cycle"},"Life cycle"),(0,o.kt)("p",null,"Any created log will pass between four phases:"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("inlineCode",{parentName:"li"},"sync")," ",(0,o.kt)("a",{parentName:"li",href:"#sync-phase-1-creation"},"Creation"),": The log is created and ",(0,o.kt)("inlineCode",{parentName:"li"},"AdHoc fields")," are collected"),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("inlineCode",{parentName:"li"},"sync")," ",(0,o.kt)("a",{parentName:"li",href:"#sync-phase-2-pre-handling"},"Pre handling"),": ",(0,o.kt)("inlineCode",{parentName:"li"},"PreHooks")," are evaluated"),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("inlineCode",{parentName:"li"},"sync"),"/",(0,o.kt)("inlineCode",{parentName:"li"},"async")," ",(0,o.kt)("a",{parentName:"li",href:"#async-phase-3-post-handling"},"Post handling"),": ",(0,o.kt)("inlineCode",{parentName:"li"},"Base fields"),", ",(0,o.kt)("inlineCode",{parentName:"li"},"PreHooks"),", ",(0,o.kt)("inlineCode",{parentName:"li"},"AdHoc fields")," are applied and ",(0,o.kt)("inlineCode",{parentName:"li"},"PostHooks")," are evaluated and applied"),(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("inlineCode",{parentName:"li"},"sync"),"/",(0,o.kt)("inlineCode",{parentName:"li"},"async")," ",(0,o.kt)("a",{parentName:"li",href:"#async-phase-4-output"},"Output"),": The final log fields are forwarded to every configured ",(0,o.kt)("inlineCode",{parentName:"li"},"Output"))),(0,o.kt)("h2",{id:"sync-phase-1-creation"},"Sync Phase 1: Creation"),(0,o.kt)("p",null,"This phase is characterized by the call to the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger")," log level method itself:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'someLogger.Trace("some msg", logger.LogFields{ ...someFields... })\n')),(0,o.kt)("p",null,"At this point, the ",(0,o.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," are collected and a referece to them is stored, in order to be able to process them by the following phases."),(0,o.kt)("h2",{id:"sync-phase-2-pre-handling"},"Sync Phase 2: Pre handling"),(0,o.kt)("p",null,"Right after the ",(0,o.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," are collected, the ",(0,o.kt)("inlineCode",{parentName:"p"},"PreHooks")," are evaluated. It means that all the configured ",(0,o.kt)("inlineCode",{parentName:"p"},"PreHooks")," functions will be called at this stage, in a synchronous way. The returned ",(0,o.kt)("inlineCode",{parentName:"p"},"PreHooks")," function values will be stored inside the created log, in order to be applied later."),(0,o.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"note")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"At this stage, the log fields aren't ready yet. We're just calculating the ",(0,o.kt)("em",{parentName:"p"},"possible")," values of the fields. These values can be overriden by ",(0,o.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," or ",(0,o.kt)("inlineCode",{parentName:"p"},"PostHooks"),", at later phases, but we don't know for sure."))),(0,o.kt)("h2",{id:"async-phase-3-post-handling"},"(A)Sync Phase 3: Post handling"),(0,o.kt)("p",null,"This is the ",(0,o.kt)("strong",{parentName:"p"},"most expensive phase"),", since it's responsible to let the log fields ready to be used by the ",(0,o.kt)("inlineCode",{parentName:"p"},"Outputs"),"."),(0,o.kt)("p",null,"Here, the ",(0,o.kt)("inlineCode",{parentName:"p"},"Base fields")," are copied from the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger")," instance that created the log into a new ",(0,o.kt)("inlineCode",{parentName:"p"},"LogFields")," map, followed by the copy of the evaluated ",(0,o.kt)("inlineCode",{parentName:"p"},"PreHooks")," (from phase 2), followed by the copy of the ",(0,o.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," (from phase 1). The ",(0,o.kt)("inlineCode",{parentName:"p"},"PostHooks")," values will be last to be evaluated and copied to the final log fields. More information, see the ",(0,o.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/override_order"},"Fields override order")," page."),(0,o.kt)("p",null,"At the end of the phase 3, the ",(0,o.kt)("inlineCode",{parentName:"p"},"message")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"level")," log fields will be applied to the ",(0,o.kt)("inlineCode",{parentName:"p"},"LogFields")," map, overriding any previous value that used the keys respectively configured. Something like this:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"logFields[lvlFieldKey] = log.lvl\nlogFields[msgFieldKey] = log.msg\n")),(0,o.kt)("div",{className:"admonition admonition-info alert alert--info"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"}))),"info")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"If the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger")," is set to be ",(0,o.kt)("inlineCode",{parentName:"p"},"async"),", this phase will be executed by a different go routine than the one that created the log"))),(0,o.kt)("h2",{id:"async-phase-4-output"},"(A)Sync Phase 4: Output"),(0,o.kt)("p",null,"Just a ",(0,o.kt)("inlineCode",{parentName:"p"},"for")," loop over the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger")," configured ",(0,o.kt)("inlineCode",{parentName:"p"},"Outputs")," slice, forwarding the final log fields, created at the phase 3:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// Not real production-code (just to illustrate)\nfor _, output := range logger.outputs {\n  output(lvl, msg, logFields)\n}\n")),(0,o.kt)("div",{className:"admonition admonition-info alert alert--info"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"}))),"info")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"If the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger")," is set to be ",(0,o.kt)("inlineCode",{parentName:"p"},"async"),", this phase will be executed by a different go routine than the one that created the log"))))}m.isMDXComponent=!0}}]);