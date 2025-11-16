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
