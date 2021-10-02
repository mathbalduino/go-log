import React from 'react';
import clsx from 'clsx';
import styles from './HomepageFeatures.module.css';
import CodeBlock from '@theme/CodeBlock';
import Link from '@docusaurus/Link';

const easyToUseCode = `package main

import "github.com/mathbalduino/go-log"

func main() {
  // Outputs to stdout
  yourLogger := logger.NewDefault()

  yourLogger.Info("New INFO log message")
  // [ INFO ] New INFO log message
  
  yourLogger.Error("New ERROR log message")
  // [ ERROR ] New ERROR log message
  // ...
}

`

const baseFieldsCode = `package main

import (
	"github.com/mathbalduino/go-log"
	"os"
)

func main() {
  // Outputs to stdout
  userLogger := logger.New(logger.DefaultConfig()).
    Fields(logger.LogFields{"module": "user"}).
    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))

  userLogger.Info("New log")
  // { "lvl": 4, "module": "user", "msg": "New log" }
}
`

const dynamicFieldsCode = `package main

import (
  "github.com/mathbalduino/go-log"
  "os"
  "time"
)

func main() {
  // Outputs to stdout
  userLogger := logger.New(logger.DefaultConfig()).
    Fields(logger.LogFields{"module": "user"}).
    PreHooks(logger.Hooks{
      "timestamp": func(logger.Log) interface{} {
        return time.Now().UnixNano()
      },
    }).
    Outputs(logger.OutputJsonToWriter(os.Stdout, nil))

  userLogger.Info("New log", logger.LogFields{"id": 98})
  // { "id": 98, "lvl": 4, "module": "user", 
  //    "msg": "New log", "timestamp": 1633182921043120309 }
}

`

const outputCode = `package main

import (
  "github.com/mathbalduino/go-log"
  "fmt"
)

// Parse log fields to JSON and send them to the cloud
func OutputToCloud(_ uint64, _ string, fields logger.LogFields) {
  parseToJsonAndSendToCloud(fields)
}

func OutputToStdout(lvl uint64, msg string, _ logger.LogFields) {
  fmt.Printf("LogLevel: %d | LogMsg: %s\\n", lvl, msg)
}

func main() {
  yourLogger := logger.New(logger.DefaultConfig()).
    Outputs(OutputToStdout, OutputToCloud)
  
  yourLogger.Info("New log") 
  // stdout: "LogLevel: 4 | LogMsg: New log\\n"
  // cloud: { "lvl": 4, "msg": "New log" }
}
`

const FeatureList = [
  {
    title: 'Easy to Use',
    header: <CodeBlock className="language-go">{easyToUseCode}</CodeBlock>,
    description: (
      <>
        Just create a new default Logger instance and start to focus on your
        application development
      </>
    ),
  },
  {
    title: 'Base Log Fields',
    header: <CodeBlock className="language-go">{baseFieldsCode}</CodeBlock>,
    description: (
      <>
        You can define how many <Link to='/docs/basic-concepts/base_fields'>Base fields </Link>
        as you want. These fields are constant values and will be used to compose every log
        created by the Logger instance
      </>
    ),
  },
  {
    title: 'Dynamic Log Fields',
    header: <CodeBlock className="language-go">{dynamicFieldsCode}</CodeBlock>,
    description: (
      <>
        Some fields my need to be calculated every time a new log is created. To solve
        this problem, you can use <Link to='/docs/basic-concepts/pre_hooks'>PreHooks</Link>,
        <Link to='/docs/basic-concepts/adhoc_fields'> AdHoc fields</Link> and
        <Link to='/docs/basic-concepts/post_hooks'> PostHooks</Link>
      </>
    ),
  },
  {
    title: 'Output configuration',
    header: (
      <CodeBlock className="language-go">{outputCode}</CodeBlock>
    ),
    description: (
      <>
        Every log has a destiny, and you can set as many <Link to='/docs/basic-concepts/outputs'>Outputs </Link>
        as you want. If you don't want to implement your own outputs, just use the builtins
      </>
    ),
  },
];

function Feature({header, title, description}) {
  return (
    <div className={clsx('col col--6')}>
      {header}
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
