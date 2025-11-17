package screens

import (
	"fmt"

	gf "github.com/base-go/GoFlow"
)

// InputWidgetsPage showcases all input widgets
type InputWidgetsPage struct {
	gf.BaseWidget
	checkboxValue *gf.Signal[bool]
	switchValue   *gf.Signal[bool]
	sliderValue   *gf.Signal[float64]
	rangeStart    *gf.Signal[float64]
	rangeEnd      *gf.Signal[float64]
	radioValue    *gf.Signal[string]
}

func NewInputWidgetsPage() *InputWidgetsPage {
	return &InputWidgetsPage{
		checkboxValue: gf.CreateSignal(false),
		switchValue:   gf.CreateSignal(true),
		sliderValue:   gf.CreateSignal(50.0),
		rangeStart:    gf.CreateSignal(20.0),
		rangeEnd:      gf.CreateSignal(80.0),
		radioValue:    gf.CreateSignal("option1"),
	}
}

func (p *InputWidgetsPage) Build(context gf.BuildContext) gf.Widget {
	// Page header
	title := gf.Text{Data: "Input Widgets"}
	backBtn := p.buildBackButton()

	header := gf.Row{
		Children: []gf.Widget{
			backBtn,
			p.spacer(20),
			&title,
		},
	}

	// Sections
	sections := []gf.Widget{
		p.buildSection("Checkbox", p.buildCheckboxDemo()),
		p.spacer(30),
		p.buildSection("Switch", p.buildSwitchDemo()),
		p.spacer(30),
		p.buildSection("Slider", p.buildSliderDemo()),
		p.spacer(30),
		p.buildSection("Range Slider", p.buildRangeSliderDemo()),
		p.spacer(30),
		p.buildSection("Radio Buttons", p.buildRadioDemo()),
	}

	allChildren := []gf.Widget{&header, p.spacer(30)}
	allChildren = append(allChildren, sections...)

	content := gf.Column{
		Children: allChildren,
	}

	container := gf.Container{
		Padding: gf.NewEdgeInsetsAll(20.0),
		Child:   &content,
	}

	return &container
}

func (p *InputWidgetsPage) buildCheckboxDemo() gf.Widget {
	checked := p.checkboxValue.Get()

	checkbox := gf.Checkbox{
		Value: checked,
		OnChanged: func(value bool) {
			p.checkboxValue.Set(value)
			fmt.Printf("Checkbox changed: %v\n", value)
		},
	}

	label := gf.Text{Data: fmt.Sprintf("Checkbox is %v", checked)}

	row := gf.Row{
		Children: []gf.Widget{
			&checkbox,
			p.spacer(15),
			&label,
		},
	}

	return &row
}

func (p *InputWidgetsPage) buildSwitchDemo() gf.Widget {
	enabled := p.switchValue.Get()

	switchWidget := gf.Switch{
		Value: enabled,
		OnChanged: func(value bool) {
			p.switchValue.Set(value)
			fmt.Printf("Switch changed: %v\n", value)
		},
	}

	label := gf.Text{
		Data: fmt.Sprintf("Switch is %s", func() string {
			if enabled {
				return "ON"
			}
			return "OFF"
		}()),
	}

	row := gf.Row{
		Children: []gf.Widget{
			&switchWidget,
			p.spacer(15),
			&label,
		},
	}

	return &row
}

func (p *InputWidgetsPage) buildSliderDemo() gf.Widget {
	value := p.sliderValue.Get()

	slider := gf.Slider{
		Value: value,
		Min:   0.0,
		Max:   100.0,
		OnChanged: func(v float64) {
			p.sliderValue.Set(v)
			fmt.Printf("Slider value: %.1f\n", v)
		},
	}

	label := gf.Text{Data: fmt.Sprintf("Value: %.1f", value)}

	column := gf.Column{
		Children: []gf.Widget{
			&label,
			p.spacer(10),
			&slider,
		},
	}

	return &column
}

func (p *InputWidgetsPage) buildRangeSliderDemo() gf.Widget {
	start := p.rangeStart.Get()
	end := p.rangeEnd.Get()

	rangeSlider := gf.RangeSlider{
		Values: gf.RangeValues{Start: start, End: end},
		Min:    0.0,
		Max:    100.0,
		OnChanged: func(values gf.RangeValues) {
			p.rangeStart.Set(values.Start)
			p.rangeEnd.Set(values.End)
			fmt.Printf("Range: %.1f - %.1f\n", values.Start, values.End)
		},
	}

	label := gf.Text{Data: fmt.Sprintf("Range: %.1f - %.1f", start, end)}

	column := gf.Column{
		Children: []gf.Widget{
			&label,
			p.spacer(10),
			&rangeSlider,
		},
	}

	return &column
}

func (p *InputWidgetsPage) buildRadioDemo() gf.Widget {
	currentValue := p.radioValue.Get()

	options := []struct {
		value string
		label string
	}{
		{"option1", "Option 1"},
		{"option2", "Option 2"},
		{"option3", "Option 3"},
	}

	var radioButtons []gf.Widget
	for _, opt := range options {
		optValue := opt.value
		radio := gf.Radio{
			Value:     optValue,
			GroupValue: currentValue,
			OnChanged: func(value interface{}) {
				if strVal, ok := value.(string); ok {
					p.radioValue.Set(strVal)
					fmt.Printf("Radio selected: %s\n", strVal)
				}
			},
		}

		label := gf.Text{Data: opt.label}

		row := gf.Row{
			Children: []gf.Widget{
				&radio,
				p.spacer(10),
				&label,
			},
		}

		radioButtons = append(radioButtons, &row, p.spacer(10))
	}

	selectedLabel := gf.Text{Data: fmt.Sprintf("Selected: %s", currentValue)}
	allChildren := append([]gf.Widget{&selectedLabel, p.spacer(15)}, radioButtons...)

	column := gf.Column{
		Children: allChildren,
	}

	return &column
}

func (p *InputWidgetsPage) buildSection(title string, content gf.Widget) gf.Widget {
	titleWidget := gf.Text{Data: title}

	titleContainer := gf.Container{
		Padding: gf.NewEdgeInsetsSymmetric(8.0, 0),
		Child:   &titleWidget,
	}

	divider := gf.Divider{}

	column := gf.Column{
		Children: []gf.Widget{
			&titleContainer,
			p.spacer(10),
			&divider,
			p.spacer(20),
			content,
		},
	}

	return &column
}

func (p *InputWidgetsPage) buildBackButton() gf.Widget {
	text := gf.Text{Data: "← Back"}

	btn := gf.Container{
		Padding: gf.NewEdgeInsetsSymmetric(8.0, 16.0),
		Color:   gf.NewColor(200, 200, 200, 255),
		Child:   &text,
	}

	detector := gf.GestureDetector{
		Child: &btn,
		OnTap: func() {
			gf.Get.Back()
		},
	}

	return &detector
}

func (p *InputWidgetsPage) spacer(height float64) gf.Widget {
	container := gf.Container{
		Height: &height,
	}
	return &container
}

