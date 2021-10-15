"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[755],{3905:function(e,t,n){n.d(t,{Zo:function(){return p},kt:function(){return g}});var o=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},a=Object.keys(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var s=o.createContext({}),c=function(e){var t=o.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},p=function(e){var t=c(e.components);return o.createElement(s.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},u=o.forwardRef((function(e,t){var n=e.components,r=e.mdxType,a=e.originalType,s=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),u=c(n),g=r,f=u["".concat(s,".").concat(g)]||u[g]||d[g]||a;return n?o.createElement(f,l(l({ref:t},p),{},{components:n})):o.createElement(f,l({ref:t},p))}));function g(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var a=n.length,l=new Array(a);l[0]=u;var i={};for(var s in t)hasOwnProperty.call(t,s)&&(i[s]=t[s]);i.originalType=e,i.mdxType="string"==typeof e?e:r,l[1]=i;for(var c=2;c<a;c++)l[c]=n[c];return o.createElement.apply(null,l)}return o.createElement.apply(null,n)}u.displayName="MDXCreateElement"},6172:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return i},contentTitle:function(){return s},metadata:function(){return c},toc:function(){return p},default:function(){return u}});var o=n(7462),r=n(3366),a=(n(7294),n(3905)),l=["components"],i={sidebar_position:4},s="AdHoc fields",c={unversionedId:"basic-concepts/adhoc_fields",id:"basic-concepts/adhoc_fields",isDocsHomePage:!1,title:"AdHoc fields",description:"AdHoc fields will be applied at the Thirty phase of the life cycle, right after the PreHooks are executed.",source:"@site/docs/basic-concepts/adhoc_fields.md",sourceDirName:"basic-concepts",slug:"/basic-concepts/adhoc_fields",permalink:"/go-log/docs/basic-concepts/adhoc_fields",editUrl:"https://github.com/mathbalduino/go-log/edit/docs/docs/basic-concepts/adhoc_fields.md",tags:[],version:"current",sidebarPosition:4,frontMatter:{sidebar_position:4},sidebar:"tutorialSidebar",previous:{title:"PreHooks",permalink:"/go-log/docs/basic-concepts/pre_hooks"},next:{title:"PostHooks",permalink:"/go-log/docs/basic-concepts/post_hooks"}},p=[],d={toc:p};function u(e){var t=e.components,n=(0,r.Z)(e,l);return(0,a.kt)("wrapper",(0,o.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"adhoc-fields"},"AdHoc fields"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," will be applied at the ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/life_cycle#async-phase-3-post-handling"},"Thirty phase")," of the life cycle, right after the ",(0,a.kt)("inlineCode",{parentName:"p"},"PreHooks")," are executed."),(0,a.kt)("p",null,"Every Log level method can accept ",(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," (even custom ones), because all log levels must call ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger.Log()")," in the end (that accepts ",(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields"),"):"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Log(lvl uint64, msg string, adHocFields []LogFields) { ... }\n")),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger.Log()")," receives the ",(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," as a ",(0,a.kt)("inlineCode",{parentName:"p"},"slice")," of ",(0,a.kt)("inlineCode",{parentName:"p"},"LogFields")," just to ease the forwarding from the log custom levels: "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"// logLevels.go\nfunc Trace(msg string, adHocFields ...LogFields) { \n  l.Log(LvlTrace, msg, adHocFields)\n}\n")),(0,a.kt)("p",null,"Example of ",(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," usage:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'someLogger := logger.New(logger.DefaultConfig()).\n  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\nsomeLogger.Info("some log", logger.LogFields{\n  "adHoc-A": "value-A",\n  "adHoc-B": "value-B",\n})\n/*\n  {\n    "msg": "some log",\n    "lvl": 4,\n    "adHoc-A": "value-A",\n    "adHoc-B": "value-B"\n  }\n*/\n')),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields"),' are defined as variadic arguments just to simulate "optional arguments", that don\'t officially exist in ',(0,a.kt)("inlineCode",{parentName:"p"},"go"),". Note that if you pass more than one ",(0,a.kt)("inlineCode",{parentName:"p"},"LogFields")," variadic argument, the latter ones will override the previous ones:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'someLogger := logger.New(logger.DefaultConfig()).\n  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\nsomeLogger.Info("some log", \n  logger.LogFields{"adHoc-A": "value-A", "adHoc-B": "value-B"},\n  logger.LogFields{"adHoc-A": "new value"},\n)\n/*\n  {\n    "msg": "some log",\n    "lvl": 4,\n    "adHoc-A": "new value",\n    "adHoc-B": "value-B"\n  }\n*/\n')),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," are very suitable to log values that are different for every created log, like the ID of some user, for example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'someLogger := logger.New(logger.DefaultConfig()).\n  Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\nsomeLogger.Info("User created", logger.LogFields{"id": 556})\n/*\n  {\n    "msg": "User created",\n    "lvl": 4,\n    "id": 556\n  }\n*/\n')))}u.isMDXComponent=!0}}]);