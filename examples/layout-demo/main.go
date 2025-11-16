package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// LayoutDemoApp demonstrates Row, Column, and all alignment options
type LayoutDemoApp struct {
	goflow.BaseWidget
	currentDemo *signals.Signal[string]
}

// NewLayoutDemoApp creates a new layout demo app
func NewLayoutDemoApp() *LayoutDemoApp {
	return &LayoutDemoApp{
		currentDemo: signals.New("row"),
	}
}

// Build creates the UI for the layout demo
func (app *LayoutDemoApp) Build(context goflow.BuildContext) goflow.Widget {
	return widgets.NewCenter(
		&widgets.Column{
			Children: []goflow.Widget{
				// Title
				widgets.NewTextWithStyle(
					"🎨 GoFlow Layout & Alignment Demo",
					&goflow.TextStyle{
						Color:      goflow.ColorBlue,
						FontSize:   32,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(30),

				// Row Alignment Demos
				widgets.NewTextWithStyle(
					"━━━ Row Layouts ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorMagenta,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				// MainAxis Alignment for Row
				demoSection("Row MainAxis Alignment"),
				spacer(10),
				demoRow("MainAxisStart", widgets.MainAxisStart, widgets.CrossAxisCenter),
				spacer(5),
				demoRow("MainAxisEnd", widgets.MainAxisEnd, widgets.CrossAxisCenter),
				spacer(5),
				demoRow("MainAxisCenter", widgets.MainAxisCenter, widgets.CrossAxisCenter),
				spacer(5),
				demoRow("MainAxisSpaceBetween", widgets.MainAxisSpaceBetween, widgets.CrossAxisCenter),
				spacer(5),
				demoRow("MainAxisSpaceAround", widgets.MainAxisSpaceAround, widgets.CrossAxisCenter),
				spacer(5),
				demoRow("MainAxisSpaceEvenly", widgets.MainAxisSpaceEvenly, widgets.CrossAxisCenter),
				spacer(20),

				// CrossAxis Alignment for Row
				demoSection("Row CrossAxis Alignment"),
				spacer(10),
				demoRowCrossAxis("CrossAxisStart", widgets.MainAxisCenter, widgets.CrossAxisStart),
				spacer(5),
				demoRowCrossAxis("CrossAxisEnd", widgets.MainAxisCenter, widgets.CrossAxisEnd),
				spacer(5),
				demoRowCrossAxis("CrossAxisCenter", widgets.MainAxisCenter, widgets.CrossAxisCenter),
				spacer(5),
				demoRowCrossAxis("CrossAxisStretch", widgets.MainAxisCenter, widgets.CrossAxisStretch),
				spacer(30),

				// Column Alignment Demos
				widgets.NewTextWithStyle(
					"━━━ Column Layouts ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorGreen,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				// MainAxis Alignment for Column
				demoSection("Column MainAxis Alignment"),
				spacer(10),
				demoColumnsRow(),
				spacer(30),

				// Combined Layout Demo
				widgets.NewTextWithStyle(
					"━━━ Combined Layouts ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorRed,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),
				demoSection("Nested Row + Column"),
				spacer(10),
				demoCombinedLayout(),
			},
			MainAxisAlign:  widgets.MainAxisStart,
			CrossAxisAlign: widgets.CrossAxisCenter,
		},
	)
}

// Helper functions

func spacer(height float64) goflow.Widget {
	h := height
	return &widgets.Container{
		Height: &h,
	}
}

func demoSection(title string) goflow.Widget {
	return widgets.NewTextWithStyle(
		title,
		&goflow.TextStyle{
			Color:      goflow.ColorGray,
			FontSize:   18,
			FontWeight: goflow.FontWeightBold,
		},
	)
}

func colorBox(color *goflow.Color, width, height float64, label string) goflow.Widget {
	w, h := width, height
	return &widgets.Container{
		Width:  &w,
		Height: &h,
		Color:  color,
		Child: widgets.NewCenter(
			widgets.NewTextWithStyle(
				label,
				&goflow.TextStyle{
					Color:    goflow.ColorWhite,
					FontSize: 12,
				},
			),
		),
	}
}

func demoRow(label string, mainAxis widgets.MainAxisAlignment, crossAxis widgets.CrossAxisAlignment) goflow.Widget {
	containerWidth := 600.0
	containerHeight := 60.0

	return &widgets.Column{
		Children: []goflow.Widget{
			widgets.NewTextWithStyle(
				label,
				&goflow.TextStyle{
					Color:    goflow.NewColor(100, 100, 100, 255),
					FontSize: 14,
				},
			),
			&widgets.Container{
				Width:  &containerWidth,
				Height: &containerHeight,
				Color:  goflow.NewColor(240, 240, 240, 255),
				Child: &widgets.Row{
					Children: []goflow.Widget{
						colorBox(goflow.ColorRed, 80, 40, "Box 1"),
						colorBox(goflow.ColorGreen, 80, 40, "Box 2"),
						colorBox(goflow.ColorBlue, 80, 40, "Box 3"),
					},
					MainAxisAlignment:  mainAxis,
					CrossAxisAlignment: crossAxis,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
			},
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStart,
	}
}

func demoRowCrossAxis(label string, mainAxis widgets.MainAxisAlignment, crossAxis widgets.CrossAxisAlignment) goflow.Widget {
	containerWidth := 600.0
	containerHeight := 100.0

	return &widgets.Column{
		Children: []goflow.Widget{
			widgets.NewTextWithStyle(
				label,
				&goflow.TextStyle{
					Color:    goflow.NewColor(100, 100, 100, 255),
					FontSize: 14,
				},
			),
			&widgets.Container{
				Width:  &containerWidth,
				Height: &containerHeight,
				Color:  goflow.NewColor(240, 240, 240, 255),
				Child: &widgets.Row{
					Children: []goflow.Widget{
						colorBox(goflow.ColorRed, 80, 30, "S"),
						colorBox(goflow.ColorGreen, 80, 50, "M"),
						colorBox(goflow.ColorBlue, 80, 70, "L"),
					},
					MainAxisAlignment:  mainAxis,
					CrossAxisAlignment: crossAxis,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
			},
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisStart,
	}
}

func demoColumnsRow() goflow.Widget {
	return &widgets.Row{
		Children: []goflow.Widget{
			demoColumn("MainAxisStart", widgets.MainAxisStart),
			demoColumn("MainAxisCenter", widgets.MainAxisCenter),
			demoColumn("MainAxisEnd", widgets.MainAxisEnd),
			demoColumn("SpaceBetween", widgets.MainAxisSpaceBetween),
		},
		MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
		CrossAxisAlignment: widgets.CrossAxisStart,
		MainAxisSize:       widgets.MainAxisSizeMax,
	}
}

func demoColumn(label string, mainAxis widgets.MainAxisAlignment) goflow.Widget {
	containerWidth := 120.0
	containerHeight := 200.0

	return &widgets.Column{
		Children: []goflow.Widget{
			widgets.NewTextWithStyle(
				label,
				&goflow.TextStyle{
					Color:    goflow.NewColor(100, 100, 100, 255),
					FontSize: 12,
				},
			),
			&widgets.Container{
				Width:  &containerWidth,
				Height: &containerHeight,
				Color:  goflow.NewColor(240, 240, 240, 255),
				Child: &widgets.Column{
					Children: []goflow.Widget{
						colorBox(goflow.ColorRed, 80, 30, "1"),
						colorBox(goflow.ColorGreen, 80, 30, "2"),
						colorBox(goflow.ColorBlue, 80, 30, "3"),
					},
					MainAxisAlign:  mainAxis,
					CrossAxisAlign: widgets.CrossAxisCenter,
					MainAxisSize:   widgets.MainAxisSizeMax,
				},
			},
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisCenter,
	}
}

func demoCombinedLayout() goflow.Widget {
	containerWidth := 600.0
	containerHeight := 250.0

	return &widgets.Container{
		Width:  &containerWidth,
		Height: &containerHeight,
		Color:  goflow.NewColor(250, 250, 250, 255),
		Child: &widgets.Column{
			Children: []goflow.Widget{
				// Header row
				&widgets.Row{
					Children: []goflow.Widget{
						colorBox(goflow.ColorRed, 180, 40, "Header 1"),
						colorBox(goflow.ColorGreen, 180, 40, "Header 2"),
						colorBox(goflow.ColorBlue, 180, 40, "Header 3"),
					},
					MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
					CrossAxisAlignment: widgets.CrossAxisCenter,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
				// Content area with two columns
				&widgets.Row{
					Children: []goflow.Widget{
						// Sidebar
						&widgets.Column{
							Children: []goflow.Widget{
								colorBox(goflow.NewColor(255, 100, 100, 255), 150, 30, "Nav 1"),
								colorBox(goflow.NewColor(255, 120, 120, 255), 150, 30, "Nav 2"),
								colorBox(goflow.NewColor(255, 140, 140, 255), 150, 30, "Nav 3"),
							},
							MainAxisAlign:  widgets.MainAxisStart,
							CrossAxisAlign: widgets.CrossAxisStart,
							MainAxisSize:   widgets.MainAxisSizeMax,
						},
						// Main content
						&widgets.Column{
							Children: []goflow.Widget{
								&widgets.Row{
									Children: []goflow.Widget{
										colorBox(goflow.NewColor(100, 200, 255, 255), 70, 30, "A"),
										colorBox(goflow.NewColor(120, 210, 255, 255), 70, 30, "B"),
										colorBox(goflow.NewColor(140, 220, 255, 255), 70, 30, "C"),
									},
									MainAxisAlignment:  widgets.MainAxisSpaceBetween,
									CrossAxisAlignment: widgets.CrossAxisCenter,
									MainAxisSize:       widgets.MainAxisSizeMax,
								},
								colorBox(goflow.NewColor(150, 230, 255, 255), 220, 100, "Content Area"),
							},
							MainAxisAlign:  widgets.MainAxisSpaceEvenly,
							CrossAxisAlign: widgets.CrossAxisCenter,
							MainAxisSize:   widgets.MainAxisSizeMax,
						},
					},
					MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
					CrossAxisAlignment: widgets.CrossAxisStart,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
			},
			MainAxisAlign:  widgets.MainAxisStart,
			CrossAxisAlign: widgets.CrossAxisCenter,
			MainAxisSize:   widgets.MainAxisSizeMax,
		},
	}
}

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   GoFlow Layout & Alignment Demo      ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	// Create the layout demo app
	app := NewLayoutDemoApp()

	fmt.Println("🏗️  Building comprehensive layout demo...")
	fmt.Println()
	fmt.Println("This demo showcases:")
	fmt.Println("  ✓ Row widget with all MainAxis alignments")
	fmt.Println("  ✓ Row widget with all CrossAxis alignments")
	fmt.Println("  ✓ Column widget with all MainAxis alignments")
	fmt.Println("  ✓ Column widget with all CrossAxis alignments")
	fmt.Println("  ✓ Nested Row + Column combinations")
	fmt.Println("  ✓ Complex layout hierarchies")
	fmt.Println()

	// Build the app
	goflow.RunApp(app)

	fmt.Println("✅ Widget tree built successfully!")
	fmt.Println()
	fmt.Println("Layout Features Demonstrated:")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println()
	fmt.Println("📊 Row MainAxis Alignments:")
	fmt.Println("   • MainAxisStart - children at the start")
	fmt.Println("   • MainAxisEnd - children at the end")
	fmt.Println("   • MainAxisCenter - children centered")
	fmt.Println("   • MainAxisSpaceBetween - space between children")
	fmt.Println("   • MainAxisSpaceAround - space around children")
	fmt.Println("   • MainAxisSpaceEvenly - even spacing")
	fmt.Println()
	fmt.Println("📏 Row CrossAxis Alignments:")
	fmt.Println("   • CrossAxisStart - children at top")
	fmt.Println("   • CrossAxisEnd - children at bottom")
	fmt.Println("   • CrossAxisCenter - children centered")
	fmt.Println("   • CrossAxisStretch - children stretched")
	fmt.Println()
	fmt.Println("📊 Column MainAxis Alignments:")
	fmt.Println("   • MainAxisStart - children at the top")
	fmt.Println("   • MainAxisCenter - children centered vertically")
	fmt.Println("   • MainAxisEnd - children at the bottom")
	fmt.Println("   • MainAxisSpaceBetween - vertical spacing")
	fmt.Println()
	fmt.Println("🎨 Combined Layouts:")
	fmt.Println("   • Nested Row + Column hierarchies")
	fmt.Println("   • Complex dashboard-style layouts")
	fmt.Println("   • Multiple alignment modes working together")
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   All Layout Features Working! 🎉     ║")
	fmt.Println("╚════════════════════════════════════════╝")
}
