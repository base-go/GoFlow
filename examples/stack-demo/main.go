package main

import (
	"fmt"

	"github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// StackDemoApp demonstrates Stack and Positioned widgets
type StackDemoApp struct {
	goflow.BaseWidget
}

// NewStackDemoApp creates a new stack demo app
func NewStackDemoApp() *StackDemoApp {
	return &StackDemoApp{}
}

// Build creates the UI for the stack demo
func (app *StackDemoApp) Build(context goflow.BuildContext) goflow.Widget {
	return widgets.NewCenter(
		&widgets.Column{
			Children: []goflow.Widget{
				// Title
				widgets.NewTextWithStyle(
					"🎨 GoFlow Stack & Positioned Demo",
					&goflow.TextStyle{
						Color:      goflow.ColorBlue,
						FontSize:   32,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(30),

				// Stack Alignment Demos
				widgets.NewTextWithStyle(
					"━━━ Stack Alignment Modes ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorMagenta,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				// Row of alignment demos
				&widgets.Row{
					Children: []goflow.Widget{
						stackAlignmentDemo("TopLeft", widgets.StackAlignmentTopLeft),
						stackAlignmentDemo("TopCenter", widgets.StackAlignmentTopCenter),
						stackAlignmentDemo("TopRight", widgets.StackAlignmentTopRight),
					},
					MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
					CrossAxisAlignment: widgets.CrossAxisStart,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
				spacer(10),
				&widgets.Row{
					Children: []goflow.Widget{
						stackAlignmentDemo("CenterLeft", widgets.StackAlignmentCenterLeft),
						stackAlignmentDemo("Center", widgets.StackAlignmentCenter),
						stackAlignmentDemo("CenterRight", widgets.StackAlignmentCenterRight),
					},
					MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
					CrossAxisAlignment: widgets.CrossAxisStart,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
				spacer(10),
				&widgets.Row{
					Children: []goflow.Widget{
						stackAlignmentDemo("BottomLeft", widgets.StackAlignmentBottomLeft),
						stackAlignmentDemo("BottomCenter", widgets.StackAlignmentBottomCenter),
						stackAlignmentDemo("BottomRight", widgets.StackAlignmentBottomRight),
					},
					MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
					CrossAxisAlignment: widgets.CrossAxisStart,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
				spacer(30),

				// Stack Fit Modes
				widgets.NewTextWithStyle(
					"━━━ Stack Fit Modes ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorGreen,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				&widgets.Row{
					Children: []goflow.Widget{
						stackFitDemo("Loose", widgets.StackFitLoose),
						stackFitDemo("Expand", widgets.StackFitExpand),
						stackFitDemo("Passthrough", widgets.StackFitPassthrough),
					},
					MainAxisAlignment:  widgets.MainAxisSpaceEvenly,
					CrossAxisAlignment: widgets.CrossAxisStart,
					MainAxisSize:       widgets.MainAxisSizeMax,
				},
				spacer(30),

				// Positioned Examples
				widgets.NewTextWithStyle(
					"━━━ Positioned Widget ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorRed,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				positionedDemo(),
				spacer(30),

				// Complex Example
				widgets.NewTextWithStyle(
					"━━━ Complex Stack Example ━━━",
					&goflow.TextStyle{
						Color:      goflow.ColorCyan,
						FontSize:   24,
						FontWeight: goflow.FontWeightBold,
					},
				),
				spacer(20),

				complexStackExample(),
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

func stackAlignmentDemo(label string, alignment widgets.StackAlignment) goflow.Widget {
	containerWidth := 150.0
	containerHeight := 150.0

	return &widgets.Column{
		Children: []goflow.Widget{
			widgets.NewTextWithStyle(
				label,
				&goflow.TextStyle{
					Color:    goflow.NewColor(100, 100, 100, 255),
					FontSize: 12,
				},
			),
			spacer(5),
			&widgets.Container{
				Width:  &containerWidth,
				Height: &containerHeight,
				Color:  goflow.NewColor(240, 240, 240, 255),
				Child: &widgets.Stack{
					Children: []goflow.Widget{
						// Large background box
						colorBox(
							goflow.NewColor(200, 200, 255, 255),
							120,
							120,
							"BG",
						),
						// Small foreground box
						colorBox(
							goflow.NewColor(255, 100, 100, 255),
							50,
							50,
							"FG",
						),
					},
					Alignment: alignment,
					Fit:       widgets.StackFitLoose,
				},
			},
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisCenter,
	}
}

func stackFitDemo(label string, fit widgets.StackFit) goflow.Widget {
	containerWidth := 180.0
	containerHeight := 180.0

	return &widgets.Column{
		Children: []goflow.Widget{
			widgets.NewTextWithStyle(
				label,
				&goflow.TextStyle{
					Color:    goflow.NewColor(100, 100, 100, 255),
					FontSize: 12,
				},
			),
			spacer(5),
			&widgets.Container{
				Width:  &containerWidth,
				Height: &containerHeight,
				Color:  goflow.NewColor(240, 240, 240, 255),
				Child: &widgets.Stack{
					Children: []goflow.Widget{
						// Small child - behavior depends on fit mode
						colorBox(
							goflow.NewColor(100, 200, 100, 255),
							80,
							80,
							"Child",
						),
						// Smaller overlay
						colorBox(
							goflow.NewColor(255, 150, 50, 200),
							40,
							40,
							"Top",
						),
					},
					Alignment: widgets.StackAlignmentCenter,
					Fit:       fit,
				},
			},
		},
		MainAxisAlign:  widgets.MainAxisStart,
		CrossAxisAlign: widgets.CrossAxisCenter,
	}
}

func positionedDemo() goflow.Widget {
	containerWidth := 600.0
	containerHeight := 300.0

	left10 := 10.0
	top10 := 10.0
	right10 := 10.0
	bottom10 := 10.0
	left200 := 200.0
	top100 := 100.0

	return &widgets.Container{
		Width:  &containerWidth,
		Height: &containerHeight,
		Color:  goflow.NewColor(250, 250, 250, 255),
		Child: &widgets.Stack{
			Children: []goflow.Widget{
				// Background
				&widgets.ColoredBox{
					Color: goflow.NewColor(230, 230, 230, 255),
				},
				// Top-left positioned
				&widgets.Positioned{
					Left: &left10,
					Top:  &top10,
					Child: colorBox(
						goflow.ColorRed,
						80,
						60,
						"Top-Left",
					),
				},
				// Top-right positioned
				&widgets.Positioned{
					Right: &right10,
					Top:   &top10,
					Child: colorBox(
						goflow.ColorGreen,
						80,
						60,
						"Top-Right",
					),
				},
				// Bottom-left positioned
				&widgets.Positioned{
					Left:   &left10,
					Bottom: &bottom10,
					Child: colorBox(
						goflow.ColorBlue,
						80,
						60,
						"Bottom-Left",
					),
				},
				// Bottom-right positioned
				&widgets.Positioned{
					Right:  &right10,
					Bottom: &bottom10,
					Child: colorBox(
						goflow.ColorYellow,
						80,
						60,
						"Bottom-Right",
					),
				},
				// Center positioned
				&widgets.Positioned{
					Left: &left200,
					Top:  &top100,
					Child: colorBox(
						goflow.ColorMagenta,
						100,
						80,
						"Absolute\nPosition",
					),
				},
			},
			Fit: widgets.StackFitExpand,
		},
	}
}

func complexStackExample() goflow.Widget {
	containerWidth := 600.0
	containerHeight := 400.0

	left0 := 0.0
	right0 := 0.0
	bottom0 := 0.0
	left20 := 20.0
	top20 := 20.0

	return &widgets.Container{
		Width:  &containerWidth,
		Height: &containerHeight,
		Color:  goflow.NewColor(255, 255, 255, 255),
		Child: &widgets.Stack{
			Children: []goflow.Widget{
				// Background image placeholder
				&widgets.ColoredBox{
					Color: goflow.NewColor(100, 150, 200, 255),
				},
				// Semi-transparent overlay
				&widgets.Positioned{
					Left:   &left0,
					Right:  &right0,
					Bottom: &bottom0,
					Child: &widgets.Container{
						Height: float64Ptr(150),
						Color:  goflow.NewColor(0, 0, 0, 180),
						Child: &widgets.Padding{
							Padding: goflow.NewEdgeInsetsAll(20),
							Child: &widgets.Column{
								Children: []goflow.Widget{
									widgets.NewTextWithStyle(
										"Stack Demo Card",
										&goflow.TextStyle{
											Color:      goflow.ColorWhite,
											FontSize:   24,
											FontWeight: goflow.FontWeightBold,
										},
									),
									spacer(10),
									widgets.NewTextWithStyle(
										"This demonstrates a complex stack with positioned overlay",
										&goflow.TextStyle{
											Color:    goflow.NewColor(220, 220, 220, 255),
											FontSize: 14,
										},
									),
								},
								MainAxisAlign:  widgets.MainAxisStart,
								CrossAxisAlign: widgets.CrossAxisStart,
							},
						},
					},
				},
				// Top-left badge
				&widgets.Positioned{
					Left: &left20,
					Top:  &top20,
					Child: &widgets.Container{
						Padding: goflow.NewEdgeInsetsSymmetric(16, 8),
						Color:   goflow.NewColor(255, 100, 100, 255),
						Child: widgets.NewTextWithStyle(
							"NEW",
							&goflow.TextStyle{
								Color:      goflow.ColorWhite,
								FontSize:   12,
								FontWeight: goflow.FontWeightBold,
							},
						),
					},
				},
				// Top-right icon
				&widgets.Positioned{
					Right: &left20,
					Top:   &top20,
					Child: colorBox(
						goflow.NewColor(255, 255, 255, 200),
						40,
						40,
						"❤️",
					),
				},
			},
			Fit: widgets.StackFitExpand,
		},
	}
}

func float64Ptr(f float64) *float64 {
	return &f
}

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   GoFlow Stack & Positioned Demo      ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	// Create the stack demo app
	app := NewStackDemoApp()

	fmt.Println("🏗️  Building comprehensive stack demo...")
	fmt.Println()
	fmt.Println("This demo showcases:")
	fmt.Println("  ✓ Stack widget with all alignment modes")
	fmt.Println("  ✓ Stack fit modes (Loose, Expand, Passthrough)")
	fmt.Println("  ✓ Positioned widget for absolute positioning")
	fmt.Println("  ✓ Complex layered layouts")
	fmt.Println("  ✓ Overlay patterns (badges, dialogs, etc.)")
	fmt.Println()

	// Build the app
	goflow.RunApp(app)

	fmt.Println("✅ Widget tree built successfully!")
	fmt.Println()
	fmt.Println("Stack Features Demonstrated:")
	fmt.Println("─────────────────────────────────────────")
	fmt.Println()
	fmt.Println("📊 Stack Alignment Modes:")
	fmt.Println("   • TopLeft, TopCenter, TopRight")
	fmt.Println("   • CenterLeft, Center, CenterRight")
	fmt.Println("   • BottomLeft, BottomCenter, BottomRight")
	fmt.Println()
	fmt.Println("📏 Stack Fit Modes:")
	fmt.Println("   • Loose - children sized to natural size")
	fmt.Println("   • Expand - children expanded to fill stack")
	fmt.Println("   • Passthrough - pass constraints through")
	fmt.Println()
	fmt.Println("🎯 Positioned Widget:")
	fmt.Println("   • Absolute positioning with Left/Top/Right/Bottom")
	fmt.Println("   • Flexible combinations of positioning")
	fmt.Println("   • Perfect for overlays and badges")
	fmt.Println()
	fmt.Println("🎨 Complex Patterns:")
	fmt.Println("   • Image with text overlay")
	fmt.Println("   • Corner badges and icons")
	fmt.Println("   • Semi-transparent layers")
	fmt.Println("   • Card-style layouts")
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   All Stack Features Working! 🎉      ║")
	fmt.Println("╚════════════════════════════════════════╝")
}
