package loxeLog

import "fmt"

// ErrIncompleteFileWrite is an error token used to represent
// the situation where the OS couldn't write all the given bytes
// to the given io.Writer interface
var ErrIncompleteFileWrite = fmt.Errorf("file write failed. There are file unwritten bytes")

// ErrEqLvlMsgFieldNames is an error token used to represent
// the situation where the given Configuration struct contains
// the LvlFieldName and MsgFieldName fields with the same value
var ErrEqLvlMsgFieldNames = fmt.Errorf("the name of the 'level' and 'message' log fields cannot be equal")
