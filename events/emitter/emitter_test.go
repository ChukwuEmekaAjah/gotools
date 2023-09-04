package emitter

import (
	"fmt"
	"testing"
)

func TestNewEmitter(t *testing.T) {
	AlertEmitter := New()

	if AlertEmitter == nil {
		t.Error("New event emitter should not be nil")
	}
}

func TestOnEmitter(t *testing.T) {
	AlertEmitter := New()

	AlertEmitter.AddListener("shoot", func(eventName string, arguments ...interface{}) {
		fmt.Println("Input values are ", arguments)
	})

	AlertEmitter.On("shoot", func(eventName string, arguments ...interface{}) {
		fmt.Println("They are ", arguments)
	})

	see := func(eventName string, arguments ...interface{}) {
		fmt.Println("Let us see ", arguments)
	}
	fmt.Println("Equality is", &see == &see)
	AlertEmitter.On("shoot", func(eventName string, arguments ...interface{}) {
		fmt.Println("Let us see ", arguments)
	})

	// go AlertEmitter.Listen("shoot")

	AlertEmitter.Emit("shoot", "Chuks", "Ajah")
	AlertEmitter.Emit("shoot", "emeka", "Ukah")

	AlertEmitter.Emit("shoot", "we", "dey", "wait", "you")
}

func TestEventNames(t *testing.T) {
	AlertEmitter := New()

	eventNames := []string{"first", "second", "three"}
	listener := func(eventName string, arguments ...interface{}) {
		fmt.Println("Input values are ", arguments)
	}

	for _, event := range eventNames {
		AlertEmitter.AddListener(event, listener)
	}

	if len(AlertEmitter.EventNames()) != len(eventNames) {
		t.Errorf("Event names should be %d in number", len(eventNames))
	}
}

// func TestRemoveListener(t *testing.T) {
// 	AlertEmitter := New()

// 	eventNames := []string{"first", "second", "three"}
// 	listener := func(eventName string, arguments ...interface{}) {
// 		fmt.Println("Input values are ", arguments)
// 	}

// 	for _, event := range eventNames {
// 		AlertEmitter.AddListener(event, listener)
// 	}

// 	if AlertEmitter.ListenerCount("first") != 1 {
// 		t.Errorf("Expected number of listeners for %s to be %d but got %d", "first", 1, AlertEmitter.ListenerCount("first"))
// 	}

// 	AlertEmitter.RemoveListener("first", listener)

// 	if AlertEmitter.ListenerCount("first") != 0 {
// 		t.Errorf("Expected number of listeners after removal for %s to be %d but got %d", "first", 0, AlertEmitter.ListenerCount("first"))
// 	}
// }
