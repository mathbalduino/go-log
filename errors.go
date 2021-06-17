package loxeLog

import "fmt"

var ErrIncompleteFileWrite = fmt.Errorf("File write failed. There are unwritten characters")
