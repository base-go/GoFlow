# Cupertino Widgets

iOS and macOS-specific widgets following Apple's Human Interface Guidelines.

## CupertinoTextField

iOS-style text input field with rounded corners.

```go
textField := cupertino.NewCupertinoTextField()
textField.Placeholder = "Enter text"
```

### Properties

- `Controller` - TextEditingController for managing text
- `Placeholder` - Placeholder text when empty
- `Prefix` - Leading widget (usually icon)
- `Suffix` - Trailing widget
- `Style` - Text style
- `ObscureText` - Hide text (for passwords)
- `Enabled` - Whether field is enabled (default: true)
- `OnChanged` - Callback when text changes
- `OnSubmitted` - Callback when submitted

### Basic TextField

```go
controller := material.NewTextEditingController("")
textField := cupertino.NewCupertinoTextField()
textField.Controller = controller
textField.Placeholder = "Email"
```

### TextField with Prefix Icon

```go
textField := cupertino.NewCupertinoTextField()
textField.Placeholder = "Search"
textField.Prefix = &widgets.Container{
    Padding: goflow.NewEdgeInsets(0, 8, 0, 0),
    Child: widgets.NewIcon(widgets.IconSearch),
}
```

### Password Field

```go
passwordField := cupertino.NewCupertinoTextField()
passwordField.Placeholder = "Password"
passwordField.ObscureText = true
passwordField.Prefix = &widgets.Container{
    Padding: goflow.NewEdgeInsets(0, 8, 0, 0),
    Child: widgets.NewIcon(widgets.IconWarning),
}
```

### Example: Form Field with Label

```go
func buildFormField(label, placeholder string, controller *material.TextEditingController) goflow.Widget {
    labelStyle := goflow.NewTextStyle()
    labelStyle.FontSize = 14.0
    labelStyle.Color = goflow.NewColor(142, 142, 147, 255) // iOS gray

    textField := cupertino.NewCupertinoTextField()
    textField.Controller = controller
    textField.Placeholder = placeholder

    return &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{
                Data: label,
                Style: labelStyle,
            },
            &widgets.Container{Height: floatPtr(8.0)},
            textField,
        },
        CrossAxisAlign: widgets.CrossAxisStart,
    }
}
```

### Example: Search Field

```go
func buildSearchField() goflow.Widget {
    searchController := material.NewTextEditingController("")

    textField := cupertino.NewCupertinoTextField()
    textField.Controller = searchController
    textField.Placeholder = "Search"

    textField.Prefix = &widgets.Container{
        Padding: goflow.NewEdgeInsets(0, 8, 0, 0),
        Child: widgets.NewIcon(widgets.IconSearch),
    }

    textField.Suffix = widgets.NewIconButton(
        widgets.IconClose,
        func() {
            searchController.Text.Set("")
        },
    )

    textField.OnChanged = func(text string) {
        performSearch(text)
    }

    return textField
}
```

### Example: Reactive Validation

```go
emailController := material.NewTextEditingController("")
isValid := signals.NewComputed(func() bool {
    email := emailController.Text.Get()
    return strings.Contains(email, "@")
})

textField := cupertino.NewCupertinoTextField()
textField.Controller = emailController
textField.Placeholder = "email@example.com"

textField.Suffix = func() goflow.Widget {
    if isValid.Get() {
        return widgets.NewIcon(widgets.IconCheck)
    }
    return nil
}()
```

## CupertinoSwitch

iOS-style toggle switch.

```go
switchWidget := cupertino.NewCupertinoSwitch(
    false,
    func(value bool) {
        fmt.Printf("Switch: %v\n", value)
    },
)
```

### Properties

- `Value` - Current state (true/false)
- `OnChanged` - Callback when toggled
- `ActiveColor` - Color when on (default: iOS blue)

### Basic Switch

```go
isEnabled := signals.New(false)

switchWidget := cupertino.NewCupertinoSwitch(
    isEnabled.Get(),
    func(value bool) {
        isEnabled.Set(value)
    },
)
```

