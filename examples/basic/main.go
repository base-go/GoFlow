package main

import (
	"fmt"

	"github.com/base-go/GoFlow/signals"
)

func main() {
	fmt.Println("=== Basic Signal Usage ===")

	// Create a signal
	counter := signals.New(0)

	// Create a computed value
	doubled := signals.NewComputed(func() int {
		return counter.Get() * 2
	})

	// Create an effect that prints when counter changes
	dispose := signals.NewEffect(func() {
		fmt.Printf("Counter: %d, Doubled: %d\n", counter.Get(), doubled.Get())
	})
	defer dispose()

	// Update the counter
	counter.Set(5)
	counter.Set(10)
	counter.Update(func(v int) int {
		return v + 5
	})
}
