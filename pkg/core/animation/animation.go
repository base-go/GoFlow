package animation

import (
	"math"
	"time"
)

// Animation represents an animation value over time
type Animation interface {
	// Value returns the current animation value
	Value() float64

	// AddListener adds a listener for animation changes
	AddListener(listener func())

	// RemoveListener removes a listener
	RemoveListener(listener func())

	// AddStatusListener adds a listener for animation status changes
	AddStatusListener(listener func(AnimationStatus))

	// RemoveStatusListener removes a status listener
	RemoveStatusListener(listener func(AnimationStatus))
}

// AnimationStatus represents the status of an animation
type AnimationStatus int

const (
	AnimationStatusDismissed AnimationStatus = iota
	AnimationStatusForward
	AnimationStatusReverse
	AnimationStatusCompleted
)

// AnimationController controls an animation
type AnimationController struct {
	duration         time.Duration
	value            float64
	lowerBound       float64
	upperBound       float64
	status           AnimationStatus
	listeners        []func()
	statusListeners  []func(AnimationStatus)
	lastUpdateTime   time.Time
	isAnimating      bool
	animationDirection int // 1 for forward, -1 for reverse
}

// NewAnimationController creates a new animation controller
func NewAnimationController(duration time.Duration) *AnimationController {
	return &AnimationController{
		duration:        duration,
		value:           0.0,
		lowerBound:      0.0,
		upperBound:      1.0,
		status:          AnimationStatusDismissed,
		listeners:       make([]func(), 0),
		statusListeners: make([]func(AnimationStatus), 0),
		isAnimating:     false,
	}
}

// Value returns the current animation value
func (a *AnimationController) Value() float64 {
	return a.value
}

// Status returns the current animation status
func (a *AnimationController) Status() AnimationStatus {
	return a.status
}

// IsAnimating returns true if the animation is currently running
func (a *AnimationController) IsAnimating() bool {
	return a.isAnimating
}

// Forward starts the animation forward
func (a *AnimationController) Forward() {
	a.isAnimating = true
	a.animationDirection = 1
	a.lastUpdateTime = time.Now()
	a.setStatus(AnimationStatusForward)
}

// Reverse starts the animation in reverse
func (a *AnimationController) Reverse() {
	a.isAnimating = true
	a.animationDirection = -1
	a.lastUpdateTime = time.Now()
	a.setStatus(AnimationStatusReverse)
}

// Stop stops the animation
func (a *AnimationController) Stop() {
	a.isAnimating = false
}

// Reset resets the animation to the beginning
func (a *AnimationController) Reset() {
	a.value = a.lowerBound
	a.isAnimating = false
	a.setStatus(AnimationStatusDismissed)
	a.notifyListeners()
}

// Repeat repeats the animation
func (a *AnimationController) Repeat(min, max float64) {
	// Simplified: just restart
	a.Forward()
}

// Update updates the animation (call this in your game loop)
func (a *AnimationController) Update() {
	if !a.isAnimating {
		return
	}

	now := time.Now()
	elapsed := now.Sub(a.lastUpdateTime)
	a.lastUpdateTime = now

	// Calculate new value
	delta := float64(elapsed) / float64(a.duration)
	a.value += delta * float64(a.animationDirection)

	// Check bounds
	if a.animationDirection > 0 && a.value >= a.upperBound {
		a.value = a.upperBound
		a.isAnimating = false
		a.setStatus(AnimationStatusCompleted)
	} else if a.animationDirection < 0 && a.value <= a.lowerBound {
		a.value = a.lowerBound
		a.isAnimating = false
		a.setStatus(AnimationStatusDismissed)
	}

	a.notifyListeners()
}

// AddListener adds a listener for animation changes
func (a *AnimationController) AddListener(listener func()) {
	a.listeners = append(a.listeners, listener)
}

