// Matheus Leonel Balduino
// Everywhere, under @mathbalduino
//   @mathbalduino on GitHub
//   @mathbalduino on Instagram
//   @mathbalduino on Twitter
// Live at mathbalduino.com.br
// 2021-11-29 4:44 PM

package main

import (
	"fmt"
	logger "github.com/mathbalduino/go-log"
	"strings"
)

func treeToString(tree []treeNode, treeDepth int) string {
	finalStrB := strings.Builder{}
	for i, log := range tree {
		isLast := i == len(tree)-1

		// Builds visual indentation prefix
		if treeDepth > 0 {
			finalStrB.WriteString(strings.Repeat("   ", treeDepth-1))
			if isLast {
				finalStrB.WriteString(logger.ColorizeStrByLvl(log.parentPtr.Lvl, "'--"))
			} else {
				finalStrB.WriteString(logger.ColorizeStrByLvl(log.parentPtr.Lvl, "|--"))
			}
		}

		// Builds the actual prefix of the log level
		lvlPrefix := fmt.Sprintf("[ %s ] ", logger.LvlToString(log.Lvl))
		finalStrB.WriteString(logger.ColorizeStrByLvl(log.Lvl, lvlPrefix))

		// Message new lines indentation
		msgIndentation := strings.Repeat("   ", treeDepth)
		if msgIndentation != "" && isLast {
			msgIndentation = strings.TrimSuffix(msgIndentation, "   ") + "¨  "
		}
		if len(log.childs) > 0 {
			msgIndentation += logger.ColorizeStrByLvl(log.Lvl, "|  ")
		} else {
			msgIndentation += "   "
		}

		// Builds message
		msg := strings.Split(log.Msg, "\n")
		for i, m := range msg {
			finalStrB.WriteString(logger.ColorizeStrByLvl(log.Lvl, m))
			if i < (len(msg) - 1) {
				finalStrB.WriteString("\n")
				finalStrB.WriteString(msgIndentation)
			}
		}
		finalStrB.WriteString("\n")

		// Build the child strings
		childStr := treeToString(log.childs, treeDepth+1)
		if len(log.childs) > 1 {
			childStr = strings.ReplaceAll(
				childStr,
				"\n"+strings.Repeat("   ", treeDepth)+" ",
				"\n"+strings.Repeat("   ", treeDepth)+logger.ColorizeStrByLvl(log.Lvl, "|"))
		}
		if len(log.childs) > 0 {
			childStr = strings.ReplaceAll(childStr, "¨  ", "   ")
		}
		finalStrB.WriteString(childStr)
	}

	return finalStrB.String()
}
