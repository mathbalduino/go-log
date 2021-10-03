"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[937],{3905:function(e,n,t){t.d(n,{Zo:function(){return c},kt:function(){return p}});var r=t(7294);function o(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function a(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function l(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?a(Object(t),!0).forEach((function(n){o(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):a(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function i(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var s=r.createContext({}),g=function(e){var n=r.useContext(s),t=n;return e&&(t="function"==typeof e?e(n):l(l({},n),e)),t},c=function(e){var n=g(e.components);return r.createElement(s.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},u=r.forwardRef((function(e,n){var t=e.components,o=e.mdxType,a=e.originalType,s=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),u=g(t),p=o,m=u["".concat(s,".").concat(p)]||u[p]||d[p]||a;return t?r.createElement(m,l(l({ref:n},c),{},{components:t})):r.createElement(m,l({ref:n},c))}));function p(e,n){var t=arguments,o=n&&n.mdxType;if("string"==typeof e||o){var a=t.length,l=new Array(a);l[0]=u;var i={};for(var s in n)hasOwnProperty.call(n,s)&&(i[s]=n[s]);i.originalType=e,i.mdxType="string"==typeof e?e:o,l[1]=i;for(var g=2;g<a;g++)l[g]=t[g];return r.createElement.apply(null,l)}return r.createElement.apply(null,t)}u.displayName="MDXCreateElement"},5074:function(e,n,t){t.r(n),t.d(n,{frontMatter:function(){return i},contentTitle:function(){return s},metadata:function(){return g},toc:function(){return c},default:function(){return u}});var r=t(7462),o=t(3366),a=(t(7294),t(3905)),l=["components"],i={sidebar_position:9},s="Log levels",g={unversionedId:"basic-concepts/log_levels",id:"basic-concepts/log_levels",isDocsHomePage:!1,title:"Log levels",description:"Basic log levels",source:"@site/docs/basic-concepts/log_levels.md",sourceDirName:"basic-concepts",slug:"/basic-concepts/log_levels",permalink:"/go-log/docs/basic-concepts/log_levels",editUrl:"https://github.com/mathbalduino/go-log/edit/main/docs/docs/basic-concepts/log_levels.md",tags:[],version:"current",sidebarPosition:9,frontMatter:{sidebar_position:9},sidebar:"tutorialSidebar",previous:{title:"Life cycle",permalink:"/go-log/docs/basic-concepts/life_cycle"},next:{title:"Async Logger",permalink:"/go-log/docs/basic-concepts/async_logger"}},c=[{value:"Basic log levels",id:"basic-log-levels",children:[]},{value:"Extra Log Levels",id:"extra-log-levels",children:[{value:"ErrorParser",id:"errorparser",children:[]},{value:"Default ErrorParser",id:"default-errorparser",children:[]}]},{value:"Standalone log levels",id:"standalone-log-levels",children:[]},{value:"Extending the log levels",id:"extending-the-log-levels",children:[]}],d={toc:c};function u(e){var n=e.components,t=(0,o.Z)(e,l);return(0,a.kt)("wrapper",(0,r.Z)({},d,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"log-levels"},"Log levels"),(0,a.kt)("h2",{id:"basic-log-levels"},"Basic log levels"),(0,a.kt)("p",null,"There are 6 basic log levels implemented by default:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"// logLevels.go\nfunc (l *Logger) Trace(msg string, adHocFields ...LogFields) { ... }\nfunc (l *Logger) Debug(msg string, adHocFields ...LogFields) { ... }\nfunc (l *Logger) Info(msg string, adHocFields ...LogFields) { ... }\nfunc (l *Logger) Warn(msg string, adHocFields ...LogFields) { ... }\nfunc (l *Logger) Error(msg string, adHocFields ...LogFields) { ... }\nfunc (l *Logger) Fatal(msg string, adHocFields ...LogFields) { ... }\n")),(0,a.kt)("p",null,"In order to call these methods, you will need a valid ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," instance. See ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/logger_creation"},"Logger creation")," for details."),(0,a.kt)("p",null,"The only difference between them is the value of the ",(0,a.kt)("inlineCode",{parentName:"p"},"lvl")," log field. Note that even if it's common to ",(0,a.kt)("inlineCode",{parentName:"p"},"panic")," when calling ",(0,a.kt)("inlineCode",{parentName:"p"},"Fatal"),", you will have to implement some ",(0,a.kt)("inlineCode",{parentName:"p"},"Output")," to simulate this behaviour (see ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/outputs#outputpaniconfatal"},"OutputPanicOnFatal")," for details)."),(0,a.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,a.kt)("div",{parentName:"div",className:"admonition-heading"},(0,a.kt)("h5",{parentName:"div"},(0,a.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,a.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,a.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"note")),(0,a.kt)("div",{parentName:"div",className:"admonition-content"},(0,a.kt)("p",{parentName:"div"},"The ",(0,a.kt)("inlineCode",{parentName:"p"},"lvl")," of the log is represented using an ",(0,a.kt)("inlineCode",{parentName:"p"},"uint64")," and values that correspond to the power of two (1, 2, 4, 8, 16, ...). This way, it's possible to apply the ",(0,a.kt)("inlineCode",{parentName:"p"},"and")," bitwise operation and easily check to see if some log level X is enabled. For details, see ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/configuration#lvlsenabled-usage"},"LevelsEnabled usage")))),(0,a.kt)("h2",{id:"extra-log-levels"},"Extra Log Levels"),(0,a.kt)("p",null,"If you've read the ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/intro"},"Introduction"),", you may have noticed these two log levels called ",(0,a.kt)("inlineCode",{parentName:"p"},"ErrorFrom")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"FatalFrom"),"."),(0,a.kt)("p",null,"These two methods are another way to call ",(0,a.kt)("inlineCode",{parentName:"p"},"Error")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"Fatal"),", respectively, but using an ",(0,a.kt)("inlineCode",{parentName:"p"},"error")," interface directly."),(0,a.kt)("h3",{id:"errorparser"},"ErrorParser"),(0,a.kt)("p",null,"If you're using a concise ",(0,a.kt)("inlineCode",{parentName:"p"},"error")," handling strategy, it should be easy to manipulate errors and\nextract more information from them, beyond the error string."),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"Configuration")," struct accepts a function called ",(0,a.kt)("inlineCode",{parentName:"p"},"ErrorParser"),", that will take some ",(0,a.kt)("inlineCode",{parentName:"p"},"error")," interface and return a tuple containing the string that better represents the error, and some log fields extracted from the error itself:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"// configuration.go\ntype Configuration struct {\n  ...\n  ErrorParser func(error) (string, LogFields)\n  ...\n}\n")),(0,a.kt)("p",null,"This function will be called from both ",(0,a.kt)("inlineCode",{parentName:"p"},"ErrorFrom")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"FatalFrom")," methods, in order to extract more information about the error, to finally call the correspondent log method (",(0,a.kt)("inlineCode",{parentName:"p"},"Error")," and",(0,a.kt)("inlineCode",{parentName:"p"},"Fatal"),", respectively)."),(0,a.kt)("p",null,"Instead of doing something like this:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'type SomeErrorStruct struct{\n  Value int\n}\nfunc (s *SomeErrorStruct) Error() string { return "error msg" }\n\nfunc maybeReturnsError() error { return &SomeErrorStruct{Value: 10} }\n\nfunc main() {\n  someLogger := logger.NewDefault()\n\n  e := maybeReturnsError()\n  if e != nil {\n    err := e.(*SomeErrorStruct)\n    someLogger.Fatal(e.Error(), logger.LogFields{ "errorValue": err.Value })\n  }\n}\n')),(0,a.kt)("p",null,"You can do this:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'type SomeErrorStruct struct{\n  Value int\n}\nfunc (s *SomeErrorStruct) Error() string { return "error msg" }\n\nfunc maybeReturnsError() error { return &SomeErrorStruct{Value: 10} }\n\nfunc MyErrorParser(e error) (string, logger.LogFields) {\n  return e.Error(), logger.LogFields{\n      "errorValue": e.(*SomeErrorStruct).Value,\n  }\n}\n\nfunc main() {\n  config := logger.DefaultConfig()\n  config.ErrorParser = MyErrorParser\n  someLogger := logger.New(config).\n    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\n\n  e := maybeReturnsError()\n  if e != nil {\n    someLogger.FatalFrom(e)\n    // { "errorValue": 10,"lvl": 32, "msg": "error msg" }\n  }\n}\n')),(0,a.kt)("p",null,"Note that the fields returned by the ",(0,a.kt)("inlineCode",{parentName:"p"},"ErrorParser")," will be placed ",(0,a.kt)("strong",{parentName:"p"},"before")," the ",(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields"),", causing them to be overwritten by the ",(0,a.kt)("inlineCode",{parentName:"p"},"AdHoc fields")," if there's a clash of keys. Example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'// Suppose that the ErrorParser is set to:\nfunc MyErrorParser(_ error) (string, logger.LogFields) {\n  return "some example error msg", logger.LogFields{ \n    "someKey": "value", \n    "anotherKey": "another value",\n  }\n}\n\nfunc main() {\n  config := logger.DefaultConfig()\n  config.ErrorParser = MyErrorParser\n  someLogger := logger.New(config).\n    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\n\n  // And you create the following log:\n  someLogger.ErrorFrom(fmt.Errorf("any error"), logger.LogFields{\n    "someKey": "newValue",\n    "thirtyKey": "another another value",\n  })\n  /*\n    {\n      "msg": "some example error msg",\n      "lvl": 16,\n      "someKey": "newValue",\n      "anotherKey": "another value",\n      "thirtyKey": "another another value"\n    }\n  */\n}\n')),(0,a.kt)("h3",{id:"default-errorparser"},"Default ErrorParser"),(0,a.kt)("p",null,"The library has a builtin ",(0,a.kt)("inlineCode",{parentName:"p"},"ErrorParser"),", that will return the ",(0,a.kt)("inlineCode",{parentName:"p"},"e.Error()")," as the msg and store the ",(0,a.kt)("inlineCode",{parentName:"p"},"error")," interface value inside the log fields, using the ",(0,a.kt)("inlineCode",{parentName:"p"},"DefaultErrorKey")," as the field key:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'// configuration.go\n// Production code-fragment\nfunc DefaultErrorParser(err error) (string, LogFields) {\n    return err.Error(), LogFields{DefaultErrorKey: err}\n}\nconst DefaultErrorKey = "error"\n')),(0,a.kt)("p",null,"This logic is used to build the ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/outputs#outputpaniconfatal"},"OutputPanicOnFatal"),", that extracts the error value (under the ",(0,a.kt)("inlineCode",{parentName:"p"},"DefaultErrorKey")," key) to forward it to the ",(0,a.kt)("inlineCode",{parentName:"p"},"panic")," call."),(0,a.kt)("h2",{id:"standalone-log-levels"},"Standalone log levels"),(0,a.kt)("p",null,"If you don't want to create and maintain a ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," instance by yourself, you can just use the standalone version of the log levels:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"// standalone.go\nfunc Trace(msg string, adHocFields ...LogFields) { ... }\nfunc Debug(msg string, adHocFields ...LogFields) { ... }\nfunc Info(msg string, adHocFields ...LogFields) { ... }\nfunc Warn(msg string, adHocFields ...LogFields) { ... }\nfunc Error(msg string, adHocFields ...LogFields) { ... }\nfunc Fatal(msg string, adHocFields ...LogFields) { ... }\n")),(0,a.kt)("p",null,"Under the hood, they will use a default ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," instance. Something like this:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"func Trace(msg string, adHocFields ...LogFields) {\n  NewDefault().Trace(msg, adHocFields...)\n}\n")),(0,a.kt)("p",null,"For more information, see ",(0,a.kt)("a",{parentName:"p",href:"/go-log/docs/basic-concepts/logger_creation"},"Default Logger"),"."),(0,a.kt)("h2",{id:"extending-the-log-levels"},"Extending the log levels"),(0,a.kt)("p",null,"If you want to create new custom log levels, you will need to create a new type that wrap the original ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger"),"."),(0,a.kt)("p",null,"The most straightway to do it, is to use embedded struct fields. Something like this:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type NewLogger struct {\n  *logger.Logger\n}\n")),(0,a.kt)("p",null,"This way, the ",(0,a.kt)("inlineCode",{parentName:"p"},"NewLogger")," will preserve the original log level methods and you can add new ones. Every log level should call the base log level method, that will actually handle the log. This method is exported by the ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," api:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"// logLevels.go\nfunc (l *Logger) Log(lvl uint64, msg string, adHocFields []LogFields) { ... }\n")),(0,a.kt)("p",null,"You can notice that the ",(0,a.kt)("inlineCode",{parentName:"p"},"adHoc")," fields are represented as an slice, while the log level methods (",(0,a.kt)("inlineCode",{parentName:"p"},"Fatal"),", ",(0,a.kt)("inlineCode",{parentName:"p"},"Info"),", etc) implement it as a variadic. Well, variadic arguments are just a slice, so if you want to create custom log levels using variadics too, you can just forward it to the ",(0,a.kt)("inlineCode",{parentName:"p"},"Log")," method:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type NewLogger struct {\n  *logger.Logger\n}\n\nfunc (l *NewLogger) NewLogLevel(msg string, adHocFields ...logger.LogFields) {\n  l.Logger.Log(<newLogLevelUint64>, msg, adHocFields)\n}\n")),(0,a.kt)("p",null,"Remember that it's not recommended to overlap new, custom log levels values (",(0,a.kt)("inlineCode",{parentName:"p"},"uint64"),"), with the builtin ones. They're all exported, so you can just check them or do something like this:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"const (\n  newLogLevel uint64 = 1 << (iota + 6)\n  anotherLogLevel\n  anotherOne\n)\n")),(0,a.kt)("p",null,"Note that it will continue the sequence after the last builtin log level, ",(0,a.kt)("inlineCode",{parentName:"p"},"LvlFatal"),". There's six builtin log levels (Trace, Debug, Info, Warn, Error and Fatal), that's why there's a ",(0,a.kt)("inlineCode",{parentName:"p"},"6")," being added to the ",(0,a.kt)("inlineCode",{parentName:"p"},"iota"),"."),(0,a.kt)("p",null,"If you want to omit some builtin log levels, just don't use embedded struct fields:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type NewLogger struct {\n  original *logger.Logger\n}\n\nfunc (l *NewLogger) Info(msg string, adHocFields ...logger.LogFields) {\n  l.original.Info(msg, adHocFields...)\n}\nfunc (l *NewLogger) Error(msg string, adHocFields ...logger.LogFields) {\n  l.original.Error(msg, adHocFields...)\n}\n")),(0,a.kt)("p",null,"This way, there will be only two methods in the new custom ",(0,a.kt)("inlineCode",{parentName:"p"},"Logger")," api: ",(0,a.kt)("inlineCode",{parentName:"p"},"Info")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"Error"),"."),(0,a.kt)("p",null,"You can see a more concrete example by looking at the ",(0,a.kt)("inlineCode",{parentName:"p"},"LoggerCLI"),", ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/mathbalduino/go-log/blob/main/loggerCLI/logLevels.go"},"here")," and ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/mathbalduino/go-log/blob/79263810d94dd1f2d112727824d1c5256b27951b/loggerCLI/new.go#L9"},"here"),"."))}u.isMDXComponent=!0}}]);