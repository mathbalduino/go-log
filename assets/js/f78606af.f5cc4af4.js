"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[657],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return m}});var i=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);t&&(i=i.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,i)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,i,o=function(e,t){if(null==e)return{};var n,i,o={},a=Object.keys(e);for(i=0;i<a.length;i++)n=a[i],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(i=0;i<a.length;i++)n=a[i],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var c=i.createContext({}),s=function(e){var t=i.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},u=function(e){var t=s(e.components);return i.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return i.createElement(i.Fragment,{},t)}},d=i.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,c=e.parentName,u=l(e,["components","mdxType","originalType","parentName"]),d=s(n),m=o,g=d["".concat(c,".").concat(m)]||d[m]||p[m]||a;return n?i.createElement(g,r(r({ref:t},u),{},{components:n})):i.createElement(g,r({ref:t},u))}));function m(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,r=new Array(a);r[0]=d;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l.mdxType="string"==typeof e?e:o,r[1]=l;for(var s=2;s<a;s++)r[s]=n[s];return i.createElement.apply(null,r)}return i.createElement.apply(null,n)}d.displayName="MDXCreateElement"},6954:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return l},contentTitle:function(){return c},metadata:function(){return s},toc:function(){return u},default:function(){return d}});var i=n(7462),o=n(3366),a=(n(7294),n(3905)),r=["components"],l={sidebar_position:3},c="Dynamic Configuration",s={unversionedId:"advanced/dynamic_config",id:"advanced/dynamic_config",isDocsHomePage:!1,title:"Dynamic Configuration",description:"It's possible to change the Logger Configuration while it's executing, using the Logger.Configuration() method, but you will have to be careful.",source:"@site/docs/advanced/dynamic_config.md",sourceDirName:"advanced",slug:"/advanced/dynamic_config",permalink:"/go-log/docs/advanced/dynamic_config",editUrl:"https://github.com/mathbalduino/go-log/edit/main/docs/docs/advanced/dynamic_config.md",tags:[],version:"current",sidebarPosition:3,frontMatter:{sidebar_position:3},sidebar:"tutorialSidebar",previous:{title:"Log creation",permalink:"/go-log/docs/advanced/log_creation"},next:{title:"Error tokens",permalink:"/go-log/docs/advanced/error_tokens"}},u=[{value:"Step by step",id:"step-by-step",children:[]}],p={toc:u};function d(e){var t=e.components,n=(0,o.Z)(e,r);return(0,a.kt)("wrapper",(0,i.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"dynamic-configuration"},"Dynamic Configuration"),(0,a.kt)("p",null,"It's possible to change the ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," while it's executing, using the ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger.Configuration()")," method, but you will have to be careful."),(0,a.kt)("p",null,"It's always recommended that you use just one go routine to modify the values of the ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," struct, otherwise you will need to fix write-concurrency issues between the writers. The library will ",(0,a.kt)("em",{parentName:"p"},"never")," modify the values of the ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," by itself, only ",(0,a.kt)("em",{parentName:"p"},"read"),"."),(0,a.kt)("h2",{id:"step-by-step"},"Step by step"),(0,a.kt)("p",null,"When using only one writer, and knowing that the library will just ",(0,a.kt)("strong",{parentName:"p"},"read")," the ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration"),", you can safely just call the ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger.Configuration()")," method passing a new ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," struct."),(0,a.kt)("p",null,"The library doesn't control some concurrency issues that can arise when you do it. If you want to change just one ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," struct field, like ",(0,a.kt)("inlineCode",{parentName:"p"},"LvlsEnabled"),", there's no issue. Note that since the library just ",(0,a.kt)("em",{parentName:"p"},"reads")," the ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration"),", and you changed just one thing, it's an atomic operation."),(0,a.kt)("p",null,"If you needs to change more than one thing, it is, naturally, not atomic. The library, right now, doesn't have a way to guarantee the atomicity of these operations. It means that if you want to change the ",(0,a.kt)("inlineCode",{parentName:"p"},"LvlFieldName")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"MsgFieldName")," at once, logs will be saved in one of the following states:"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"Old ",(0,a.kt)("inlineCode",{parentName:"li"},"LvlFieldName")," and ",(0,a.kt)("inlineCode",{parentName:"li"},"MsgFieldName")," values (before ",(0,a.kt)("inlineCode",{parentName:"li"},"Logger.Configuration()")," call)"),(0,a.kt)("li",{parentName:"ol"},"Old ",(0,a.kt)("inlineCode",{parentName:"li"},"LvlFieldName")," value with the new ",(0,a.kt)("inlineCode",{parentName:"li"},"MsgFieldName")," value (some log was being created at the same time of the ",(0,a.kt)("inlineCode",{parentName:"li"},"Logger.Configuration()")," call)"),(0,a.kt)("li",{parentName:"ol"},"New ",(0,a.kt)("inlineCode",{parentName:"li"},"LvlFieldName")," value with the old ",(0,a.kt)("inlineCode",{parentName:"li"},"MsgFieldName")," value (some log was being created at the same time of the ",(0,a.kt)("inlineCode",{parentName:"li"},"Logger.Configuration()")," call)"),(0,a.kt)("li",{parentName:"ol"},"New ",(0,a.kt)("inlineCode",{parentName:"li"},"LvlFieldName")," and ",(0,a.kt)("inlineCode",{parentName:"li"},"MsgFieldName")," values (",(0,a.kt)("inlineCode",{parentName:"li"},"Logger.Configuration()")," call completed)")),(0,a.kt)("p",null,"There's a plan to implement some blocking ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger.Configuration()")," method variation, that solves this issue, in the future. Please, let me know if it's necessary."))}d.isMDXComponent=!0}}]);