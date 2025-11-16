package widgets

import (
	"github.com/base-go/GoFlow/goflow"
)

// Image displays an image
type Image struct {
	goflow.BaseWidget

	// Image source
	Source ImageSource

	// How to inscribe the image into the space
	Fit ImageFit

	// Width and height
	Width  *float64
	Height *float64

	// Semantic label for accessibility
	SemanticLabel string
}

// ImageSource defines where the image comes from
type ImageSource interface {
	GetPath() string
}

// AssetImage loads an image from assets
type AssetImage struct {
	Path string
}

func (a *AssetImage) GetPath() string {
	return a.Path
}

// NetworkImage loads an image from a URL
type NetworkImage struct {
	URL string
}

func (n *NetworkImage) GetPath() string {
	return n.URL
}

// FileImage loads an image from a file path
type FileImage struct {
	Path string
}

func (f *FileImage) GetPath() string {
	return f.Path
}

// ImageFit defines how to inscribe the image
type ImageFit int

const (
	ImageFitFill ImageFit = iota   // Fill the box, distorting aspect ratio
	ImageFitContain                // Contain within box, preserving aspect ratio
	ImageFitCover                  // Cover the box, preserving aspect ratio
	ImageFitFitWidth               // Fit width, height may overflow
	ImageFitFitHeight              // Fit height, width may overflow
	ImageFitNone                   // No scaling
	ImageFitScaleDown              // Align within box, scale down if needed
)

// NewImage creates a new Image widget
func NewImage(source ImageSource) *Image {
	return &Image{
		Source: source,
		Fit:    ImageFitContain,
	}
}

// CreateElement creates a render object element
func (i *Image) CreateElement() goflow.Element {
	return &RenderObjectElement{
		BaseElement: goflow.BaseElement{},
		widget:      i,
		renderObject: &RenderImage{
			BaseRenderBox: goflow.NewBaseRenderBox(),
			source:        i.Source,
			fit:           i.Fit,
			width:         i.Width,
			height:        i.Height,
		},
	}
}

// Build returns nil (this is a render object widget)
func (i *Image) Build(context goflow.BuildContext) goflow.Widget {
	return nil
}

// RenderImage is the render object for Image
type RenderImage struct {
	*goflow.BaseRenderBox
	source ImageSource
	fit    ImageFit
	width  *float64
	height *float64
}

// PerformLayout performs layout
func (r *RenderImage) PerformLayout() {
	constraints := r.GetConstraints()

	// Determine size
	var width, height float64

	if r.width != nil {
		width = *r.width
	} else {
		width = constraints.MaxWidth
	}

	if r.height != nil {
		height = *r.height
	} else {
		height = constraints.MaxHeight
	}

	r.SetSize(constraints.Constrain(goflow.NewSize(width, height)))
}

// Layout performs the layout
func (r *RenderImage) Layout(constraints *goflow.Constraints) {
	r.BaseRenderBox.Layout(constraints)
	r.PerformLayout()
}

// Paint paints the image
func (r *RenderImage) Paint(canvas goflow.Canvas, offset *goflow.Offset) {
	// In a real implementation, this would draw the actual image
	// For now, we'll just draw a placeholder rectangle
	paint := goflow.NewPaint()
	paint.Color = goflow.NewColor(200, 200, 200, 255)
	rect := goflow.NewRect(offset, r.GetSize())
	canvas.DrawRect(rect, paint)
}
