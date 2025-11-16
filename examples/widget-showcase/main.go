package main

import (
	"fmt"
	"time"

	"github.com/base-go/GoFlow/pkg/core/animation"
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

func main() {
	app := &WidgetShowcaseApp{}
	goflow.RunApp(app)
}

// WidgetShowcaseApp demonstrates all new widgets
type WidgetShowcaseApp struct {
	goflow.BaseWidget
}

func (app *WidgetShowcaseApp) Build(context goflow.BuildContext) goflow.Widget {
	return &MaterialApp{
		Title: "GoFlow Widget Showcase",
		Home:  NewShowcaseHome(),
	}
}

// MaterialApp is a simple Material Design app wrapper
type MaterialApp struct {
	goflow.BaseWidget
	Title string
	Home  goflow.Widget
}

func (ma *MaterialApp) Build(context goflow.BuildContext) goflow.Widget {
	return ma.Home
}

// ShowcaseHome is the home page with tabs for different widget categories
type ShowcaseHome struct {
	goflow.BaseWidget
	currentTab int
}

func NewShowcaseHome() *ShowcaseHome {
	return &ShowcaseHome{
		currentTab: 0,
	}
}

func (sh *ShowcaseHome) Build(context goflow.BuildContext) goflow.Widget {
	tabs := []string{
		"Forms",
		"Inputs",
		"Animations",
		"Data Display",
		"Pickers",
	}

	return &widgets.Column{
		Children: []goflow.Widget{
			// Header
			&widgets.Container{
				Color: goflow.NewColor(33, 150, 243, 255),
				Child: &widgets.Text{
					Data: "GoFlow Widget Showcase",
					Style: goflow.NewTextStyle().
						WithFontSize(24).
						WithColor(goflow.NewColor(255, 255, 255, 255)),
				},
				Padding: goflow.NewEdgeInsetsAll(16),
			},

			// Tab bar
			sh.buildTabBar(tabs),

			// Content
			sh.buildTabContent(),
		},
	}
}

func (sh *ShowcaseHome) buildTabBar(tabs []string) goflow.Widget {
	tabWidgets := make([]goflow.Widget, len(tabs))

	for i, tab := range tabs {
		tabIndex := i
		isSelected := sh.currentTab == i

		bgColor := goflow.NewColor(200, 200, 200, 255)
		textColor := goflow.NewColor(0, 0, 0, 255)

		if isSelected {
			bgColor = goflow.NewColor(33, 150, 243, 255)
			textColor = goflow.NewColor(255, 255, 255, 255)
		}

		tabWidgets[i] = &widgets.GestureDetector{
			OnTap: func() {
				sh.currentTab = tabIndex
			},
			Child: &widgets.Container{
				Color: bgColor,
				Child: &widgets.Text{
					Data:  tab,
					Style: goflow.NewTextStyle().WithColor(textColor),
				},
				Padding: goflow.NewEdgeInsets(12, 8, 12, 8),
			},
		}
	}

	return &widgets.Row{
		MainAxisAlignment: widgets.MainAxisAlignmentSpaceEvenly,
		Children:          tabWidgets,
	}
}

func (sh *ShowcaseHome) buildTabContent() goflow.Widget {
	switch sh.currentTab {
	case 0:
		return NewFormShowcase()
	case 1:
		return NewInputShowcase()
	case 2:
		return NewAnimationShowcase()
	case 3:
		return NewDataDisplayShowcase()
	case 4:
		return NewPickerShowcase()
	default:
		return &widgets.Text{
			Data:  "Unknown tab",
			Style: goflow.NewTextStyle(),
		}
	}
}

// FormShowcase demonstrates form widgets
type FormShowcase struct {
	goflow.BaseWidget
	formKey *widgets.FormKey
}

func NewFormShowcase() *FormShowcase {
	return &FormShowcase{
		formKey: widgets.NewFormKey(),
	}
}

func (fs *FormShowcase) Build(context goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Child: widgets.NewForm(
			fs.formKey,
			&widgets.Column{
				CrossAxisAlignment: widgets.CrossAxisAlignmentStretch,
				Children: []goflow.Widget{
					&widgets.Text{
						Data:  "Form Example",
						Style: goflow.NewTextStyle().WithFontSize(20),
					},

					&widgets.SizedBox{Height: 16},

					// Email field
					widgets.NewTextFormField().
						WithDecoration(&widgets.InputDecoration{
							HintText: "Enter your email",
						}).
						WithValidator(widgets.ComposeValidators(
							widgets.RequiredValidator("Email is required"),
							widgets.EmailValidator("Please enter a valid email"),
						)),

					&widgets.SizedBox{Height: 16},

					// Password field
					widgets.NewTextFormField().
						WithObscureText(true).
						WithDecoration(&widgets.InputDecoration{
							HintText: "Enter your password",
						}).
						WithValidator(widgets.MinLengthValidator(6, "Password must be at least 6 characters")),

					&widgets.SizedBox{Height: 24},

					// Submit button
					&widgets.GestureDetector{
						OnTap: func() {
							if fs.formKey.CurrentState().Validate() {
								fmt.Println("Form is valid!")
								fs.formKey.CurrentState().Save()
							} else {
								fmt.Println("Form has errors")
							}
						},
						Child: &widgets.Container{
							Color: goflow.NewColor(33, 150, 243, 255),
							Child: &widgets.Center{
								Child: &widgets.Text{
									Data: "Submit",
									Style: goflow.NewTextStyle().
										WithColor(goflow.NewColor(255, 255, 255, 255)),
								},
							},
							Padding: goflow.NewEdgeInsets(16, 12, 16, 12),
						},
					},
				},
			},
		),
		Padding: goflow.NewEdgeInsetsAll(16),
	}
}

