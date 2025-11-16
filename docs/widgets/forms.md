# Form Widgets

Form widgets allow user input and data collection.

## TextField

Text input field with Material or Cupertino styling.

### Material TextField

```go
textField := material.NewTextField()
textField.Decoration = &material.InputDecoration{
    LabelText: "Email",
    HintText: "Enter your email",
    HelperText: "We'll never share your email",
}
textField.OnChanged = func(text string) {
    fmt.Println("Text changed:", text)
}
```

### Cupertino TextField

```go
textField := cupertino.NewCupertinoTextField()
textField.Placeholder = "Enter your email"
textField.OnChanged = func(text string) {
    fmt.Println("Text changed:", text)
}
```

### Properties

**Material TextField:**
- `Controller` - Text editing controller with signal
- `Decoration` - Input decoration (label, hint, icons)
- `KeyboardType` - Type of keyboard to show
- `MaxLines` - Number of lines (1 for single-line)
- `ObscureText` - Hide text (for passwords)
- `Enabled` - Whether field is enabled
- `OnChanged` - Called when text changes
- `OnSubmitted` - Called when user submits

**Cupertino TextField:**
- `Controller` - Text editing controller
- `Placeholder` - Placeholder text
- `Prefix` - Leading widget
- `Suffix` - Trailing widget
- `ObscureText` - Hide text (for passwords)
- `OnChanged` - Called when text changes

### Text Editing Controller

Manage text field value with signals:

```go
controller := material.NewTextEditingController("Initial text")

// Read current value
currentText := controller.Text.Get()

// Update value programmatically
controller.Text.Set("New text")

// React to changes
signals.NewEffect(func() {
    fmt.Println("Text is:", controller.Text.Get())
})
```

### Input Decoration

Customize Material TextField appearance:

```go
&material.InputDecoration{
    Label Text: "Username",
    HintText: "johndoe",
    HelperText: "Your unique username",
    ErrorText: "", // Set to show error
    PrefixIcon: widgets.NewIcon(widgets.IconPerson),
    SuffixIcon: widgets.NewIcon(widgets.IconCheck),
    Border: material.InputBorderOutline,
    FillColor: goflow.NewColor(245, 245, 245, 255),
    Filled: true,
}
```

### Example: Password Field

```go
passwordField := material.NewTextField()
passwordField.ObscureText = true
passwordField.Decoration = &material.InputDecoration{
    LabelText: "Password",
    PrefixIcon: widgets.NewIcon(widgets.IconLock),
    Border: material.InputBorderOutline,
}
```

### Example: Reactive Form Validation

```go
email := signals.New("")
emailValid := signals.NewComputed(func() bool {
    e := email.Get()
    return strings.Contains(e, "@")
})

emailField := material.NewTextField()
emailField.Controller.Text = email
emailField.Decoration = &material.InputDecoration{
    LabelText: "Email",
    ErrorText: func() string {
        if !emailValid.Get() && email.Get() != "" {
            return "Invalid email"
        }
        return ""
    }(),
}
```

## Checkbox

Material Design checkbox.

```go
checked := signals.New(false)

checkbox := material.NewCheckbox(checked.Get(), func(value bool) {
    checked.Set(value)
    fmt.Println("Checkbox:", value)
})
```

### Properties

- `Value` - Current state (true/false)
- `OnChanged` - Callback when toggled
- `ActiveColor` - Color when checked
- `CheckColor` - Color of checkmark
- `Tristate` - Allow null/indeterminate state

### Example: Checkbox with Label

```go
&widgets.Row{
    Children: []goflow.Widget{
        material.NewCheckbox(isChecked, onChanged),
        &widgets.Text{Data: "I agree to the terms"},
    },
}
```

### Example: Reactive Checkbox

```go
termsAccepted := signals.New(false)

&widgets.Row{
    Children: []goflow.Widget{
        material.NewCheckbox(termsAccepted.Get(), func(value bool) {
            termsAccepted.Set(value)
        }),
        &widgets.Text{Data: "I accept the terms and conditions"},
    },
}

// Somewhere else in UI
submitButton := material.NewButton(
    &widgets.Text{Data: "Submit"},
    func() {
        if !termsAccepted.Get() {
            showError("Please accept terms")
            return
        }
        submit()
    },
)
```

## Radio

Material Design radio button.

```go
selectedOption := signals.New("option1")

&widgets.Column{
    Children: []goflow.Widget{
        &widgets.Row{
            Children: []goflow.Widget{
                material.NewRadio("option1", selectedOption.Get(), func(value interface{}) {
                    selectedOption.Set(value.(string))
                }),
                &widgets.Text{Data: "Option 1"},
            },
        },
        &widgets.Row{
            Children: []goflow.Widget{
                material.NewRadio("option2", selectedOption.Get(), func(value interface{}) {
                    selectedOption.Set(value.(string))
                }),
                &widgets.Text{Data: "Option 2"},
            },
        },
    },
}
```

### Properties

- `Value` - Value of this radio button
- `GroupValue` - Currently selected value in group
- `OnChanged` - Callback when selected
- `ActiveColor` - Color when selected

### Example: Radio Group

```go
type RadioOption struct {
    Value string
    Label string
}

func buildRadioGroup(options []RadioOption, selectedValue *signals.Signal[string]) goflow.Widget {
    children := make([]goflow.Widget, len(options))

    for i, opt := range options {
        option := opt // Capture for closure
        children[i] = &widgets.Row{
            Children: []goflow.Widget{
                material.NewRadio(
                    option.Value,
                    selectedValue.Get(),
                    func(value interface{}) {
                        selectedValue.Set(value.(string))
                    },
                ),
                &widgets.Text{Data: option.Label},
            },
        }
    }

    return &widgets.Column{Children: children}
}

// Usage
gender := signals.New("male")
buildRadioGroup([]RadioOption{
    {Value: "male", Label: "Male"},
    {Value: "female", Label: "Female"},
    {Value: "other", Label: "Other"},
}, gender)
```

