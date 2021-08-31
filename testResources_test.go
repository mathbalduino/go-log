package logger

import (
	"sort"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	rContextWithCancel, wContextWithCancel = newResource()
	rAsyncHandleLog, wAsyncHandleLog       = newResource()
	rNewWaitGroup, wNewWaitGroup           = newResource()
	rAtomicAddUint64, wAtomicAddUint64     = newResource()
	rHandleLog, wHandleLog                 = newResource()
	rStdOut, wStdOut                       = newResource()
)

// Tests run in parallel, so it's required to control the concurrency over
// global variables (such as the mocked functions). The functions/types below
// handle it

type testResource struct {
	index  uint64 // unique ID
	isRead bool
	mutex  *sync.RWMutex
}

// generates a new mutex + unique ID
var newResource = func() func() (testResource, testResource) {
	i := uint64(0)
	return func() (testResource, testResource) {
		idx := atomic.AddUint64(&i, 1) - 1
		rwMutex := &sync.RWMutex{}
		return testResource{idx, true, rwMutex}, testResource{idx, false, rwMutex}
	}
}()

// sort the resources in ID ascending order, and call "Lock". "Unlock" in the reverse order
func raceFreeTest(fn func(*testing.T), resources ...testResource) func(*testing.T) {
	return func(t *testing.T) {
		sort.Slice(resources, func(i, j int) bool { return resources[i].index < resources[j].index })
		for _, resource := range resources {
			if resource.isRead {
				resource.mutex.RLock()
				defer resource.mutex.RUnlock() // Safe to call inside the loop
			} else {
				resource.mutex.Lock()
				defer resource.mutex.Unlock() // Safe to call inside the loop
			}
		}
		fn(t)
	}
}
