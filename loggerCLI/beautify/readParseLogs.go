// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-29 4:32 PM

package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type log struct {
	Msg       string
	Lvl       uint64
	Timestamp string
	Parent    string
}

func readParseLogs() []log {
	scanner := bufio.NewScanner(os.Stdin)
	var logs []log
	for scanner.Scan() {
		var log_ log
		e := json.Unmarshal(scanner.Bytes(), &log_)
		if e != nil {
			panic(e)
		}
		logs = append(logs, log_)
	}

	return logs
}
