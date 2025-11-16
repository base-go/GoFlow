package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/signals"
)

func main() {
	fmt.Println("=== Counter Application ===")

	// State
	count := signals.New(0)
	multiplier := signals.New(2)

	// Derived state
	doubled := signals.NewComputed(func() int {
		return count.Get() * 2
	})

	tripled := signals.NewComputed(func() int {
		return count.Get() * 3
	})

	customMultiplied := signals.NewComputed(func() int {
		return count.Get() * multiplier.Get()
	})

	// Effects
	dispose1 := signals.NewEffect(func() {
		fmt.Printf("Count changed: %d\n", count.Get())
	})
	defer dispose1()

	dispose2 := signals.NewEffect(func() {
		// Use Peek to avoid reacting to every count change
		currentCount := signals.Untracked(func() int {
			return count.Peek()
		})
		mult := multiplier.Get()
		fmt.Printf("Multiplier changed to %d (count is %d)\n", mult, currentCount)
	})
	defer dispose2()

	// Simulate app interactions
	fmt.Println("\n--- Incrementing counter ---")
	count.Update(func(v int) int { return v + 1 })
	count.Update(func(v int) int { return v + 1 })

	fmt.Println("\n--- Checking computed values ---")
	fmt.Printf("Doubled: %d\n", doubled.Get())
	fmt.Printf("Tripled: %d\n", tripled.Get())
	fmt.Printf("Custom (×%d): %d\n", multiplier.Peek(), customMultiplied.Get())

	fmt.Println("\n--- Changing multiplier ---")
	multiplier.Set(5)

	fmt.Println("\n--- Batch increment and multiply ---")
	signals.Batch(func() {
		count.Update(func(v int) int { return v + 3 })
		multiplier.Set(10)
	})

	fmt.Printf("Final custom (×%d): %d\n", multiplier.Peek(), customMultiplied.Get())
}
