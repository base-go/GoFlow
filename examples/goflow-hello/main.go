package main

import (
	"github.com/base-go/GoFlow/goflow"
	"github.com/base-go/GoFlow/widgets"
)

// HelloApp is a simple "Hello, GoFlow!" application
type HelloApp struct {
	goflow.BaseWidget
}

// Build creates the widget tree for the hello app
func (a *HelloApp) Build(context goflow.BuildContext) goflow.Widget {
	return widgets.NewCenter(
		&widgets.Column{
			Children: []goflow.Widget{
				widgets.NewTextWithStyle(
					"Hello, GoFlow!",
					&goflow.TextStyle{
						Color:      goflow.ColorBlue,
						FontSize:   32,
						FontWeight: goflow.FontWeightBold,
					},
				),
				widgets.NewTextWithStyle(
					"A Flutter-like GUI framework for Go",
					&goflow.TextStyle{
						Color:    goflow.ColorGray,
						FontSize: 16,
					},
				),
			},
			MainAxisAlign:  widgets.MainAxisCenter,
			CrossAxisAlign: widgets.CrossAxisCenter,
		},
	)
}

func main() {
	app := &HelloApp{}
	goflow.RunApp(app)
}