## Switch

Toggle switch with Material or Cupertino styling.

### Material Switch

```go
enabled := signals.New(true)

switchWidget := material.NewSwitch(enabled.Get(), func(value bool) {
    enabled.Set(value)
})
```

### Cupertino Switch

```go
enabled := signals.New(true)

switchWidget := cupertino.NewCupertinoSwitch(enabled.Get(), func(value bool) {
    enabled.Set(value)
})
```

### Properties

**Material Switch:**
- `Value` - Current state
- `OnChanged` - Callback when toggled
- `ActiveColor` - Color when on
- `ActiveTrackColor` - Track color when on
- `InactiveThumbColor` - Thumb color when off
- `InactiveTrackColor` - Track color when off

**Cupertino Switch:**
- `Value` - Current state
- `OnChanged` - Callback when toggled
- `ActiveColor` - Color when on (defaults to iOS blue)

### Example: Settings Toggle

```go
notifications := signals.New(true)

material.NewListTile(&widgets.Row{
    Children: []goflow.Widget{
        &widgets.Expanded{
            Child: &widgets.Column{
                Children: []goflow.Widget{
                    &widgets.Text{Data: "Notifications"},
                    &widgets.Text{Data: "Receive push notifications"},
                },
            },
        },
        material.NewSwitch(notifications.Get(), func(value bool) {
            notifications.Set(value)
        }),
    },
})
```

## Slider

Value slider with Material or Cupertino styling.

### Material Slider

```go
volume := signals.New(50.0)

slider := material.NewSlider(volume.Get(), 0, 100, func(value float64) {
    volume.Set(value)
})
```

### Cupertino Slider

```go
brightness := signals.New(0.5)

slider := cupertino.NewCupertinoSlider(brightness.Get(), 0, 1, func(value float64) {
    brightness.Set(value)
})
```

### Properties

**Material Slider:**
- `Value` - Current value
- `Min` - Minimum value
- `Max` - Maximum value
- `Divisions` - Number of discrete divisions (optional)
- `OnChanged` - Called when value changes
- `ActiveColor` - Color of active track
- `InactiveColor` - Color of inactive track
- `Label` - Label shown when dragging

**Cupertino Slider:**
- `Value` - Current value
- `Min` - Minimum value
- `Max` - Maximum value
- `Divisions` - Number of discrete divisions (optional)
- `OnChanged` - Called when value changes
- `ActiveColor` - Color of active track

### Example: Volume Control

```go
volume := signals.New(50.0)

&widgets.Column{
    Children: []goflow.Widget{
        &widgets.Row{
            Children: []goflow.Widget{
                widgets.NewIcon(widgets.IconVolumeDown),
                &widgets.Expanded{
                    Child: material.NewSlider(volume.Get(), 0, 100, func(v float64) {
                        volume.Set(v)
                        setVolume(v)
                    }),
                },
                widgets.NewIcon(widgets.IconVolumeUp),
            },
        },
        &widgets.Text{
            Data: fmt.Sprintf("Volume: %.0f%%", volume.Get()),
        },
    },
}
```

### Example: Discrete Slider

```go
rating := signals.New(3.0)
divisions := 5

slider := material.NewSlider(rating.Get(), 0, 5, func(v float64) {
    rating.Set(v)
})
slider.Divisions = &divisions
slider.Label = fmt.Sprintf("%.0f stars", rating.Get())
```

## Form Best Practices

### 1. Use Signals for Form State

```go
type FormData struct {
    Name     *signals.Signal[string]
    Email    *signals.Signal[string]
    Age      *signals.Signal[int]
    Agree    *signals.Signal[bool]
}

func NewFormData() *FormData {
    return &FormData{
        Name:  signals.New(""),
        Email: signals.New(""),
        Age:   signals.New(0),
        Agree: signals.New(false),
    }
}
```

### 2. Validate Reactively

```go
formValid := signals.NewComputed(func() bool {
    return form.Name.Get() != "" &&
           form.Email.Get() != "" &&
           form.Agree.Get()
})

submitButton.Enabled = formValid.Get()
```

### 3. Show Errors Conditionally

```go
emailError := signals.NewComputed(func() string {
    email := form.Email.Get()
    if email == "" {
        return ""
    }
    if !isValidEmail(email) {
        return "Invalid email address"
    }
    return ""
})

textField.Decoration.ErrorText = emailError.Get()
```

### 4. Provide Visual Feedback

```go
&widgets.Column{
    Children: []goflow.Widget{
        buildTextField(form.Name, "Name"),
        buildTextField(form.Email, "Email"),
        material.NewButton(
            &widgets.Text{Data: "Submit"},
            func() {
                if !formValid.Get() {
                    showSnackbar("Please fill all required fields")
                    return
                }
                submitForm()
            },
        ),
    },
}
```

### 5. Handle Loading States

```go
isLoading := signals.New(false)

submitButton := material.NewButton(
    func() goflow.Widget {
        if isLoading.Get() {
            return &widgets.Text{Data: "Submitting..."}
        }
        return &widgets.Text{Data: "Submit"}
    }(),
    func() {
        if isLoading.Get() {
            return
        }

        isLoading.Set(true)
        go func() {
            err := submitForm()
            isLoading.Set(false)
            if err != nil {
                showError(err)
            }
        }()
    },
)
submitButton.Enabled = !isLoading.Get()
```
