package event

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	mgr := NewManager()
	var a, b, c int32
	mgr.FocusOn("event1", func(event string, data interface{}) {
		assert.Equal(t, event, "event1")
		atomic.AddInt32(&a, data.(int32))
		atomic.AddInt32(&c, data.(int32))
	})
	mgr.FocusOn("event2", func(event string, data interface{}) {
		assert.Equal(t, event, "event2")
		atomic.AddInt32(&b, data.(int32))
		atomic.AddInt32(&c, data.(int32))
	})
	wait1 := mgr.Trigger("event1", int32(3))
	wait2 := mgr.Trigger("event2", int32(-2))
	wait3 := mgr.Trigger("event3", nil)
	wait1()
	wait2()
	wait3()
	assert.Equal(t, a, int32(3))
	assert.Equal(t, b, int32(-2))
	assert.Equal(t, c, int32(1))
}

func BenchmarkTrigger(b *testing.B) {
	mgr := NewManager()
	mgr.FocusOn("event", func(event string, data interface{}) {})
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mgr.Trigger("event", nil)
	}
}

func ExampleNewManager() {
	mgr := NewManager()
	// focus some event
	mgr.FocusOn("some_event", func(event string, data interface{}) {
		// do somethings
		fmt.Println(event, data)
	})
	// trigger event happened
	wait := mgr.Trigger("some_event", "It was happened")
	// wait will block until all focus handlers finished
	wait()
	// result:
	// some_event It was happened
}
