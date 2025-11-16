package main

import (
	"fmt"

	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/signals"
	"github.com/base-go/GoFlow/widgets"
)

// PlaygroundApp is a comprehensive example testing all GoFlow features
type PlaygroundApp struct {
	goflow.BaseWidget
	// Basic signals
	counter     *signals.Signal[int]
	temperature *signals.Signal[float64]
	username    *signals.Signal[string]

	// Computed signals
	counterDoubled  *signals.Computed[int]
	temperatureInF  *signals.Computed[float64]
	greeting        *signals.Computed[string]

	// Collections
	todos       *signals.SignalSlice[string]
	settings    *signals.SignalMap[string, bool]
	todoCount   *signals.Computed[int]
	activeTodos *signals.Computed[[]string]
}

// NewPlaygroundApp creates a new playground app with all features initialized
func NewPlaygroundApp() *PlaygroundApp {
	counter := signals.New(0)
	temperature := signals.New(20.0)
	username := signals.New("Guest")

	todos := signals.NewSlice([]string{
		"Learn GoFlow basics",
		"Build UI components",
		"Test reactive state",
	})

	settings := signals.NewMap(map[string]bool{
		"darkMode":      false,
		"notifications": true,
		"autoSave":      true,
	})

	app := &PlaygroundApp{
		counter:     counter,
		temperature: temperature,
		username:    username,
		todos:       todos,
		settings:    settings,

		// Computed signals demonstrate reactive derivation
		counterDoubled: signals.NewComputed(func() int {
			return counter.Get() * 2
		}),
		temperatureInF: signals.NewComputed(func() float64 {
			return temperature.Get()*9/5 + 32
		}),
		greeting: signals.NewComputed(func() string {
			return fmt.Sprintf("Hello, %s!", username.Get())
		}),
		todoCount: signals.NewComputed(func() int {
			return len(todos.Get())
		}),
		activeTodos: signals.NewComputed(func() []string {
			allTodos := todos.Get()
			// In a real app, we'd filter by completion status
			return allTodos
		}),
	}

	return app
}

