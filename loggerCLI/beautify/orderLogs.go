// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-29 4:38 PM

package main

import (
	"sort"
	"strconv"
)

// side-effect
func orderLogs(logs []log) {
	sort.Slice(logs, func(i_, j_ int) bool {
		i, eI := strconv.ParseInt(logs[i_].Timestamp, 10, 64)
		if eI != nil {
			panic(eI)
		}
		j, eJ := strconv.ParseInt(logs[j_].Timestamp, 10, 64)
		if eJ != nil {
			panic(eI)
		}

		return i < j
	})
}
