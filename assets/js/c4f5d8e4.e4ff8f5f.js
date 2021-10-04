"use strict";(self.webpackChunkmy_website=self.webpackChunkmy_website||[]).push([[195],{3667:function(e,n,t){t.r(n),t.d(n,{default:function(){return B}});var o=t(2263),g=t(6742),s=t(6698),a=t(6010),l=t(7294),r=t(7462),i="features_pzNA",c=t(6213),u=[{title:"Easy to Use",header:l.createElement(c.Z,{className:"language-go"},'package main\n\nimport "github.com/mathbalduino/go-log"\n\nfunc main() {\n  // Outputs to stdout\n  yourLogger := logger.NewDefault()\n\n  yourLogger.Info("New INFO log message")\n  // [ INFO ] New INFO log message\n  \n  yourLogger.Error("New ERROR log message")\n  // [ ERROR ] New ERROR log message\n  // ...\n}\n\n'),description:l.createElement(l.Fragment,null,"Just create a new default Logger instance and start to focus on your application development")},{title:"Base Log Fields",header:l.createElement(c.Z,{className:"language-go"},'package main\n\nimport (\n  "github.com/mathbalduino/go-log"\n  "os"\n)\n\nfunc main() {\n  // Outputs to stdout\n  userLogger := logger.New(logger.DefaultConfig()).\n    Fields(logger.LogFields{"module": "user"}).\n    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\n\n  userLogger.Info("New log")\n  // { "lvl": 4, "module": "user", "msg": "New log" }\n}\n'),description:l.createElement(l.Fragment,null,"You can define how many ",l.createElement(g.Z,{to:"/docs/basic-concepts/base_fields"},"Base fields "),"as you want. These fields are constant values and will be used to compose every log created by the Logger instance")},{title:"Dynamic Log Fields",header:l.createElement(c.Z,{className:"language-go"},'package main\n\nimport (\n  "github.com/mathbalduino/go-log"\n  "os"\n  "time"\n)\n\nfunc main() {\n  // Outputs to stdout\n  userLogger := logger.New(logger.DefaultConfig()).\n    Fields(logger.LogFields{"module": "user"}).\n    PreHooks(logger.Hooks{\n      "timestamp": func(logger.Log) interface{} {\n        return time.Now().UnixNano()\n      },\n    }).\n    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))\n\n  userLogger.Info("New log", logger.LogFields{"id": 98})\n  // { "id": 98, "lvl": 4, "module": "user", \n  //    "msg": "New log", "timestamp": 1633182921043120309 }\n}\n\n'),description:l.createElement(l.Fragment,null,"Some fields my need to be calculated every time a new log is created. To solve this problem, you can use ",l.createElement(g.Z,{to:"/docs/basic-concepts/pre_hooks"},"PreHooks"),",",l.createElement(g.Z,{to:"/docs/basic-concepts/adhoc_fields"}," AdHoc fields")," and",l.createElement(g.Z,{to:"/docs/basic-concepts/post_hooks"}," PostHooks"))},{title:"Output configuration",header:l.createElement(c.Z,{className:"language-go"},'package main\n\nimport (\n  "github.com/mathbalduino/go-log"\n  "fmt"\n)\n\n// Parse log fields to JSON and send them to the cloud\nfunc OutputToCloud(_ uint64, _ string, fields logger.LogFields) {\n  parseToJsonAndSendToCloud(fields)\n}\n\nfunc OutputToStdout(lvl uint64, msg string, _ logger.LogFields) {\n  fmt.Printf("LogLevel: %d | LogMsg: %s\\n", lvl, msg)\n}\n\nfunc main() {\n  yourLogger := logger.New(logger.DefaultConfig()).\n    Outputs(OutputToStdout, OutputToCloud)\n  \n  yourLogger.Info("New log") \n  // stdout: "LogLevel: 4 | LogMsg: New log\\n"\n  // cloud: { "lvl": 4, "msg": "New log" }\n}\n'),description:l.createElement(l.Fragment,null,"Every log has a destiny, and you can set as many ",l.createElement(g.Z,{to:"/docs/basic-concepts/outputs"},"Outputs "),"as you want. If you don't want to implement your own outputs, just use the builtins")}];function A(e){var n=e.header,t=e.title,o=e.description;return l.createElement("div",{className:(0,a.Z)("col col--6")},n,l.createElement("div",{className:"text--center padding-horiz--md"},l.createElement("h3",null,t),l.createElement("p",null,o)))}function d(){return l.createElement("section",{className:i},l.createElement("div",{className:"container"},l.createElement("div",{className:"row"},u.map((function(e,n){return l.createElement(A,(0,r.Z)({key:n},e))})))))}var m="heroBanner_1ZBZ",Q="buttons_irzW";function E(){var e=(0,o.Z)().siteConfig,n=t(1931).Z;return l.createElement("header",{className:(0,a.Z)("hero hero--primary",m)},l.createElement("div",{className:"container"},l.createElement("img",{src:n}),l.createElement("h1",{className:"hero__title"},e.title),l.createElement("p",{className:"hero__subtitle"},e.tagline),l.createElement("div",{className:Q},l.createElement(g.Z,{className:"button button--secondary button--lg",to:"/docs/intro"},"Read the Introduction \ud83d\udcdc"))))}function B(){return l.createElement(s.Z,{description:"go-log library online documentation"},l.createElement(E,null),l.createElement("main",null,l.createElement(d,null)))}},1931:function(e,n){n.Z="data:image/png;base64,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"}}]);