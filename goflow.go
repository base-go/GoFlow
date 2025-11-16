// Package goflow provides a declarative UI framework for Go, inspired by Flutter.
// It offers a single import path for all framework functionality.
//
// Example usage:
//
//	import gf "github.com/base-go/GoFlow"
//
//	func main() {
//		count := gf.CreateSignal(0)
//
//		app := gf.Container{
//			Child: gf.Center{
//				Child: gf.Column{
//					Children: []gf.Widget{
//						gf.Text{Content: "Counter Example"},
//						gf.Text{Content: fmt.Sprintf("Count: %d", count.Get())},
//					},
//				},
//			},
//		}
//
//		gf.RunApp(app)
//	}
package goflow

import (
	"github.com/base-go/GoFlow/pkg/core/animation"
	framework "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
	"github.com/base-go/GoFlow/pkg/hotreload"
	"github.com/base-go/GoFlow/pkg/input"
	"github.com/base-go/GoFlow/pkg/navigation"
	"github.com/base-go/GoFlow/pkg/testing"
	"github.com/base-go/GoFlow/pkg/ui/adaptive"
	"github.com/base-go/GoFlow/pkg/ui/cupertino"
	"github.com/base-go/GoFlow/pkg/ui/material"
)

// ============================================================================
// Core Framework Types
// ============================================================================

// Widget is the base interface for all UI components.
type Widget = framework.Widget

// Element represents a widget instance in the widget tree.
type Element = framework.Element

// BuildContext provides context information during widget builds.
type BuildContext = framework.BuildContext

// RenderObject handles layout and painting operations.
type RenderObject = framework.RenderObject

// App represents the application root.
type App = framework.App

// BaseWidget provides a base implementation for stateless widgets.
type BaseWidget = framework.BaseWidget

// BaseElement provides a base implementation for elements.
type BaseElement = framework.BaseElement

// BaseRenderBox provides a base implementation for render objects.
type BaseRenderBox = framework.BaseRenderBox

// ============================================================================
// Geometry Types
// ============================================================================

// Size represents 2D dimensions (width, height).
type Size = framework.Size

// Offset represents a 2D position (x, y).
type Offset = framework.Offset

// Rect represents a rectangle (position + size).
type Rect = framework.Rect

// Constraints define layout constraints for widgets.
type Constraints = framework.Constraints

// EdgeInsets represents spacing/padding on all four sides.
type EdgeInsets = framework.EdgeInsets

// ============================================================================
// Styling Types
// ============================================================================

// Color represents an RGBA color.
type Color = framework.Color

// Paint defines drawing properties (color, stroke, etc.).
type Paint = framework.Paint

// TextStyle defines text appearance (font, size, color, etc.).
type TextStyle = framework.TextStyle

// Canvas provides a drawing surface.
type Canvas = framework.Canvas

// ============================================================================
// Color Constants
// ============================================================================

var (
	ColorBlack       = framework.ColorBlack
	ColorWhite       = framework.ColorWhite
	ColorRed         = framework.ColorRed
	ColorGreen       = framework.ColorGreen
	ColorBlue        = framework.ColorBlue
	ColorYellow      = framework.ColorYellow
	ColorCyan        = framework.ColorCyan
	ColorMagenta     = framework.ColorMagenta
	ColorGray        = framework.ColorGray
	ColorTransparent = framework.ColorTransparent
)

// ============================================================================
// Platform Constants
// ============================================================================

const (
	PlatformAndroid = framework.PlatformAndroid
	PlatformIOS     = framework.PlatformIOS
	PlatformMacOS   = framework.PlatformMacOS
	PlatformLinux   = framework.PlatformLinux
	PlatformWindows = framework.PlatformWindows
	PlatformWeb     = framework.PlatformWeb
	PlatformFuchsia = framework.PlatformFuchsia
)

// ============================================================================
// Core Framework Functions
// ============================================================================

// RunApp is the main entry point to start a GoFlow application.
var RunApp = framework.RunApp

// NewApp creates a new application instance with specified dimensions.
var NewApp = framework.NewApp

// Platform detection and management
var (
	DetectPlatform = framework.DetectPlatform
	GetPlatform    = framework.GetPlatform
	SetPlatform    = framework.SetPlatform
)

