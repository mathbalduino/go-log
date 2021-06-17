package loxeLog

import "fmt"

var ErrIncompleteFileWrite = fmt.Errorf("file write failed. There are unwritten characters")
