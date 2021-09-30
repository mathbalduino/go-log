"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[736],{3905:function(e,t,n){n.d(t,{Zo:function(){return s},kt:function(){return m}});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function u(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),p=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},s=function(e){var t=p(e.components);return r.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},c=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,l=e.parentName,s=u(e,["components","mdxType","originalType","parentName"]),c=p(n),m=a,g=c["".concat(l,".").concat(m)]||c[m]||d[m]||o;return n?r.createElement(g,i(i({ref:t},s),{},{components:n})):r.createElement(g,i({ref:t},s))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=c;var u={};for(var l in t)hasOwnProperty.call(t,l)&&(u[l]=t[l]);u.originalType=e,u.mdxType="string"==typeof e?e:a,i[1]=u;for(var p=2;p<o;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}c.displayName="MDXCreateElement"},7733:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return u},contentTitle:function(){return l},metadata:function(){return p},toc:function(){return s},default:function(){return c}});var r=n(7462),a=n(3366),o=(n(7294),n(3905)),i=["components"],u={sidebar_position:7},l="Outputs",p={unversionedId:"basic-concepts/outputs",id:"basic-concepts/outputs",isDocsHomePage:!1,title:"Outputs",description:"Every log needs to be written to somewhere, otherwise why would we create it? In this library, Output functions represent this final destiny:",source:"@site/docs/basic-concepts/outputs.md",sourceDirName:"basic-concepts",slug:"/basic-concepts/outputs",permalink:"/go-log/docs/basic-concepts/outputs",editUrl:"https://github.com/mathbalduino/go-log/edit/main/docs/docs/basic-concepts/outputs.md",tags:[],version:"current",sidebarPosition:7,frontMatter:{sidebar_position:7},sidebar:"tutorialSidebar",previous:{title:"Fields override order",permalink:"/go-log/docs/basic-concepts/override_order"},next:{title:"Life cycle",permalink:"/go-log/docs/basic-concepts/life_cycle"}},s=[{value:"Defining Outputs",id:"defining-outputs",children:[]},{value:"Ordering",id:"ordering",children:[]},{value:"Builtin outputs",id:"builtin-outputs",children:[{value:"OutputToWriter",id:"outputtowriter",children:[]},{value:"OutputJsonToWriter",id:"outputjsontowriter",children:[]},{value:"OutputAnsiToStdout",id:"outputansitostdout",children:[]},{value:"OutputPanicOnFatal",id:"outputpaniconfatal",children:[]}]}],d={toc:s};function c(e){var t=e.components,n=(0,a.Z)(e,i);return(0,o.kt)("wrapper",(0,r.Z)({},d,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"outputs"},"Outputs"),(0,o.kt)("p",null,"Every log needs to be written to somewhere, otherwise why would we create it? In this library, ",(0,o.kt)("inlineCode",{parentName:"p"},"Output")," functions represent this final destiny:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\n// Just an alias to a simple function\ntype Output = func(lvl uint64, msg string, fields LogFields)\n")),(0,o.kt)("p",null,"Inside these functions, you can do whatever you want. In general, it will be a write operation to the filesystem, database, or some request to the cloud."),(0,o.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"note")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"Even with the ",(0,o.kt)("inlineCode",{parentName:"p"},"lvl")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"msg")," being directly accessible via the function parameters, they're inside the ",(0,o.kt)("inlineCode",{parentName:"p"},"fields")," map param too, so you can just parse the ",(0,o.kt)("inlineCode",{parentName:"p"},"LogFields")," map right away."))),(0,o.kt)("div",{className:"admonition admonition-danger alert alert--danger"},(0,o.kt)("div",{parentName:"div",className:"admonition-heading"},(0,o.kt)("h5",{parentName:"div"},(0,o.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,o.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"12",height:"16",viewBox:"0 0 12 16"},(0,o.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"}))),"danger")),(0,o.kt)("div",{parentName:"div",className:"admonition-content"},(0,o.kt)("p",{parentName:"div"},"You have to be very carefull when writing ",(0,o.kt)("inlineCode",{parentName:"p"},"Outputs"),", because the library is not prepared to handle any kind of ",(0,o.kt)("inlineCode",{parentName:"p"},"panic")," that can occur inside them, and it can cause issues. "),(0,o.kt)("p",{parentName:"div"},(0,o.kt)("em",{parentName:"p"},(0,o.kt)("strong",{parentName:"em"},"Except if it's the last ",(0,o.kt)("inlineCode",{parentName:"strong"},"Output")," (last thing to be executed), avoid ",(0,o.kt)("inlineCode",{parentName:"strong"},"panic")," calls inside ",(0,o.kt)("inlineCode",{parentName:"strong"},"Outputs")))))),(0,o.kt)("h2",{id:"defining-outputs"},"Defining Outputs"),(0,o.kt)("p",null,"Just like ",(0,o.kt)("inlineCode",{parentName:"p"},"Hooks"),", you can set a new ",(0,o.kt)("inlineCode",{parentName:"p"},"Output")," using the ",(0,o.kt)("inlineCode",{parentName:"p"},"Outputs"),"/",(0,o.kt)("inlineCode",{parentName:"p"},"RawOutputs")," methods:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\nfunc (l *Logger) Outputs(output Output, outputs ...Output) *Logger { ... }\nfunc (l *Logger) RawOutputs(output Output, outputs ...Output) *Logger { ... }\n")),(0,o.kt)("p",null,"Note that it's a variadic function, so you can pass as many outputs as you want. The order will be preserved."),(0,o.kt)("p",null,"The ",(0,o.kt)("inlineCode",{parentName:"p"},"Outputs")," will ",(0,o.kt)("strong",{parentName:"p"},"append")," the new outputs to the old ones, while the ",(0,o.kt)("inlineCode",{parentName:"p"},"RawOutputs")," will ignore the old ones and use just the new outputs."),(0,o.kt)("h2",{id:"ordering"},"Ordering"),(0,o.kt)("p",null,"Since these functions are stored as an slice inside the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger"),", the order can be preserved and you can use it at you benefit. You can save it first to the database A, and, in the next ",(0,o.kt)("inlineCode",{parentName:"p"},"Output"),", read from database A, process, and save to database B, for example."),(0,o.kt)("p",null,"At the end of the life cycle of every created log, there's a for loop that will iterate over the ",(0,o.kt)("inlineCode",{parentName:"p"},"Outputs")," slice:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// Not real production-code (just to illustrate)\nfor _, output := range logger.outputs {\n  output(lvl, msg, logFields)\n}\n")),(0,o.kt)("p",null,"You must handle possible ",(0,o.kt)("inlineCode",{parentName:"p"},"panic")," calls that may occur inside the outputs, because it will be not handled by the library."),(0,o.kt)("h2",{id:"builtin-outputs"},"Builtin outputs"),(0,o.kt)("p",null,"There's 4 builtin outputs, ready to be used:"),(0,o.kt)("h3",{id:"outputtowriter"},"OutputToWriter"),(0,o.kt)("p",null,"Writes the log to some ",(0,o.kt)("inlineCode",{parentName:"p"},"io.Writer")," (usually, a file), after being parsed using the ",(0,o.kt)("inlineCode",{parentName:"p"},"OutputParser"),":"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\ntype OutputParser = func(LogFields) ([]byte, error) // just an alias\nfunc OutputToWriter(w io.Writer, parser OutputParser, onError func(error)) Output {\n  ...\n}\n")),(0,o.kt)("p",null,"Note that there's a thirty argument: ",(0,o.kt)("inlineCode",{parentName:"p"},"onError"),". It is used to handle possible errors when trying to parse the log using the ",(0,o.kt)("inlineCode",{parentName:"p"},"OutputParser")," or trying to write it to the ",(0,o.kt)("inlineCode",{parentName:"p"},"io.Writer"),". It is intended to be used as a last fallback."),(0,o.kt)("p",null,"This is, in fact, a function that will return another function. Note that if you pass it directly to the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger"),", the compiler will stop you. The returned function is the real ",(0,o.kt)("inlineCode",{parentName:"p"},"Output"),", pay attention. Example:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// compiler error\nlogger.NewDefault().\n  Outputs(OutputToWriter)\n\n// ok\nlogger.NewDefault().\n  Outputs(OutputToWriter(w, p, func(error) {}))\n")),(0,o.kt)("h3",{id:"outputjsontowriter"},"OutputJsonToWriter"),(0,o.kt)("p",null,"Writes the log to some ",(0,o.kt)("inlineCode",{parentName:"p"},"io.Writer")," (usually, a file), after being parsed to ",(0,o.kt)("inlineCode",{parentName:"p"},"json"),":"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\nfunc OutputJsonToWriter(w io.Writer, onError func(error)) Output {\n  ...\n}\n")),(0,o.kt)("p",null,"Note that the ",(0,o.kt)("inlineCode",{parentName:"p"},"onError")," argument has the same purpose as the one in ",(0,o.kt)("a",{parentName:"p",href:"#outputtowriter"},"OutputToWriter"),"."),(0,o.kt)("p",null,"This is, in fact, a function that will return another function. Note that if you pass it directly to the ",(0,o.kt)("inlineCode",{parentName:"p"},"Logger"),", the compiler will stop you. The returned function is the real ",(0,o.kt)("inlineCode",{parentName:"p"},"Output"),", pay attention."),(0,o.kt)("h3",{id:"outputansitostdout"},"OutputAnsiToStdout"),(0,o.kt)("p",null,"Writes the log to the ",(0,o.kt)("inlineCode",{parentName:"p"},"stdout"),", displaying just the ",(0,o.kt)("inlineCode",{parentName:"p"},"level")," and the ",(0,o.kt)("inlineCode",{parentName:"p"},"message"),", using ",(0,o.kt)("inlineCode",{parentName:"p"},"ANSI")," codes to colorize it accordingly to it's ",(0,o.kt)("inlineCode",{parentName:"p"},"level"),". If your ",(0,o.kt)("inlineCode",{parentName:"p"},"stdout")," don't have support for ",(0,o.kt)("inlineCode",{parentName:"p"},"ANSI")," codes, don't use this ",(0,o.kt)("inlineCode",{parentName:"p"},"Output")," (not common, since in general it will be some terminal)."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\nfunc OutputAnsiToStdout(lvl uint64, msg string, _ LogFields) {\n  ...\n}\n")),(0,o.kt)("h3",{id:"outputpaniconfatal"},"OutputPanicOnFatal"),(0,o.kt)("p",null,"As you will see, the ",(0,o.kt)("inlineCode",{parentName:"p"},"Fatal")," log level doesn't do anything special. In order to unlock it's ability to ",(0,o.kt)("inlineCode",{parentName:"p"},"panic"),", you will need to use this special ",(0,o.kt)("inlineCode",{parentName:"p"},"Output"),":"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\nfunc OutputPanicOnFatal(lvl uint64, msg string, fields LogFields) {\n  ...\n}\n")),(0,o.kt)("p",null,"Just set it to be the last ",(0,o.kt)("inlineCode",{parentName:"p"},"Output")," and it will call ",(0,o.kt)("inlineCode",{parentName:"p"},"panic")," if the received log is a Fatal one. "),(0,o.kt)("p",null,"If there's some ",(0,o.kt)("inlineCode",{parentName:"p"},"error")," value inside the ",(0,o.kt)("inlineCode",{parentName:"p"},"LogFields"),", it will be given to the ",(0,o.kt)("inlineCode",{parentName:"p"},"panic")," call, otherwise, a new ",(0,o.kt)("inlineCode",{parentName:"p"},"error")," will be created using the ",(0,o.kt)("inlineCode",{parentName:"p"},"msg")," argument and ",(0,o.kt)("inlineCode",{parentName:"p"},"fmt.Errorf()"),"."),(0,o.kt)("p",null,"This ",(0,o.kt)("inlineCode",{parentName:"p"},"Output")," will search for the error value inside the log fields using the ",(0,o.kt)("inlineCode",{parentName:"p"},"DefaultErrorKey")," key:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'// configuration.go\nfunc DefaultErrorParser(err error) (string, LogFields) {\n    return err.Error(), LogFields{DefaultErrorKey: err}\n}\nconst DefaultErrorKey = "error"\n')),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},"// outputs.go\nfunc OutputPanicOnFatal(lvl uint64, msg string, fields LogFields) {\n  ...\n  err := fields[DefaultErrorKey]\n  ...\n}\n")),(0,o.kt)("p",null,"If you're using a different ",(0,o.kt)("inlineCode",{parentName:"p"},"ErrorParser"),", other than the ",(0,o.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/log_levels#default-errorparser"},"DefaultErrorParser"),", make sure that the error value is stored inside the ",(0,o.kt)("inlineCode",{parentName:"p"},"LogFields")," under the ",(0,o.kt)("inlineCode",{parentName:"p"},"DefaultErrorKey")," key. ",(0,o.kt)("del",{parentName:"p"},"Or not, it's up to you")))}c.isMDXComponent=!0}}]);