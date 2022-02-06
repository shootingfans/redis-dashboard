// Package routine is defined a routine group to manage many routines
package routine

import "sync"

// Group is manage some routine in a group
type Group interface {
	// Add is append routine func to  group
	Add(fn func())
	// Run is run group and wait all finish
	Run()
}

type group struct {
	waitGroup sync.WaitGroup
	locker    sync.Mutex
	routines  []func()
}

func (g *group) Add(fn func()) {
	g.locker.Lock()
	defer g.locker.Unlock()
	g.waitGroup.Add(1)
	g.routines = append(g.routines, fn)
}

func (g *group) Run() {
	g.locker.Lock()
	defer g.locker.Unlock()
	for _, f := range g.routines {
		go func(fn func()) {
			defer g.waitGroup.Done()
			saveRun(fn)
		}(f)
	}
	g.waitGroup.Wait()
}

func saveRun(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	fn()
}

// NewGroup create routine group
func NewGroup() Group {
	return new(group)
}
