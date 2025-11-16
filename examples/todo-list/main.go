package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/signals"
)

type Todo struct {
	ID        int
	Text      string
	Completed bool
}

func main() {
	fmt.Println("=== Todo List Example with SignalSlice ===")

	// Reactive todo list
	todos := signals.NewSlice([]Todo{})
	nextID := 1

	// Computed values
	totalCount := todos.Len()
	completedCount := signals.NewComputed(func() int {
		count := 0
		for _, todo := range todos.Get() {
			if todo.Completed {
				count++
			}
		}
		return count
	})

	remainingCount := signals.NewComputed(func() int {
		return totalCount.Get() - completedCount.Get()
	})

	// Effect to display status
	dispose := signals.NewEffect(func() {
		total := totalCount.Get()
		completed := completedCount.Get()
		remaining := remainingCount.Get()
		fmt.Printf("\n📊 Status: %d total, %d completed, %d remaining\n", total, completed, remaining)
	})
	defer dispose()

	// Add todos
	fmt.Println("\n➕ Adding todos...")
	todos.Append(Todo{ID: nextID, Text: "Learn Go", Completed: false})
	nextID++

	todos.Append(Todo{ID: nextID, Text: "Build GoFlow", Completed: false})
	nextID++

	todos.Append(Todo{ID: nextID, Text: "Write examples", Completed: false})
	nextID++

	// Complete a todo
	fmt.Println("\n✅ Completing 'Learn Go'...")
	todos.Update(func(list []Todo) []Todo {
		newList := make([]Todo, len(list))
		copy(newList, list)
		for i, todo := range newList {
			if todo.Text == "Learn Go" {
				newList[i].Completed = true
				break
			}
		}
		return newList
	})

	// Add multiple todos with batch
	fmt.Println("\n➕ Adding multiple todos (batched)...")
	signals.Batch(func() {
		todos.Append(Todo{ID: nextID, Text: "Test signals", Completed: false})
		nextID++
		todos.Append(Todo{ID: nextID, Text: "Optimize performance", Completed: true})
		nextID++
	})

	// Display all todos
	fmt.Println("\n📝 All todos:")
	for _, todo := range todos.Get() {
		status := "⬜"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("  %s %s\n", status, todo.Text)
	}

	// Filter active todos
	activeTodos := todos.Filter(func(todo Todo) bool {
		return !todo.Completed
	})

	fmt.Println("\n🎯 Active todos:")
	for _, todo := range activeTodos.Get() {
		fmt.Printf("  - %s\n", todo.Text)
	}
}