### Styled Switch

```go
switchWidget := cupertino.NewCupertinoSwitch(true, onChanged)
switchWidget.ActiveColor = goflow.NewColor(52, 199, 89, 255) // iOS green
```

### Example: Settings Toggle

```go
func buildSettingToggle(title string, value *signals.Signal[bool]) goflow.Widget {
    switchWidget := cupertino.NewCupertinoSwitch(
        value.Get(),
        func(newValue bool) {
            value.Set(newValue)
        },
    )

    return &widgets.Row{
        Children: []goflow.Widget{
            &widgets.Expanded{
                Child: &widgets.Text{Data: title},
            },
            switchWidget,
        },
        MainAxisAlignment: widgets.MainAxisSpaceBetween,
    }
}

// Usage
notificationsEnabled := signals.New(true)
buildSettingToggle("Notifications", notificationsEnabled)
```

### Example: Multiple Switches

```go
type Setting struct {
    Label   string
    Enabled *signals.Signal[bool]
}

func buildSettingsList(settings []Setting) goflow.Widget {
    items := make([]goflow.Widget, len(settings))

    for i, setting := range settings {
        items[i] = &widgets.Container{
            Padding: goflow.NewEdgeInsets(16, 12, 16, 12),
            Child: &widgets.Row{
                Children: []goflow.Widget{
                    &widgets.Expanded{
                        Child: &widgets.Text{Data: setting.Label},
                    },
                    cupertino.NewCupertinoSwitch(
                        setting.Enabled.Get(),
                        func(value bool) {
                            setting.Enabled.Set(value)
                        },
                    ),
                },
            },
        }
    }

    return &widgets.ListView{
        Children: items,
    }
}
```

## CupertinoSlider

iOS-style slider for selecting values.

```go
slider := cupertino.NewCupertinoSlider(
    50.0,    // value
    0.0,     // min
    100.0,   // max
    func(value float64) {
        fmt.Printf("Value: %.0f\n", value)
    },
)
```

### Properties

- `Value` - Current value
- `Min` - Minimum value
- `Max` - Maximum value
- `Divisions` - Number of discrete steps (optional)
- `OnChanged` - Callback when value changes
- `ActiveColor` - Color of active portion (default: iOS blue)

### Basic Slider

```go
volume := signals.New(50.0)

slider := cupertino.NewCupertinoSlider(
    volume.Get(),
    0.0,
    100.0,
    func(value float64) {
        volume.Set(value)
    },
)
```

### Styled Slider

```go
slider := cupertino.NewCupertinoSlider(value, min, max, onChanged)
slider.ActiveColor = goflow.NewColor(255, 59, 48, 255) // iOS red
```

### Discrete Slider

```go
divisions := 10
slider := cupertino.NewCupertinoSlider(5.0, 0.0, 10.0, onChanged)
slider.Divisions = &divisions // Snap to integer values
```

### Example: Volume Control

```go
func buildVolumeControl() goflow.Widget {
    volume := signals.New(50.0)

    slider := cupertino.NewCupertinoSlider(
        volume.Get(),
        0.0,
        100.0,
        func(value float64) {
            volume.Set(value)
            setVolume(int(value))
        },
    )

    return &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconHome), // Volume icon
                    &widgets.Expanded{Child: slider},
                    &widgets.Text{
                        Data: fmt.Sprintf("%.0f%%", volume.Get()),
                    },
                },
                MainAxisAlignment: widgets.MainAxisSpaceBetween,
            },
        },
    }
}
```

### Example: Brightness Slider

