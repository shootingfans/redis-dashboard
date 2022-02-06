package listener

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestNewManager(t *testing.T) {
	mgr := NewManager()
	var num int32
	for range [10][]int{} {
		mgr.AddLister(func() {
			atomic.AddInt32(&num, 1)
		})
	}
	f := mgr.AddLister(func() {
		atomic.AddInt32(&num, 10)
	})
	mgr.AddLister(func() {
		atomic.AddInt32(&num, 5)
		panic("test")
	})
	go func() {
		f()
	}()
	mgr.Notify()()
	assert.Equal(t, num, int32(25))
}

func ExampleNewManager() {
	mgr := NewManager()
	var a, b, c bool
	// add some listener to manager
	mgr.AddLister(func() {
		a = true
	})
	mgr.AddLister(func() {
		b = true
	})
	m2 := mgr.AddLister(func() {
		c = true
		panic("panic") // when panic the manager will recover safety
	})
	go func() {
		m2() // m2 is function, when call m2 that will block until listener called
	}()
	call := mgr.Notify() // notify all listeners
	call()               // call is a function to block and wait until all listener called
	fmt.Println(a, b, c)
	// true true true
}