// Constraint helpers
var (
	TightConstraints     = framework.TightConstraints
	LooseConstraints     = framework.LooseConstraints
	UnboundedConstraints = framework.UnboundedConstraints
)

// Geometry helpers
var (
	NewSize   = framework.NewSize
	NewOffset = framework.NewOffset
	NewRect   = framework.NewRect
	ZeroSize  = framework.ZeroSize
	ZeroOffset = framework.ZeroOffset
)

// Styling helpers
var (
	NewColor     = framework.NewColor
	NewPaint     = framework.NewPaint
	NewTextStyle = framework.NewTextStyle
	NewEdgeInsets = framework.NewEdgeInsets
)

// ============================================================================
// Layout Widgets
// ============================================================================

// Column arranges widgets vertically.
type Column = widgets.Column

// Row arranges widgets horizontally.
type Row = widgets.Row

// Stack overlays widgets on top of each other.
type Stack = widgets.Stack

// Wrap flows widgets and wraps to the next line when needed.
type Wrap = widgets.Wrap

// Center centers its child widget.
type Center = widgets.Center

// Align positions its child within itself.
type Align = widgets.Align

// Container is a multi-purpose widget combining common styling and layout.
type Container = widgets.Container

// Padding adds padding around its child.
type Padding = widgets.Padding

// SizedBox creates a fixed-size box.
type SizedBox = widgets.SizedBox

// Flexible controls how a child flexes in a Row or Column.
type Flexible = widgets.Flexible

// Expanded makes a child fill available space in a Row or Column.
type Expanded = widgets.Expanded

// ============================================================================
// Display Widgets
// ============================================================================

// Text displays a string of text.
type Text = widgets.Text

// Icon displays an icon.
type Icon = widgets.Icon

// Image displays an image.
type Image = widgets.Image

// Avatar displays a profile picture or placeholder.
type Avatar = widgets.Avatar

// CircleAvatar displays a circular profile picture.
type CircleAvatar = widgets.CircleAvatar

// Badge displays a notification badge.
type Badge = widgets.Badge

// Chip displays a material chip (compact element).
type Chip = widgets.Chip

// ActionChip displays an action chip.
type ActionChip = widgets.ActionChip

// Divider displays a visual separator line.
type Divider = widgets.Divider

// ============================================================================
// Input Widgets
// ============================================================================

// TextField is a text input widget.
type TextField = widgets.TextField

// Checkbox is a boolean selection widget.
type Checkbox = widgets.Checkbox

// Radio is a single selection widget within a group.
type Radio = widgets.Radio

// Switch is a toggle switch widget.
type Switch = widgets.Switch

// Slider is a range selection widget.
type Slider = widgets.Slider

// Dropdown is a dropdown menu widget.
type Dropdown = widgets.Dropdown

// Form manages form fields and validation.
type Form = widgets.Form

// FormField represents a single form field.
type FormField = widgets.FormField

// GestureDetector recognizes gestures on its child.
type GestureDetector = widgets.GestureDetector

// ============================================================================
// Data Display Widgets
// ============================================================================

// ListView displays a scrollable list of widgets.
type ListView = widgets.ListView

// DataTable displays data in a table format.
type DataTable = widgets.DataTable

// PaginatedDataTable displays a paginated data table.
type PaginatedDataTable = widgets.PaginatedDataTable

// ExpansionPanel displays expandable/collapsible panels.
type ExpansionPanel = widgets.ExpansionPanel

// ExpansionTile displays an expandable list tile.
type ExpansionTile = widgets.ExpansionTile

// TreeView displays hierarchical data in a tree structure.
type TreeView = widgets.TreeView

// Card displays content in a material card.
type Card = widgets.Card

// ============================================================================
// Animation Widgets
// ============================================================================

// AnimatedContainer animates container property changes.
type AnimatedContainer = widgets.AnimatedContainer

// AnimatedOpacity animates opacity changes.
type AnimatedOpacity = widgets.AnimatedOpacity

// AnimatedSize animates size changes.
type AnimatedSize = widgets.AnimatedSize

// FadeTransition animates opacity with an animation controller.
type FadeTransition = widgets.FadeTransition