```go
func buildBrightnessSlider() goflow.Widget {
    brightness := signals.New(75.0)

    slider := cupertino.NewCupertinoSlider(
        brightness.Get(),
        0.0,
        100.0,
        func(value float64) {
            brightness.Set(value)
            setBrightness(value / 100.0)
        },
    )

    labelStyle := goflow.NewTextStyle()
    labelStyle.FontSize = 14.0
    labelStyle.Color = goflow.NewColor(142, 142, 147, 255)

    return &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Row{
                Children: []goflow.Widget{
                    &widgets.Text{
                        Data: "Brightness",
                        Style: labelStyle,
                    },
                    &widgets.Text{
                        Data: fmt.Sprintf("%.0f%%", brightness.Get()),
                    },
                },
                MainAxisAlignment: widgets.MainAxisSpaceBetween,
            },
            &widgets.Container{Height: floatPtr(8.0)},
            slider,
        },
        CrossAxisAlign: widgets.CrossAxisStretch,
    }
}
```

### Example: Color Temperature Slider

```go
func buildTemperatureSlider() goflow.Widget {
    temperature := signals.New(6500.0)

    divisions := 20
    slider := cupertino.NewCupertinoSlider(
        temperature.Get(),
        2700.0, // Warm
        10000.0, // Cool
        func(value float64) {
            temperature.Set(value)
            setColorTemperature(value)
        },
    )
    slider.Divisions = &divisions
    slider.ActiveColor = goflow.NewColor(255, 149, 0, 255) // iOS orange

    return &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{Data: "Color Temperature"},
            slider,
            &widgets.Row{
                Children: []goflow.Widget{
                    &widgets.Text{Data: "Warm"},
                    widgets.NewSpacer(),
                    &widgets.Text{Data: "Cool"},
                },
            },
        },
        CrossAxisAlign: widgets.CrossAxisStretch,
    }
}
```

## iOS Design Guidelines

### Colors

iOS uses specific system colors:

```go
// From cupertino.DefaultLightTheme()
var (
    IOSBlue   = goflow.NewColor(0, 122, 255, 255)
    IOSGreen  = goflow.NewColor(52, 199, 89, 255)
    IOSOrange = goflow.NewColor(255, 149, 0, 255)
    IOSRed    = goflow.NewColor(255, 59, 48, 255)
    IOSGray   = goflow.NewColor(142, 142, 147, 255)
)
```

### Typography

```go
// iOS font sizes
const (
    IOSFontSizeLargeTitle = 34.0
    IOSFontSizeTitle1     = 28.0
    IOSFontSizeTitle2     = 22.0
    IOSFontSizeTitle3     = 20.0
    IOSFontSizeHeadline   = 17.0
    IOSFontSizeBody       = 17.0
    IOSFontSizeCallout    = 16.0
    IOSFontSizeSubhead    = 15.0
    IOSFontSizeFootnote   = 13.0
    IOSFontSizeCaption1   = 12.0
    IOSFontSizeCaption2   = 11.0
)
```

### Spacing

```go
// iOS standard spacing
const (
    IOSPaddingSmall  = 8.0
    IOSPaddingMedium = 16.0
    IOSPaddingLarge  = 20.0
)

// iOS minimum touch target
const IOSMinTouchTarget = 44.0
```

## Best Practices

### 1. Use iOS Standard Dimensions

```go
// ✅ Good - 44pt minimum touch target
textField := cupertino.NewCupertinoTextField()
// Already has 44pt minimum height

// ✅ Good - standard corner radius (8pt for iOS)
container.Decoration = &goflow.BoxDecoration{
    BorderRadius: 8.0,
}
```

### 2. Match iOS Color Palette

```go
theme := cupertino.DefaultLightTheme()

// ✅ Good - use theme colors
switchWidget.ActiveColor = theme.PrimaryColor // iOS blue

// ❌ Avoid - custom colors that don't match iOS
switchWidget.ActiveColor = goflow.NewColor(128, 0, 128, 255)
```

### 3. Provide Clear Placeholders

```go
// ✅ Good - descriptive placeholder
textField.Placeholder = "email@example.com"

// ❌ Avoid - vague placeholder
textField.Placeholder = "Input"
```

### 4. Use Appropriate Input Types

```go
// Password field
passwordField := cupertino.NewCupertinoTextField()
passwordField.ObscureText = true

// Search field
searchField := cupertino.NewCupertinoTextField()
searchField.Prefix = widgets.NewIcon(widgets.IconSearch)
```

