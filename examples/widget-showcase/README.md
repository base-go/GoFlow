# GoFlow Widget Showcase

A comprehensive demonstration of GoFlow's Phase 2 widget system implementation, showcasing all the new high-priority features from the 2025 roadmap.

## Features Demonstrated

### 1. Animation System
- **AnimationController**: Full control over animations
- **Tween Animations**: Multiple tween types (Size, Color, Alignment, EdgeInsets, Decoration)
- **CurvedAnimation**: Easing curves (Linear, EaseIn, EaseOut, EaseInOut, Bounce, Elastic)
- **Transitions**: FadeTransition, ScaleTransition, RotationTransition, SlideTransition
- **Hero Animations**: Shared element transitions between routes
- **Page Transitions**: Material, Cupertino, Fade, Slide, Zoom, Scale

### 2. Form System
- **Form Widget**: Complete form validation system
- **FormField States**: Validation, save, reset functionality
- **Auto-validation**: Configurable validation modes
- **Validators**:
  - Required field validator
  - Email validator
  - Min/max length validators
  - Custom validators
  - Composed validators

### 3. Form Fields
- **TextFormField**: Text input with validation
- **DropdownFormField**: Dropdown selection with validation
- **DatePickerFormField**: Date selection with validation
- **TimePickerFormField**: Time selection with validation
- **CheckboxFormField**: Checkbox with validation
- **RadioFormField**: Radio button group with validation
- **SliderFormField**: Slider with validation

### 4. Advanced Input Widgets
- **Checkbox**: Interactive checkbox with customizable colors and shapes
- **Radio**: Radio button groups with selection management
- **Switch**: Toggle switch with Material and Cupertino styles
- **Slider**: Value slider with divisions and labels
- **RangeSlider**: Dual-thumb range selection

### 5. Picker Widgets
- **DatePicker**: Calendar-based date selection
- **TimePicker**: Clock-based time selection (12h/24h formats)
- **ColorPicker**: Color palette and RGB/HSV pickers with alpha support
- **FilePicker**: Native file system dialog integration
- **DirectoryPicker**: Directory selection dialog

### 6. Data Display Widgets
- **DataTable**: Sortable, selectable data tables
  - Column sorting
  - Row selection
  - Custom cell rendering
  - Pagination support
- **Card**: Material Design cards with elevation
- **ExpansionPanel**: Collapsible panel groups
- **ExpansionTile**: Single-line expandable list items
- **TreeView**: Hierarchical data visualization
- **PaginatedDataTable**: Data tables with built-in pagination

## Running the Example

```bash
cd examples/widget-showcase
go run main.go
```

## Code Structure

```
widget-showcase/
├── main.go                 # Main application
└── README.md              # This file
```

### Application Tabs

1. **Forms Tab**: Demonstrates form validation with email and password fields
2. **Inputs Tab**: Shows all input widgets (Checkbox, Switch, Radio, Slider)
3. **Animations Tab**: Interactive animation controls and transitions
4. **Data Display Tab**: Cards, tables, and expansion panels
5. **Pickers Tab**: Color, date, time, and file pickers

## Implementation Details

### Form Validation Example

```go
formKey := widgets.NewFormKey()

form := widgets.NewForm(
    formKey,
    widgets.NewTextFormField().
        WithValidator(widgets.ComposeValidators(
            widgets.RequiredValidator("Email is required"),
            widgets.EmailValidator("Please enter a valid email"),
        )),
)

// Validate
if formKey.CurrentState().Validate() {
    formKey.CurrentState().Save()
}
```

### Animation Example

```go
controller := animation.NewAnimationController(300 * time.Millisecond)
curvedAnim := animation.NewCurvedAnimation(controller, animation.EaseInOut)

transition := widgets.NewFadeTransition(curvedAnim, myWidget)

// Control animation
controller.Forward()  // Start animation
controller.Reverse()  // Reverse animation
controller.Reset()    // Reset to beginning
```

### DataTable Example

```go
columns := []widgets.DataColumn{
    {Label: widgets.NewText("Name")},
    {Label: widgets.NewText("Age"), Numeric: true},
}

rows := []widgets.DataRow{
    {
        Cells: []widgets.DataCell{
            {Child: widgets.NewText("Alice")},
            {Child: widgets.NewText("30")},
        },
    },
}

table := widgets.NewDataTable(columns, rows).
    WithSorting(0, true)
```

### Input Widgets Example

```go
// Checkbox
checkbox := widgets.NewCheckbox(false, func(value bool) {
    fmt.Printf("Checked: %v\n", value)
})

// Slider
slider := widgets.NewSlider(50.0, 0, 100, func(value float64) {
    fmt.Printf("Value: %.0f\n", value)
})

// Switch
toggle := widgets.NewSwitch(false, func(value bool) {
    fmt.Printf("Toggled: %v\n", value)
})
```

## Phase 2 Roadmap Completion

This example demonstrates complete implementation of Phase 2 high-priority features:

- ✅ **2.1 Animation System**: AnimationController, Tweens, Curves, Transitions, Hero
- ✅ **2.2 Form System**: Form validation, auto-validation, custom validators
- ✅ **2.3 Advanced Inputs**: Checkbox, Radio, Switch, Slider, RangeSlider
- ✅ **2.4 Data Display**: DataTable, Card, ExpansionPanel, TreeView

## Next Steps

After running this example, explore:
- Creating custom validators for your forms
- Building complex animations with multiple tweens
- Implementing sortable and paginated data tables
- Customizing widget themes and colors

## Resources

- [GoFlow Documentation](../../docs/README.md)
- [ROADMAP](../../ROADMAP.md) - Full development roadmap
- [Widget Reference](../../docs/WIDGETS_REFERENCE.md) - Complete widget catalog
