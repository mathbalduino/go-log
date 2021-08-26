package loxeLog

import (
	"sort"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	rContextWithCancel = newResource()
	rAsyncHandleLog    = newResource()
	rNewWaitGroup      = newResource()
	rAtomicAddUint64   = newResource()
	rHandleLog         = newResource()
)

// Tests run in parallel, so it's required to control the concurrency over
// global variables (such as the mocked functions). The functions/types below
// handle it

type testResource struct {
	index uint64 // unique ID
	mutex *sync.Mutex
}

// generates a new mutex + unique ID
var newResource = func() func() testResource {
	i := uint64(0)
	return func() testResource {
		return testResource{
			atomic.AddUint64(&i, 1) - 1,
			&sync.Mutex{},
		}
	}
}()

// sort the resources in ID ascending order, and call "Lock". "Unlock" in the reverse order
func raceFreeTest(fn func(*testing.T), resources ...testResource) func(*testing.T) {
	return func(t *testing.T) {
		sort.Slice(resources, func(i, j int) bool { return resources[i].index < resources[j].index })
		for _, resource := range resources {
			resource.mutex.Lock()
			defer resource.mutex.Unlock() // Safe to call inside the loop
		}
		fn(t)
	}
}