// ScaleTransition animates scale with an animation controller.
type ScaleTransition = widgets.ScaleTransition

// SlideTransition animates position with an animation controller.
type SlideTransition = widgets.SlideTransition

// RotationTransition animates rotation with an animation controller.
type RotationTransition = widgets.RotationTransition

// AnimatedBuilder builds widgets based on animation values.
type AnimatedBuilder = widgets.AnimatedBuilder

// Hero creates shared element transitions between routes.
type Hero = widgets.Hero

// ============================================================================
// Other Widgets
// ============================================================================

// PageView displays swipeable pages.
type PageView = widgets.PageView

// ScrollController controls scrollable widgets.
type ScrollController = widgets.ScrollController

// FocusNode represents a focus target.
type FocusNode = widgets.FocusNode

// FocusScope manages focus within a subtree.
type FocusScope = widgets.FocusScope

// DatePicker displays a date picker.
type DatePicker = widgets.DatePicker

// TimePicker displays a time picker.
type TimePicker = widgets.TimePicker

// ColorPicker displays a color picker.
type ColorPicker = widgets.ColorPicker

// ============================================================================
// Layout Alignment Constants
// ============================================================================

// MainAxisAlignment constants
const (
	MainAxisStart        = widgets.MainAxisStart
	MainAxisEnd          = widgets.MainAxisEnd
	MainAxisCenter       = widgets.MainAxisCenter
	MainAxisSpaceBetween = widgets.MainAxisSpaceBetween
	MainAxisSpaceAround  = widgets.MainAxisSpaceAround
	MainAxisSpaceEvenly  = widgets.MainAxisSpaceEvenly
)

// CrossAxisAlignment constants
const (
	CrossAxisStart   = widgets.CrossAxisStart
	CrossAxisEnd     = widgets.CrossAxisEnd
	CrossAxisCenter  = widgets.CrossAxisCenter
	CrossAxisStretch = widgets.CrossAxisStretch
)

// MainAxisSize constants
const (
	MainAxisSizeMin = widgets.MainAxisSizeMin
	MainAxisSizeMax = widgets.MainAxisSizeMax
)

// ============================================================================
// Reactive State Management (Signals)
// ============================================================================

// Signal is a reactive value container that notifies listeners on changes.
type Signal[T any] = signals.Signal[T]

// Computed is a derived reactive value that recomputes when dependencies change.
type Computed[T any] = signals.Computed[T]

// SignalSlice is a reactive slice.
type SignalSlice[T any] = signals.SignalSlice[T]

// SignalMap is a reactive map.
type SignalMap[K comparable, V any] = signals.SignalMap[K, V]

// Signal creation and management functions
var (
	// CreateSignal creates a new signal with an initial value.
	CreateSignal = signals.New

	// CreateComputed creates a computed value that derives from other signals.
	CreateComputed = signals.NewComputed

	// CreateEffect creates a side effect that runs when dependencies change.
	CreateEffect = signals.NewEffect

	// Batch batches multiple signal updates into a single notification.
	Batch = signals.Batch

	// Untracked runs a function without tracking signal dependencies.
	Untracked = signals.Untracked

	// CreateSignalSlice creates a reactive slice.
	CreateSignalSlice = signals.NewSlice

	// CreateSignalMap creates a reactive map.
	CreateSignalMap = signals.NewMap
)

// ============================================================================
// Animation System
// ============================================================================

// Animation represents an animation.
type Animation = animation.Animation

// AnimationController controls animation playback.
type AnimationController = animation.AnimationController

// AnimationStatus represents the animation state.
type AnimationStatus = animation.AnimationStatus

// NewAnimationController creates a new animation controller.
var NewAnimationController = animation.NewAnimationController

// ============================================================================
// Material Design (Material UI)
// ============================================================================

