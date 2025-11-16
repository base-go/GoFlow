package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/signals"
)

type Item struct {
	Name  string
	Price float64
}

func main() {
	fmt.Println("=== Shopping Cart Example ===")

	// Cart items
	items := signals.New([]Item{})

	// Computed total
	total := signals.NewComputed(func() float64 {
		sum := 0.0
		for _, item := range items.Get() {
			sum += item.Price
		}
		return sum
	})

	// Computed item count
	itemCount := signals.NewComputed(func() int {
		return len(items.Get())
	})

	// Display effect
	dispose := signals.NewEffect(func() {
		count := itemCount.Get()
		totalPrice := total.Get()
		fmt.Printf("Cart: %d items, Total: $%.2f\n", count, totalPrice)
	})
	defer dispose()

	// Add items to cart
	fmt.Println("\nAdding items...")
	items.Update(func(cart []Item) []Item {
		return append(cart, Item{"Apple", 1.50})
	})

	items.Update(func(cart []Item) []Item {
		return append(cart, Item{"Banana", 0.75})
	})

	items.Update(func(cart []Item) []Item {
		return append(cart, Item{"Orange", 2.00})
	})

	// Batch update
	fmt.Println("\nBatch adding multiple items...")
	signals.Batch(func() {
		items.Update(func(cart []Item) []Item {
			return append(cart, Item{"Grape", 3.50})
		})
		items.Update(func(cart []Item) []Item {
			return append(cart, Item{"Mango", 4.00})
		})
	})
}
