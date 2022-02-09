package event

import (
	"sync"

	"github.com/shootingfans/redis-dashboard/pkg/routine"
)

// Manager is manage all event, when some event triggered that focus handlers will run
type Manager interface {

	// Trigger will tell all handlers who focus on this event handlers to run
	Trigger(event string, data interface{}) (wait func())

	// FocusOn is register a handler to focus on some event
	FocusOn(event string, handle Handle)
}

// Handle is function called when event triggered
type Handle func(event string, data interface{})

// NewManager return event manager
func NewManager() Manager {
	return &manager{
		handlers: make(map[string][]Handle),
	}
}

type manager struct {
	locker   sync.RWMutex
	handlers map[string][]Handle
}

func (m *manager) Trigger(event string, data interface{}) func() {
	m.locker.RLock()
	handlers, ok := m.handlers[event]
	m.locker.RUnlock()
	if !ok {
		return func() {}
	}
	finish := make(chan struct{})
	go func(e string, d interface{}) {
		defer close(finish)
		grp := routine.NewGroup()
		for _, hdl := range handlers {
			grp.Add(func() {
				hdl(event, data)
			})
		}
		grp.Run()
	}(event, data)
	return func() {
		<-finish
	}
}

func (m *manager) FocusOn(event string, handle Handle) {
	m.locker.Lock()
	m.handlers[event] = append(m.handlers[event], handle)
	m.locker.Unlock()
}
