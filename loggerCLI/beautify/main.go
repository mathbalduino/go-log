package main

import (
	"fmt"
	logger "github.com/mathbalduino/go-log"
	"github.com/mathbalduino/go-log/internal"
)

func main() {
	fmt.Println(logger.ColorizeStrByLvl(logger.LvlInfo, signatureTmpl))

	logs := readParseLogs()
	orderLogs(logs)
	tree := buildTree(logs)
	treeStr := treeToString(tree, 0)
	fmt.Println(treeStr)
}

// signatureTmpl is the header containing information
// about the author and library
const signatureTmpl = `||
|| ` + internal.LibraryName + ` ` + internal.LibraryModuleVersion + ` - beautify
|| by Matheus Leonel Balduino
||
|| Everywhere, under @mathbalduino:
||   GitLab:    @mathbalduino
||   Instagram: @mathbalduino
||   Twitter:   @mathbalduino
||   Website:   mathbalduino.com.br/` + internal.LibraryName + `
||
`
