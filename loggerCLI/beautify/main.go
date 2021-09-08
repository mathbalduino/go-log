package main

//go:generate sh -c "go run ../main/main.go | go run *.go"

import (
	"bufio"
	"encoding/json"
	"fmt"
	logger "gitlab.com/loxe-tools/go-log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	stdin := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		s := ""
		for scanner.Scan() {
			s += scanner.Text() + "\n"
		}
		stdin <- strings.TrimSuffix(s, "\n")
	}()

infiniteLoop:
	for {
		select {
		case str, ok := <-stdin:
			if !ok {
				break infiniteLoop
			} else {
				rawLogs := readLogs(str)
				orderLogs(rawLogs)
				final := logTree(rawLogs)
				fmt.Println(drawLogTree(final, 0))
			}
		case <-time.After(time.Second):
			return
		}
	}
}

type log struct {
	Msg       string
	Lvl       uint64
	Timestamp string
	Parent    string
}

func readLogs(input string) []log {
	jsonsStrings := strings.Split(input, "\n")
	logsFields := []log{}
	for _, jsonString := range jsonsStrings {
		var log log
		e := json.Unmarshal([]byte(jsonString), &log)
		if e != nil {
			panic(e)
		}
		logsFields = append(logsFields, log)
	}
	return logsFields
}

type ABC struct {
	log
	parentPtr *ABC
	childs    []ABC
}

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

func logTree(logs []log) []ABC {
	aux := map[string]*ABC{}
	final := make([]ABC, 0, len(logs))
	for _, log := range logs {
		timestamp := log.Timestamp
		parent := aux[log.Parent]
		if parent == nil {
			final = append(final, ABC{log, nil, nil})
			aux[timestamp] = &final[len(final)-1]
			continue
		}

		parent.childs = append(parent.childs, ABC{log, parent, nil})
		aux[timestamp] = &parent.childs[len(parent.childs)-1]
	}
	return final
}

func drawLogTree(tree []ABC, treeDepth int) string {
	str := ""
	for i, log := range tree {
		isLast := i == len(tree) - 1
		prefix := ""
		if treeDepth > 0 {
			prefix = strings.Repeat("   ", treeDepth - 1)

			if isLast {
				prefix += logger.ColorizeStrByLvl(log.parentPtr.Lvl, "'--")
			} else {
				prefix += logger.ColorizeStrByLvl(log.parentPtr.Lvl, "|--")
			}
		}

		c := "   "
		if len(log.childs) > 0 {
			c = "|  "
		}
		msg := strings.ReplaceAll(log.Msg, "\n",
			"\n" + strings.Repeat("   ", treeDepth) + c)

		str += fmt.Sprintf(
			prefix + logger.ColorizeStrByLvl(log.Lvl, "[ %s ] %s") + "\n",
			logger.LvlToString(log.Lvl),
			msg)
		nestedStrs := drawLogTree(log.childs, treeDepth + 1)
		if len(log.childs) > 1 {
			nestedStrs = strings.ReplaceAll(
				nestedStrs,
				"\n" + strings.Repeat("   ", treeDepth) + " ",
				"\n" + strings.Repeat("   ", treeDepth) + "|")
		}
		str += nestedStrs
	}
	return str
}
//func printLogTree(tree []ABC, treeDepth int) {
//	for _, log := range tree {
//		prefix := ""
//		if treeDepth > 0 {
//			parent := log.parent
//			prefix = logger.ColorizeStrByLvl(parent.Lvl, "|--")
//			for i := treeDepth - 1; i > 0; i-- {
//				parent = parent.parent
//				prefix = logger.ColorizeStrByLvl(parent.Lvl, "|  ") + prefix
//			}
//			prefix = logger.ColorizeStrByLvl(parent.Lvl, prefix)
//		}
//
//		fmt.Printf(
//			prefix + logger.ColorizeStrByLvl(log.Lvl, "[ %s ] %s") + "\n",
//			logger.LvlToString(log.Lvl),
//			strings.ReplaceAll(log.Msg, "\n", "\n" + strings.Repeat("|  ", treeDepth + 1)))
//		printLogTree(log.childs, treeDepth + 1)
//	}
//}
