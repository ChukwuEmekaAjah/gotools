package emitter

import (
	"errors"
)

type Emitter struct {
	listeners    map[string][]func(eventName string, arguments ...interface{})
	maxListeners uint
}

func New() *Emitter {
	newEmitter := Emitter{maxListeners: 10, listeners: make(map[string][]func(eventName string, arguments ...interface{}))}
	return &newEmitter
}

func (e *Emitter) AddListener(eventName string, listener func(eventName string, arguments ...interface{})) (*Emitter, error) {
	if e.maxListeners > 0 && uint(len(e.listeners)) >= e.maxListeners {
		return nil, errors.New("Maximum number of event emitter listeners reached")
	}

	_, exists := e.listeners[eventName]
	if !exists {
		e.listeners[eventName] = []func(eventName string, arguments ...interface{}){listener}
	} else {
		e.listeners[eventName] = append(e.listeners[eventName], listener)
	}

	return e, nil
}

func (e *Emitter) RemoveAllListeners(eventName string) (*Emitter, error) {
	_, exists := e.listeners[eventName]

	if !exists {
		return e, nil
	}
	delete(e.listeners, eventName)
	return e, nil
}

func (e *Emitter) Emit(eventName string, arguments ...interface{}) {

	listeners := e.listeners[eventName]
	for _, eventLister := range listeners {
		eventLister(eventName, arguments...)
	}

}

func (e *Emitter) On(eventName string, listener func(eventName string, arguments ...interface{})) {
	e.AddListener(eventName, listener)
}

func (e *Emitter) RemoveListener(eventName string, listener func(eventName string, arguments ...interface{})) *Emitter {
	listeners, exists := e.listeners[eventName]

	if !exists {
		return e
	}

	for i := 0; i < len(listeners); i++ {
		if &listeners[i] == &listener {
			listeners = append(listeners[0:i], listeners[i+1:]...)
			break
		}
	}

	e.listeners[eventName] = listeners
	return e
}

func (e Emitter) EventNames() []string {
	names := []string{}
	for event := range e.listeners {
		names = append(names, event)
	}
	return names
}

func (e *Emitter) ListenerCount(eventName string) int {
	_, exists := e.listeners[eventName]

	if !exists {
		return 0
	}

	return len(e.listeners[eventName])
}

func (e Emitter) GetMaxListeners() uint {
	return e.maxListeners
}

func (e *Emitter) SetMaxListeners(count uint) *Emitter {
	e.maxListeners = count
	return e
}

func (e Emitter) Listeners(eventName string) []func(eventName string, arguments ...interface{}) {
	_, exists := e.listeners[eventName]

	if !exists {
		return []func(eventName string, arguments ...interface{}){}
	}

	return e.listeners[eventName]
}