// RemoveListener removes a listener
func (a *AnimationController) RemoveListener(listener func()) {
	for i, l := range a.listeners {
		if &l == &listener {
			a.listeners = append(a.listeners[:i], a.listeners[i+1:]...)
			return
		}
	}
}

// AddStatusListener adds a listener for animation status changes
func (a *AnimationController) AddStatusListener(listener func(AnimationStatus)) {
	a.statusListeners = append(a.statusListeners, listener)
}

// RemoveStatusListener removes a status listener
func (a *AnimationController) RemoveStatusListener(listener func(AnimationStatus)) {
	for i, l := range a.statusListeners {
		if &l == &listener {
			a.statusListeners = append(a.statusListeners[:i], a.statusListeners[i+1:]...)
			return
		}
	}
}

func (a *AnimationController) notifyListeners() {
	for _, listener := range a.listeners {
		listener()
	}
}

func (a *AnimationController) setStatus(status AnimationStatus) {
	if a.status != status {
		a.status = status
		for _, listener := range a.statusListeners {
			listener(status)
		}
	}
}

// Dispose cleans up the controller
func (a *AnimationController) Dispose() {
	a.listeners = nil
	a.statusListeners = nil
}

// Tween interpolates between two values
type Tween struct {
	begin float64
	end   float64
}

