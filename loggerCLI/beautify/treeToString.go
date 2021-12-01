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

func treeToString(tree []treeNode, indentation string) string {
	finalStrB := strings.Builder{}
	for i, log := range tree {
		isLast := i == len(tree)-1

		finalStrB.WriteString(indentation)
		if log.parentPtr != nil {
			if isLast {
				finalStrB.WriteString(logger.ColorizeStrByLvl(log.parentPtr.Lvl, "'--"))
			} else {
				finalStrB.WriteString(logger.ColorizeStrByLvl(log.parentPtr.Lvl, "|--"))
			}
		}

		lvlPrefix := fmt.Sprintf("[ %s ] ", logger.LvlToString(log.Lvl))
		finalStrB.WriteString(logger.ColorizeStrByLvl(log.Lvl, lvlPrefix))

		splitMsg := strings.Split(log.Msg, "\n")
		for i, m := range splitMsg {
			finalStrB.WriteString(logger.ColorizeStrByLvl(log.Lvl, m))
			if i < (len(splitMsg) - 1) {
				finalStrB.WriteString("\n")
				finalStrB.WriteString(indentation)
				if log.parentPtr != nil {
					if isLast {
						finalStrB.WriteString(logger.ColorizeStrByLvl(log.parentPtr.Lvl, "   "))
					} else {
						finalStrB.WriteString(logger.ColorizeStrByLvl(log.parentPtr.Lvl, "|  "))
					}
				}
				finalStrB.WriteString("   ")
			}
		}
		finalStrB.WriteString("\n")

		newIndentation := ""
		if log.parentPtr != nil {
			if isLast {
				newIndentation = indentation + logger.ColorizeStrByLvl(log.parentPtr.Lvl, "   ")
			} else {
				newIndentation = indentation + logger.ColorizeStrByLvl(log.parentPtr.Lvl, "|  ")
			}
		}
		childStr := treeToString(log.childs, newIndentation)
		finalStrB.WriteString(childStr)
	}
	return finalStrB.String()
}