// Build creates the UI for the playground
func (app *PlaygroundApp) Build(context goflow.BuildContext) goflow.Widget {
	// Access all signals to create reactive dependencies
	count := app.counter.Get()
	doubled := app.counterDoubled.Get()
	temp := app.temperature.Get()
	tempF := app.temperatureInF.Get()
	greet := app.greeting.Get()
	todoList := app.todos.Get()
	todoTotal := app.todoCount.Get()
	settingsMap := app.settings.Get()

	// Build a comprehensive UI testing all widgets
	return widgets.NewCenter(
		&widgets.Column{
			Children: []goflow.Widget{
				// Header section
				widgets.NewTextWithStyle(
					"🎮 GoFlow Playground",
					&goflow.TextStyle{
						Color:      goflow.ColorBlue,
						FontSize:   32,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				// User greeting section
				containerWithStyle(
					widgets.NewTextWithStyle(
						greet,
						&goflow.TextStyle{
							Color:    goflow.ColorGreen,
							FontSize: 20,
						},
					),
					goflow.NewColor(240, 255, 240, 255), // Light green background
					16,
				),
				spacer(20),

				// Counter section
				widgets.NewTextWithStyle(
					"📊 Counter Demo",
					&goflow.TextStyle{
						Color:      goflow.ColorMagenta,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(10),
				widgets.NewText(fmt.Sprintf("Counter: %d", count)),
				widgets.NewText(fmt.Sprintf("Doubled: %d (computed)", doubled)),
				spacer(20),

				// Temperature section
				widgets.NewTextWithStyle(
					"🌡️ Temperature Converter",
					&goflow.TextStyle{
						Color:      goflow.NewColor(255, 140, 0, 255), // Orange color
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(10),
				widgets.NewText(fmt.Sprintf("Temperature: %.1f°C", temp)),
				widgets.NewText(fmt.Sprintf("In Fahrenheit: %.1f°F (computed)", tempF)),
				spacer(20),

				// Todo list section
				widgets.NewTextWithStyle(
					fmt.Sprintf("✅ Todo List (%d items)", todoTotal),
					&goflow.TextStyle{
						Color:      goflow.ColorRed,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(10),
				buildTodoList(todoList),
				spacer(20),

				// Settings section
				widgets.NewTextWithStyle(
					"⚙️ Settings",
					&goflow.TextStyle{
						Color:      goflow.ColorGray,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(10),
				buildSettings(settingsMap),
				spacer(20),

				// Layout demo section
				widgets.NewTextWithStyle(
					"🎨 Layout Demo",
					&goflow.TextStyle{
						Color:      goflow.ColorBlue,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(10),
				buildLayoutDemo(),
			},
			MainAxisAlign:  widgets.MainAxisCenter,
			CrossAxisAlign: widgets.CrossAxisCenter,
		},
	)
}

// Helper functions for building UI components

func spacer(height float64) goflow.Widget {
	h := height
	return &widgets.Container{
		Height: &h,
	}
}

func containerWithStyle(child goflow.Widget, color *goflow.Color, padding float64) goflow.Widget {
	return &widgets.Container{
		Child:   child,
		Color:   color,
		Padding: goflow.NewEdgeInsetsAll(padding),
	}
}

func buildTodoList(todos []string) goflow.Widget {
	children := make([]goflow.Widget, 0, len(todos))
	for i, todo := range todos {
		todoText := fmt.Sprintf("%d. %s", i+1, todo)
		children = append(children, widgets.NewText(todoText))
	}
	return &widgets.Column{
		Children:       children,
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStart,
	}
}

func buildSettings(settings map[string]bool) goflow.Widget {
	children := make([]goflow.Widget, 0)
	for key, value := range settings {
		status := "OFF"
		if value {
			status = "ON"
		}
		settingText := fmt.Sprintf("%s: %s", key, status)
		children = append(children, widgets.NewText(settingText))
	}
	return &widgets.Column{
		Children:       children,
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStart,
	}
}

func buildLayoutDemo() goflow.Widget {
	return &widgets.Column{
		Children: []goflow.Widget{
			// Container with padding
			containerWithStyle(
				widgets.NewText("Container with padding"),
				goflow.NewColor(255, 240, 240, 255), // Light red
				12,
			),
			spacer(10),
			// Container with color
			&widgets.Container{
				Width:  float64Ptr(200),
				Height: float64Ptr(40),
				Color:  goflow.ColorBlue,
				Child: widgets.NewCenter(
					widgets.NewTextWithStyle(
						"Sized container",
						&goflow.TextStyle{
							Color: goflow.ColorWhite,
						},
					),
				),
			},
			spacer(10),
			// Container with margin and padding
			&widgets.Container{
				Padding: goflow.NewEdgeInsetsSymmetric(20, 10),
				Color:   goflow.NewColor(240, 240, 255, 255), // Light blue
				Child:   widgets.NewText("Padding: H=20, V=10"),
			},
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStart,
	}
}

func float64Ptr(f float64) *float64 {
	return &f
}

// Methods to interact with the app

func (app *PlaygroundApp) IncrementCounter() {
	app.counter.Update(func(v int) int { return v + 1 })
}

func (app *PlaygroundApp) DecrementCounter() {
	app.counter.Update(func(v int) int { return v - 1 })
}

func (app *PlaygroundApp) SetTemperature(temp float64) {
	app.temperature.Set(temp)
}

func (app *PlaygroundApp) SetUsername(name string) {
	app.username.Set(name)
}

func (app *PlaygroundApp) AddTodo(todo string) {
	app.todos.Append(todo)
}

func (app *PlaygroundApp) RemoveTodo(index int) {
	app.todos.RemoveAt(index)
}

func (app *PlaygroundApp) ToggleSetting(key string) {
	app.settings.Update(func(m map[string]bool) map[string]bool {
		newMap := make(map[string]bool)
		for k, v := range m {
			newMap[k] = v
		}
		newMap[key] = !newMap[key]
		return newMap
	})
}

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   GoFlow Framework Playground Test    ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	// Create the playground app
	app := NewPlaygroundApp()

	// Set up reactive effects to track changes
	fmt.Println("📡 Setting up reactive effects...")

	counterDispose := signals.NewEffect(func() {
		count := app.counter.Get()
		fmt.Printf("  → Counter changed: %d\n", count)
	})
	defer counterDispose()

	tempDispose := signals.NewEffect(func() {
		temp := app.temperature.Get()
		fmt.Printf("  → Temperature changed: %.1f°C\n", temp)
	})
	defer tempDispose()

	todoDispose := signals.NewEffect(func() {
		count := app.todoCount.Get()
		fmt.Printf("  → Todo count changed: %d items\n", count)
	})
	defer todoDispose()

	fmt.Println()

	// Build and run the app
	fmt.Println("🏗️  Building widget tree...")
	goflow.RunApp(app)
	fmt.Println("✅ Widget tree built successfully!")
	fmt.Println()

	// Test interactions
	fmt.Println("🧪 Testing Framework Features")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println()

	// Test 1: Basic signal updates
	fmt.Println("Test 1: Basic Signal Updates")
	fmt.Println("Incrementing counter 3 times...")
	app.IncrementCounter()
	app.IncrementCounter()
	app.IncrementCounter()
	fmt.Printf("Counter value: %d\n", app.counter.Get())
	fmt.Printf("Doubled value (computed): %d\n", app.counterDoubled.Get())
	fmt.Println()

	// Test 2: Computed signals
	fmt.Println("Test 2: Computed Signals")
	fmt.Println("Setting temperature to 25°C...")
	app.SetTemperature(25)
	fmt.Printf("Temperature: %.1f°C → %.1f°F (computed)\n",
		app.temperature.Get(), app.temperatureInF.Get())
	fmt.Println()

	// Test 3: String signals
	fmt.Println("Test 3: String Signals & Computed")
	fmt.Println("Setting username to 'Alice'...")
	app.SetUsername("Alice")
	fmt.Printf("Greeting (computed): %s\n", app.greeting.Get())
	fmt.Println()

	// Test 4: Signal collections - SliceSignal
	fmt.Println("Test 4: Signal Collections - Slice")
	fmt.Println("Adding new todo...")
	app.AddTodo("Explore advanced features")
	fmt.Printf("Todo count: %d\n", app.todoCount.Get())
	fmt.Println("All todos:")
	for i, todo := range app.todos.Get() {
		fmt.Printf("  %d. %s\n", i+1, todo)
	}
	fmt.Println()

	fmt.Println("Removing first todo...")
	app.RemoveTodo(0)
	fmt.Printf("Todo count: %d\n", app.todoCount.Get())
	fmt.Println()

	// Test 5: Signal collections - MapSignal
	fmt.Println("Test 5: Signal Collections - Map")
	fmt.Println("Toggling 'darkMode' setting...")
	app.ToggleSetting("darkMode")
	fmt.Println("Current settings:")
	for key, value := range app.settings.Get() {
		fmt.Printf("  %s: %v\n", key, value)
	}
	fmt.Println()

	// Test 6: Batch updates
	fmt.Println("Test 6: Batch Updates")
	fmt.Println("Performing multiple updates in batch...")
	signals.Batch(func() {
		app.IncrementCounter()
		app.IncrementCounter()
		app.SetTemperature(30)
		app.SetUsername("Bob")
	})
	fmt.Printf("Counter: %d, Temperature: %.1f°C, User: %s\n",
		app.counter.Get(), app.temperature.Get(), app.username.Get())
	fmt.Println()

	// Test 7: Effect cleanup and final state
	fmt.Println("Test 7: Final State Check")
	fmt.Println("────────────────────────────────")
	fmt.Printf("✓ Counter: %d (doubled: %d)\n", app.counter.Get(), app.counterDoubled.Get())
	fmt.Printf("✓ Temperature: %.1f°C (%.1f°F)\n", app.temperature.Get(), app.temperatureInF.Get())
	fmt.Printf("✓ User: %s\n", app.username.Get())
	fmt.Printf("✓ Todos: %d items\n", app.todoCount.Get())
	fmt.Printf("✓ Settings: %d configured\n", len(app.settings.Get()))
	fmt.Println()

	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   All Framework Features Tested! 🎉   ║")
	fmt.Println("╚════════════════════════════════════════╝")
}