### 5. Provide Immediate Feedback

```go
switchWidget := cupertino.NewCupertinoSwitch(value, func(newValue bool) {
    // ✅ Good - immediate update
    updateSetting(newValue)
    showToast("Setting updated")
})
```

### 6. Use Sliders for Continuous Values

```go
// ✅ Good - slider for volume
volumeSlider := cupertino.NewCupertinoSlider(50.0, 0.0, 100.0, setVolume)

// ❌ Avoid - buttons for continuous values
```

### 7. Add Units to Slider Labels

```go
// ✅ Good - shows units
&widgets.Text{
    Data: fmt.Sprintf("%.0f%%", volume.Get()),
}

// ❌ Less clear - no units
&widgets.Text{
    Data: fmt.Sprintf("%.0f", volume.Get()),
}
```

### 8. Group Related Settings

```go
func buildSettingsGroup(title string, settings []Setting) goflow.Widget {
    titleStyle := goflow.NewTextStyle()
    titleStyle.FontSize = 13.0
    titleStyle.Color = goflow.NewColor(142, 142, 147, 255)

    return &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Container{
                Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
                Child: &widgets.Text{
                    Data: title,
                    Style: titleStyle,
                },
            },
            buildSettingsList(settings),
        },
        CrossAxisAlign: widgets.CrossAxisStart,
    }
}
```

## Common Patterns

### iOS Form

```go
func buildIOSForm() goflow.Widget {
    nameController := material.NewTextEditingController("")
    emailController := material.NewTextEditingController("")

    return &widgets.Column{
        Children: []goflow.Widget{
            buildFormField("Name", "John Doe", nameController),
            &widgets.Container{Height: floatPtr(16.0)},
            buildFormField("Email", "email@example.com", emailController),
            &widgets.Container{Height: floatPtr(24.0)},
            cupertino.NewButton(
                &widgets.Text{Data: "Submit"},
                func() {
                    submitForm(nameController, emailController)
                },
            ),
        },
        CrossAxisAlign: widgets.CrossAxisStretch,
    }
}
```

### iOS Settings Page

```go
func buildSettingsPage() goflow.Widget {
    return &cupertino.CupertinoPageScaffold{
        NavigationBar: cupertino.NewNavigationBar(&widgets.Text{
            Data: "Settings",
        }),
        Child: &widgets.ListView{
            Children: []goflow.Widget{
                buildSettingsGroup("Account", accountSettings),
                buildSettingsGroup("Notifications", notificationSettings),
                buildSettingsGroup("Privacy", privacySettings),
            },
        },
    }
}
```

### iOS Control Panel

```go
func buildControlPanel() goflow.Widget {
    brightness := signals.New(75.0)
    volume := signals.New(50.0)
    wifiEnabled := signals.New(true)
    bluetoothEnabled := signals.New(false)

    return &widgets.Column{
        Children: []goflow.Widget{
            buildBrightnessSlider(),
            &widgets.Container{Height: floatPtr(16.0)},
            buildVolumeControl(),
            &widgets.Container{Height: floatPtr(24.0)},
            buildSettingToggle("Wi-Fi", wifiEnabled),
            buildSettingToggle("Bluetooth", bluetoothEnabled),
        },
        CrossAxisAlign: widgets.CrossAxisStretch,
    }
}
```

## Helper Functions

```go
// Float pointer helper
func floatPtr(f float64) *float64 {
    return &f
}

// Build iOS section header
func buildSectionHeader(title string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 13.0
    style.Color = goflow.NewColor(142, 142, 147, 255)
    style.FontWeight = goflow.FontWeightBold

    return &widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
        Child: &widgets.Text{
            Data: strings.ToUpper(title),
            Style: style,
        },
    }
}

// Build iOS divider
func buildIOSDivider() goflow.Widget {
    return &widgets.Container{
        Height: floatPtr(0.5),
        Color: goflow.NewColor(200, 199, 204, 255),
    }
}
```
