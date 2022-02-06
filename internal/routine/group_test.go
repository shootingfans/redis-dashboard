package routine

import (
	"sync/atomic"
	"testing"
)

func TestNewGroup(t *testing.T) {
	rgp := NewGroup()
	var num int32
	for range [10][]int{} {
		rgp.Add(func() {
			atomic.AddInt32(&num, 1)
		})
	}
	rgp.Add(func() {
		atomic.AddInt32(&num, 10)
		panic("test")
	})
	rgp.Run()
}