// Material widget types with "Material" prefix to avoid conflicts
type (
	MaterialAppBar                = material.AppBar
	MaterialScaffold              = material.Scaffold
	MaterialButton                = material.Button
	MaterialFloatingActionButton  = material.FloatingActionButton
	MaterialCard                  = material.Card
	MaterialDrawer                = material.Drawer
	MaterialDrawerHeader          = material.DrawerHeader
	MaterialBottomNavigationBar   = material.BottomNavigationBar
	MaterialTextField             = material.TextField
	MaterialCheckbox              = material.Checkbox
	MaterialRadio                 = material.Radio
	MaterialSwitch                = material.Switch
	MaterialSlider                = material.Slider
	MaterialDialog                = material.Dialog
	MaterialAlertDialog           = material.AlertDialog
	MaterialTheme                 = material.Theme
)

// Material theme functions
var (
	DefaultLightTheme = material.DefaultLightTheme
	DefaultDarkTheme  = material.DefaultDarkTheme
)

// ============================================================================
// Cupertino (iOS-style UI)
// ============================================================================

// Cupertino widget types with "Cupertino" prefix
type (
	CupertinoScaffold       = cupertino.Scaffold
	CupertinoNavigationBar  = cupertino.NavigationBar
	CupertinoButton         = cupertino.Button
	CupertinoTextField      = cupertino.TextField
	CupertinoCard           = cupertino.Card
	CupertinoTheme          = cupertino.Theme
)

// ============================================================================
// Adaptive (Cross-platform UI)
// ============================================================================

// Adaptive widget types with "Adaptive" prefix
type (
	AdaptiveButton = adaptive.Button
	AdaptiveAppBar = adaptive.AppBar
	AdaptiveCard   = adaptive.Card
)

// ============================================================================
// Navigation (GetX-style)
// ============================================================================

// Get is the global navigation instance providing navigation methods.
var Get = navigation.Get

// Navigation types
type (
	Navigator    = navigation.Navigator
	Route        = navigation.Route
	RouteBuilder = navigation.RouteBuilder
)

// Navigation app constructors
var (
	GetApp          = navigation.GetApp
	GetMaterialApp  = navigation.GetMaterialApp
	GetCupertinoApp = navigation.GetCupertinoApp
)

// Navigation helper functions
var (
	ShowDialog            = navigation.ShowDialog
	ShowAlertDialog       = navigation.ShowAlertDialog
	ShowBottomSheet       = navigation.ShowBottomSheet
	ShowModalBottomSheet  = navigation.ShowModalBottomSheet
	ShowSnackbar          = navigation.ShowSnackbar
)

// Navigation transition constants
const (
	TransitionFade       = navigation.TransitionFade
	TransitionSlideRight = navigation.TransitionSlideRight
	TransitionSlideLeft  = navigation.TransitionSlideLeft
	TransitionSlideUp    = navigation.TransitionSlideUp
	TransitionSlideDown  = navigation.TransitionSlideDown
	TransitionZoom       = navigation.TransitionZoom
)

// ============================================================================
// Input System
// ============================================================================

// Input event types
type (
	InputEvent           = input.InputEvent
	KeyboardEvent        = input.KeyboardEvent
	MouseEvent           = input.MouseEvent
	TouchEvent           = input.TouchEvent
	EventDispatcher      = input.EventDispatcher
	GestureRecognizer    = input.GestureRecognizer
	MultiTouchRecognizer = input.MultiTouchRecognizer
)

// Input functions
var (
	NewEventDispatcher      = input.NewEventDispatcher
	NewKeyboardState        = input.NewKeyboardState
	NewMouseState           = input.NewMouseState
	NewMultiTouchRecognizer = input.NewMultiTouchRecognizer
)

// ============================================================================
// Hot Reload (Development)
// ============================================================================

// HotReloader manages hot reload during development.
type HotReloader = hotreload.HotReloader

// StateStore preserves state across hot reloads.
type StateStore = hotreload.StateStore

// Hot reload functions
var (
	NewHotReloader = hotreload.NewHotReloader
	NewStateStore  = hotreload.NewStateStore
)

// ============================================================================
// Testing
// ============================================================================

// Testing types
type (
	WidgetTester      = testing.WidgetTester
	GoldenTester      = testing.GoldenTester
	IntegrationTester = testing.IntegrationTester
	GestureSimulator  = testing.GestureSimulator
)

// Testing functions
var (
	NewWidgetTester      = testing.NewWidgetTester
	NewGoldenTester      = testing.NewGoldenTester
	NewIntegrationTester = testing.NewIntegrationTester
)