// NewTween creates a new tween
func NewTween(begin, end float64) *Tween {
	return &Tween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end
func (t *Tween) Lerp(value float64) float64 {
	return t.begin + (t.end-t.begin)*value
}

// ColorTween interpolates between two colors
type ColorTween struct {
	begin *Color
	end   *Color
}

// Color represents an RGBA color
type Color struct {
	R, G, B, A uint8
}

// NewColorTween creates a new color tween
func NewColorTween(begin, end *Color) *ColorTween {
	return &ColorTween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end colors
func (c *ColorTween) Lerp(value float64) *Color {
	return &Color{
		R: uint8(float64(c.begin.R) + float64(c.end.R-c.begin.R)*value),
		G: uint8(float64(c.begin.G) + float64(c.end.G-c.begin.G)*value),
		B: uint8(float64(c.begin.B) + float64(c.end.B-c.begin.B)*value),
		A: uint8(float64(c.begin.A) + float64(c.end.A-c.begin.A)*value),
	}
}

// Curve defines an easing curve
type Curve interface {
	Transform(t float64) float64
}

// LinearCurve is a linear curve (no easing)
type LinearCurve struct{}

func (l *LinearCurve) Transform(t float64) float64 {
	return t
}

// EaseInCurve eases in (slow start)
type EaseInCurve struct{}

func (e *EaseInCurve) Transform(t float64) float64 {
	return t * t
}

// EaseOutCurve eases out (slow end)
type EaseOutCurve struct{}

func (e *EaseOutCurve) Transform(t float64) float64 {
	return 1 - (1-t)*(1-t)
}

// EaseInOutCurve eases in and out
type EaseInOutCurve struct{}

func (e *EaseInOutCurve) Transform(t float64) float64 {
	if t < 0.5 {
		return 2 * t * t
	}
	return 1 - 2*(1-t)*(1-t)
}

// BounceInCurve bounces in
type BounceInCurve struct{}

func (b *BounceInCurve) Transform(t float64) float64 {
	return 1 - (&BounceOutCurve{}).Transform(1-t)
}

// BounceOutCurve bounces out
type BounceOutCurve struct{}

func (b *BounceOutCurve) Transform(t float64) float64 {
	if t < 1/2.75 {
		return 7.5625 * t * t
	} else if t < 2/2.75 {
		t -= 1.5 / 2.75
		return 7.5625*t*t + 0.75
	} else if t < 2.5/2.75 {
		t -= 2.25 / 2.75
		return 7.5625*t*t + 0.9375
	} else {
		t -= 2.625 / 2.75
		return 7.5625*t*t + 0.984375
	}
}

// ElasticInCurve elastic in
type ElasticInCurve struct{}

func (e *ElasticInCurve) Transform(t float64) float64 {
	if t == 0 || t == 1 {
		return t
	}
	return -math.Pow(2, 10*(t-1)) * math.Sin((t-1.1)*5*math.Pi)
}

// ElasticOutCurve elastic out
type ElasticOutCurve struct{}

func (e *ElasticOutCurve) Transform(t float64) float64 {
	if t == 0 || t == 1 {
		return t
	}
	return math.Pow(2, -10*t)*math.Sin((t-0.1)*5*math.Pi) + 1
}

// CurvedAnimation wraps an animation with a curve
type CurvedAnimation struct {
	parent Animation
	curve  Curve
}

// NewCurvedAnimation creates a new curved animation
func NewCurvedAnimation(parent Animation, curve Curve) *CurvedAnimation {
	return &CurvedAnimation{
		parent: parent,
		curve:  curve,
	}
}

// Value returns the curved value
func (c *CurvedAnimation) Value() float64 {
	return c.curve.Transform(c.parent.Value())
}

// AddListener adds a listener
func (c *CurvedAnimation) AddListener(listener func()) {
	c.parent.AddListener(listener)
}

// RemoveListener removes a listener
func (c *CurvedAnimation) RemoveListener(listener func()) {
	c.parent.RemoveListener(listener)
}

// AddStatusListener adds a status listener
func (c *CurvedAnimation) AddStatusListener(listener func(AnimationStatus)) {
	c.parent.AddStatusListener(listener)
}

// RemoveStatusListener removes a status listener
func (c *CurvedAnimation) RemoveStatusListener(listener func(AnimationStatus)) {
	c.parent.RemoveStatusListener(listener)
}

// SizeTween interpolates between two sizes
type SizeTween struct {
	begin *Size
	end   *Size
}

// Size represents width and height
type Size struct {
	Width  float64
	Height float64
}

// NewSizeTween creates a new size tween
func NewSizeTween(begin, end *Size) *SizeTween {
	return &SizeTween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end sizes
func (s *SizeTween) Lerp(value float64) *Size {
	return &Size{
		Width:  s.begin.Width + (s.end.Width-s.begin.Width)*value,
		Height: s.begin.Height + (s.end.Height-s.begin.Height)*value,
	}
}

// AlignmentTween interpolates between two alignments
type AlignmentTween struct {
	begin *Alignment
	end   *Alignment
}

// Alignment represents alignment values
type Alignment struct {
	X float64 // -1.0 (left) to 1.0 (right)
	Y float64 // -1.0 (top) to 1.0 (bottom)
}

// NewAlignmentTween creates a new alignment tween
func NewAlignmentTween(begin, end *Alignment) *AlignmentTween {
	return &AlignmentTween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end alignments
func (a *AlignmentTween) Lerp(value float64) *Alignment {
	return &Alignment{
		X: a.begin.X + (a.end.X-a.begin.X)*value,
		Y: a.begin.Y + (a.end.Y-a.begin.Y)*value,
	}
}

// EdgeInsetsTween interpolates between two edge insets
type EdgeInsetsTween struct {
	begin *EdgeInsets
	end   *EdgeInsets
}

// EdgeInsets represents padding/margin values
type EdgeInsets struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// NewEdgeInsetsTween creates a new edge insets tween
func NewEdgeInsetsTween(begin, end *EdgeInsets) *EdgeInsetsTween {
	return &EdgeInsetsTween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end edge insets
func (e *EdgeInsetsTween) Lerp(value float64) *EdgeInsets {
	return &EdgeInsets{
		Top:    e.begin.Top + (e.end.Top-e.begin.Top)*value,
		Right:  e.begin.Right + (e.end.Right-e.begin.Right)*value,
		Bottom: e.begin.Bottom + (e.end.Bottom-e.begin.Bottom)*value,
		Left:   e.begin.Left + (e.end.Left-e.begin.Left)*value,
	}
}

// BorderRadiusTween interpolates between two border radius values
type BorderRadiusTween struct {
	begin float64
	end   float64
}

// NewBorderRadiusTween creates a new border radius tween
func NewBorderRadiusTween(begin, end float64) *BorderRadiusTween {
	return &BorderRadiusTween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end border radius
func (b *BorderRadiusTween) Lerp(value float64) float64 {
	return b.begin + (b.end-b.begin)*value
}

// DecorationTween interpolates between two box decorations
type DecorationTween struct {
	begin *BoxDecoration
	end   *BoxDecoration
}

// BoxDecoration represents container decoration
type BoxDecoration struct {
	Color        *Color
	BorderRadius float64
	Border       *Border
}

// Border represents border properties
type Border struct {
	Color *Color
	Width float64
}

// NewDecorationTween creates a new decoration tween
func NewDecorationTween(begin, end *BoxDecoration) *DecorationTween {
	return &DecorationTween{
		begin: begin,
		end:   end,
	}
}

// Lerp linearly interpolates between begin and end decorations
func (d *DecorationTween) Lerp(value float64) *BoxDecoration {
	result := &BoxDecoration{
		BorderRadius: d.begin.BorderRadius + (d.end.BorderRadius-d.begin.BorderRadius)*value,
	}

	// Interpolate colors
	if d.begin.Color != nil && d.end.Color != nil {
		colorTween := NewColorTween(d.begin.Color, d.end.Color)
		result.Color = colorTween.Lerp(value)
	} else if d.end.Color != nil {
		result.Color = d.end.Color
	} else {
		result.Color = d.begin.Color
	}

	// Interpolate border
	if d.begin.Border != nil && d.end.Border != nil {
		borderColorTween := NewColorTween(d.begin.Border.Color, d.end.Border.Color)
		result.Border = &Border{
			Color: borderColorTween.Lerp(value),
			Width: d.begin.Border.Width + (d.end.Border.Width-d.begin.Border.Width)*value,
		}
	} else if d.end.Border != nil {
		result.Border = d.end.Border
	} else {
		result.Border = d.begin.Border
	}

	return result
}

// AnimationStatusListener is a function that listens to animation status changes
type AnimationStatusListener func(AnimationStatus)

// TickerProvider provides a ticker for animations
type TickerProvider interface {
	CreateTicker() *Ticker
}

// Ticker provides regular callbacks for animations
type Ticker struct {
	onTick   func(time.Duration)
	stopChan chan bool
	running  bool
}

// NewTicker creates a new ticker
func NewTicker(onTick func(time.Duration)) *Ticker {
	return &Ticker{
		onTick:   onTick,
		stopChan: make(chan bool),
		running:  false,
	}
}

// Start starts the ticker
func (t *Ticker) Start() {
	if t.running {
		return
	}
	t.running = true
	go func() {
		ticker := time.NewTicker(16 * time.Millisecond) // ~60 FPS
		defer ticker.Stop()
		startTime := time.Now()

		for {
			select {
			case <-ticker.C:
				elapsed := time.Since(startTime)
				if t.onTick != nil {
					t.onTick(elapsed)
				}
			case <-t.stopChan:
				return
			}
		}
	}()
}

// Stop stops the ticker
func (t *Ticker) Stop() {
	if !t.running {
		return
	}
	t.running = false
	t.stopChan <- true
}

// Common curve instances
var (
	Linear       = &LinearCurve{}
	EaseIn       = &EaseInCurve{}
	EaseOut      = &EaseOutCurve{}
	EaseInOut    = &EaseInOutCurve{}
	BounceIn     = &BounceInCurve{}
	BounceOut    = &BounceOutCurve{}
	ElasticIn    = &ElasticInCurve{}
	ElasticOut   = &ElasticOutCurve{}
)
