package routine

import (
	"fmt"
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

func ExampleNewGroup() {
	rgp := NewGroup()
	var b bool
	var c []int
	var sum int32
	// add some routine to group
	rgp.Add(func() {
		b = true
		atomic.AddInt32(&sum, 1)
	})
	rgp.Add(func() {
		c = append(c, 1)
		atomic.AddInt32(&sum, 1)
	})
	// run all routines and block until all routine exit
	rgp.Run()
	fmt.Println(b, c, sum)
	// true [1] 2
}