// InputShowcase demonstrates input widgets
type InputShowcase struct {
	goflow.BaseWidget
	checkboxValue bool
	switchValue   bool
	sliderValue   float64
}

func NewInputShowcase() *InputShowcase {
	return &InputShowcase{
		checkboxValue: false,
		switchValue:   false,
		sliderValue:   50.0,
	}
}

func (is *InputShowcase) Build(context goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Child: &widgets.Column{
			CrossAxisAlignment: widgets.CrossAxisAlignmentStart,
			Children: []goflow.Widget{
				&widgets.Text{
					Data:  "Input Widgets",
					Style: goflow.NewTextStyle().WithFontSize(20),
				},

				&widgets.SizedBox{Height: 16},

				// Checkbox
				&widgets.Row{
					Children: []goflow.Widget{
						widgets.NewCheckbox(is.checkboxValue, func(value bool) {
							is.checkboxValue = value
						}),
						&widgets.Text{
							Data:  fmt.Sprintf("Checkbox: %v", is.checkboxValue),
							Style: goflow.NewTextStyle(),
						},
					},
				},

				&widgets.SizedBox{Height: 16},

				// Switch
				&widgets.Row{
					Children: []goflow.Widget{
						widgets.NewSwitch(is.switchValue, func(value bool) {
							is.switchValue = value
						}),
						&widgets.Text{
							Data:  fmt.Sprintf("Switch: %v", is.switchValue),
							Style: goflow.NewTextStyle(),
						},
					},
				},

				&widgets.SizedBox{Height: 16},

				// Radio buttons
				&widgets.Text{
					Data:  "Radio Buttons:",
					Style: goflow.NewTextStyle().WithFontSize(16),
				},
				is.buildRadioGroup(),

				&widgets.SizedBox{Height: 16},

				// Slider
				&widgets.Column{
					Children: []goflow.Widget{
						&widgets.Text{
							Data:  fmt.Sprintf("Slider: %.0f", is.sliderValue),
							Style: goflow.NewTextStyle(),
						},
						widgets.NewSlider(is.sliderValue, 0, 100, func(value float64) {
							is.sliderValue = value
						}),
					},
				},
			},
		},
		Padding: goflow.NewEdgeInsetsAll(16),
	}
}

func (is *InputShowcase) buildRadioGroup() goflow.Widget {
	options := []string{"Option 1", "Option 2", "Option 3"}
	radioWidgets := make([]goflow.Widget, len(options))

	for i, option := range options {
		radioWidgets[i] = &widgets.Row{
			Children: []goflow.Widget{
				widgets.NewRadio(i, 0, func(value interface{}) {
					fmt.Printf("Selected: %v\n", value)
				}),
				&widgets.Text{
					Data:  option,
					Style: goflow.NewTextStyle(),
				},
			},
		}
	}

	return &widgets.Column{
		Children: radioWidgets,
	}
}

// AnimationShowcase demonstrates animation widgets
type AnimationShowcase struct {
	goflow.BaseWidget
	controller *animation.AnimationController
}

func NewAnimationShowcase() *AnimationShowcase {
	controller := animation.NewAnimationController(1000 * time.Millisecond)
	return &AnimationShowcase{
		controller: controller,
	}
}

func (as *AnimationShowcase) Build(context goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Child: &widgets.Column{
			Children: []goflow.Widget{
				&widgets.Text{
					Data:  "Animations",
					Style: goflow.NewTextStyle().WithFontSize(20),
				},

				&widgets.SizedBox{Height: 16},

				// Fade transition
				widgets.NewFadeTransition(
					as.controller,
					&widgets.Container{
						Width: func() *float64 {
							w := 100.0
							return &w
						}(),
						Height: func() *float64 {
							h := 100.0
							return &h
						}(),
						Color: goflow.NewColor(33, 150, 243, 255),
					},
				),

				&widgets.SizedBox{Height: 16},

				// Scale transition
				widgets.NewScaleTransition(
					as.controller,
					&widgets.Container{
						Width: func() *float64 {
							w := 100.0
							return &w
						}(),
						Height: func() *float64 {
							h := 100.0
							return &h
						}(),
						Color: goflow.NewColor(76, 175, 80, 255),
					},
				),

				&widgets.SizedBox{Height: 24},

				// Control buttons
				&widgets.Row{
					Children: []goflow.Widget{
						&widgets.GestureDetector{
							OnTap: func() {
								as.controller.Forward()
							},
							Child: &widgets.Container{
								Color: goflow.NewColor(33, 150, 243, 255),
								Child: &widgets.Text{
									Data: "Forward",
									Style: goflow.NewTextStyle().
										WithColor(goflow.NewColor(255, 255, 255, 255)),
								},
								Padding: goflow.NewEdgeInsets(12, 8, 12, 8),
							},
						},

						&widgets.SizedBox{Width: 8},

						&widgets.GestureDetector{
							OnTap: func() {
								as.controller.Reverse()
							},
							Child: &widgets.Container{
								Color: goflow.NewColor(33, 150, 243, 255),
								Child: &widgets.Text{
									Data: "Reverse",
									Style: goflow.NewTextStyle().
										WithColor(goflow.NewColor(255, 255, 255, 255)),
								},
								Padding: goflow.NewEdgeInsets(12, 8, 12, 8),
							},
						},
					},
				},
			},
		},
		Padding: goflow.NewEdgeInsetsAll(16),
	}
}

