package loxeLog

import "fmt"

// ErrLvlMsgSameKey is an error token used to represent the
// situation where the given Configuration struct contains
// the 'LvlFieldName' and 'MsgFieldName' fields with the same value
var ErrLvlMsgSameKey = fmt.Errorf("the 'level' and 'message' keys cannot be equal")

// ErrNilWaitGroup is threw by AsyncHandleLog when it receives
// a nil WaitGroup via the function argument
var ErrNilWaitGroup = fmt.Errorf("the 'WaitGroup' argument cannot be nil")

// ErrNilCtx is threw by AsyncHandleLog when it receives
// a nil context via the function argument
var ErrNilCtx = fmt.Errorf("the 'context.Context' argument cannot be nil")

// ErrNilChan is threw by AsyncHandleLog when it receives
// a nil channel via the function argument
var ErrNilChan = fmt.Errorf("the 'channel' argument cannot be nil")
