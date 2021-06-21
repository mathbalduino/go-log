package loxeLog

import "fmt"

// ErrLvlMsgSameKey is an error token used to represent the
// situation where the given Configuration struct contains
// the 'LvlFieldName' and 'MsgFieldName' fields with the same value
var ErrLvlMsgSameKey = fmt.Errorf("the 'level' and 'message' keys cannot be equal")
