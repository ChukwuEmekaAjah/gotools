package main

import (
	"fmt"

	"github.com/ChukwuemekaAjah/gotools/events/emitter"
)

func main() {
	AlertEmitter := emitter.New()

	AlertEmitter.AddListener("shoot", func(eventName string, arguments ...interface{}) {
		fmt.Println("Input values are ", arguments)
	})

	AlertEmitter.On("shoot", func(eventName string, arguments ...interface{}) {
		fmt.Println("They are ", arguments)
	})

	AlertEmitter.On("shoot", func(eventName string, arguments ...interface{}) {
		fmt.Println("Let us see ", arguments)
	})

	// go AlertEmitter.Listen("shoot")

	AlertEmitter.Emit("shoot", "Chuks", "Ajah")
	AlertEmitter.Emit("shoot", "emeka", "Ukah")

	AlertEmitter.Emit("shoot", "we", "dey", "wait", "you")
	fmt.Println("Hello world")
}
