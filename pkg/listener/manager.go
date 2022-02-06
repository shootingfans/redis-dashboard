// Package listener is defined a listener manager to listen some event and notify all listeners
package listener

import (
	"github.com/shootingfans/redis-dashboard/pkg/routine"
	"sync"
)

// Manager is manage all listener and notify them all
type Manager interface {
	// AddLister add new listener to manager
	// The result function will wait and block until function called
	AddLister(fn func()) func()

	// Notify will notify all listener
	// The result function will wait and block until all listeners called
	Notify() func()
}

type manager struct {
	locker    sync.Mutex
	listeners []func()
	waitGroup sync.WaitGroup
}

func (m *manager) AddLister(fn func()) func() {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.waitGroup.Add(1)
	m.listeners = append(m.listeners, func() {
		defer m.waitGroup.Done()
		fn()
	})
	return func() {
		m.waitGroup.Wait()
	}
}

func (m *manager) Notify() func() {
	m.locker.Lock()
	defer m.locker.Unlock()
	mgr := routine.NewGroup()
	for _, f := range m.listeners {
		mgr.Add(f)
	}
	go mgr.Run()
	return func() {
		m.waitGroup.Wait()
	}
}

// NewManager create a listener manager
func NewManager() Manager {
	return new(manager)
}