// DataDisplayShowcase demonstrates data display widgets
type DataDisplayShowcase struct {
	goflow.BaseWidget
}

func NewDataDisplayShowcase() *DataDisplayShowcase {
	return &DataDisplayShowcase{}
}

func (dds *DataDisplayShowcase) Build(context goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Child: &widgets.Column{
			Children: []goflow.Widget{
				&widgets.Text{
					Data:  "Data Display",
					Style: goflow.NewTextStyle().WithFontSize(20),
				},

				&widgets.SizedBox{Height: 16},

				// Card
				widgets.NewCard(
					&widgets.Container{
						Child: &widgets.Text{
							Data:  "This is a card widget",
							Style: goflow.NewTextStyle(),
						},
						Padding: goflow.NewEdgeInsetsAll(16),
					},
				).WithElevation(2.0),

				&widgets.SizedBox{Height: 16},

				// ExpansionTile
				widgets.NewExpansionTile(
					&widgets.Text{
						Data:  "Expandable Section",
						Style: goflow.NewTextStyle().WithFontSize(16),
					},
					[]goflow.Widget{
						&widgets.Text{
							Data:  "This content is hidden until expanded",
							Style: goflow.NewTextStyle(),
						},
					},
				),

				&widgets.SizedBox{Height: 16},

				// DataTable
				dds.buildDataTable(),
			},
		},
		Padding: goflow.NewEdgeInsetsAll(16),
	}
}

func (dds *DataDisplayShowcase) buildDataTable() goflow.Widget {
	columns := []widgets.DataColumn{
		{
			Label: &widgets.Text{
				Data:  "Name",
				Style: goflow.NewTextStyle(),
			},
		},
		{
			Label: &widgets.Text{
				Data:  "Age",
				Style: goflow.NewTextStyle(),
			},
			Numeric: true,
		},
		{
			Label: &widgets.Text{
				Data:  "Role",
				Style: goflow.NewTextStyle(),
			},
		},
	}

	rows := []widgets.DataRow{
		{
			Cells: []widgets.DataCell{
				{
					Child: &widgets.Text{
						Data:  "Alice",
						Style: goflow.NewTextStyle(),
					},
				},
				{
					Child: &widgets.Text{
						Data:  "30",
						Style: goflow.NewTextStyle(),
					},
				},
				{
					Child: &widgets.Text{
						Data:  "Developer",
						Style: goflow.NewTextStyle(),
					},
				},
			},
		},
		{
			Cells: []widgets.DataCell{
				{
					Child: &widgets.Text{
						Data:  "Bob",
						Style: goflow.NewTextStyle(),
					},
				},
				{
					Child: &widgets.Text{
						Data:  "25",
						Style: goflow.NewTextStyle(),
					},
				},
				{
					Child: &widgets.Text{
						Data:  "Designer",
						Style: goflow.NewTextStyle(),
					},
				},
			},
		},
	}

	return widgets.NewDataTable(columns, rows)
}

// PickerShowcase demonstrates picker widgets
type PickerShowcase struct {
	goflow.BaseWidget
	selectedColor *goflow.Color
}

func NewPickerShowcase() *PickerShowcase {
	return &PickerShowcase{
		selectedColor: goflow.NewColor(33, 150, 243, 255),
	}
}

func (ps *PickerShowcase) Build(context goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Child: &widgets.Column{
			Children: []goflow.Widget{
				&widgets.Text{
					Data:  "Pickers",
					Style: goflow.NewTextStyle().WithFontSize(20),
				},

				&widgets.SizedBox{Height: 16},

				// Color picker preview
				&widgets.Row{
					Children: []goflow.Widget{
						&widgets.Text{
							Data:  "Selected Color: ",
							Style: goflow.NewTextStyle(),
						},
						&widgets.Container{
							Width: func() *float64 {
								w := 50.0
								return &w
							}(),
							Height: func() *float64 {
								h := 50.0
								return &h
							}(),
							Color: ps.selectedColor,
						},
					},
				},

				&widgets.SizedBox{Height: 16},

				// Color picker
				widgets.NewColorPicker(ps.selectedColor, func(color *goflow.Color) {
					ps.selectedColor = color
				}),
			},
		},
		Padding: goflow.NewEdgeInsetsAll(16),
	}
}
