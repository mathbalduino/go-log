"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[175],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return m}});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var p=r.createContext({}),s=function(e){var t=r.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=s(e.components);return r.createElement(p.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},g=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,p=e.parentName,u=c(e,["components","mdxType","originalType","parentName"]),g=s(n),m=o,d=g["".concat(p,".").concat(m)]||g[m]||l[m]||a;return n?r.createElement(d,i(i({ref:t},u),{},{components:n})):r.createElement(d,i({ref:t},u))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,i=new Array(a);i[0]=g;var c={};for(var p in t)hasOwnProperty.call(t,p)&&(c[p]=t[p]);c.originalType=e,c.mdxType="string"==typeof e?e:o,i[1]=c;for(var s=2;s<a;s++)i[s]=n[s];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}g.displayName="MDXCreateElement"},2676:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return c},contentTitle:function(){return p},metadata:function(){return s},toc:function(){return u},default:function(){return g}});var r=n(7462),o=n(3366),a=(n(7294),n(3905)),i=["components"],c={sidebar_position:11},p="Logger creation",s={unversionedId:"basic-concepts/logger_creation",id:"basic-concepts/logger_creation",isDocsHomePage:!1,title:"Logger creation",description:"To create a new Logger, you can use these two functions exported by the root package:",source:"@site/docs/basic-concepts/logger_creation.md",sourceDirName:"basic-concepts",slug:"/basic-concepts/logger_creation",permalink:"/docs/basic-concepts/logger_creation",editUrl:"https://github.com/mathbalduino/go-log/edit/main/docs/docs/basic-concepts/logger_creation.md",tags:[],version:"current",sidebarPosition:11,frontMatter:{sidebar_position:11},sidebar:"tutorialSidebar",previous:{title:"Async Logger",permalink:"/docs/basic-concepts/async_logger"},next:{title:"Configuration",permalink:"/docs/basic-concepts/configuration"}},u=[],l={toc:u};function g(e){var t=e.components,n=(0,o.Z)(e,i);return(0,a.kt)("wrapper",(0,r.Z)({},l,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"logger-creation"},"Logger creation"),(0,a.kt)("p",null,"To create a new ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger"),", you can use these two functions exported by the root package:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"// new.go\nfunc New(config Configuration) *Logger { ... }\nfunc NewDefault() *Logger { ... }\n")),(0,a.kt)("p",null,"The first one will create an empty ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," (with the given ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration"),"), while the latter one will create a ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," using the ",(0,a.kt)("a",{parentName:"p",href:"/docs/basic-concepts/configuration#default-configuration"},"Default Configuration")," and setting two outputs (in this order): "),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"/docs/basic-concepts/outputs#outputansitostdout"},"OutputAnsiToStdout")),(0,a.kt)("li",{parentName:"ol"},(0,a.kt)("a",{parentName:"li",href:"/docs/basic-concepts/outputs#outputpaniconfatal"},"OutputPanicOnFatal"))),(0,a.kt)("p",null,"More information about why the order of the ",(0,a.kt)("inlineCode",{parentName:"p"},"Outputs")," is important ",(0,a.kt)("a",{parentName:"p",href:"/docs/basic-concepts/outputs#ordering"},"here"),"."),(0,a.kt)("div",{className:"admonition admonition-caution alert alert--warning"},(0,a.kt)("div",{parentName:"div",className:"admonition-heading"},(0,a.kt)("h5",{parentName:"div"},(0,a.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,a.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"16",height:"16",viewBox:"0 0 16 16"},(0,a.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))),"caution")),(0,a.kt)("div",{parentName:"div",className:"admonition-content"},(0,a.kt)("p",{parentName:"div"},"Before using the ",(0,a.kt)("inlineCode",{parentName:"p"},"New")," function, check the created ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration"),". If there's some error inside it (",(0,a.kt)("inlineCode",{parentName:"p"},"lvl")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"msg")," using the same string, for example), it will ",(0,a.kt)("inlineCode",{parentName:"p"},"panic"),". More information about how to structure your ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," struct ",(0,a.kt)("a",{parentName:"p",href:"/docs/basic-concepts/configuration"},"here")))))}g.isMDXComponent=!0}}]);