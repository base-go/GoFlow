# Material Design Widgets

Material-specific widgets following Google's Material Design guidelines.

## Card

Material Design card - a sheet of material displaying related content.

```go
material.NewCard(&widgets.Text{Data: "Card content"})
```

### Properties

- `Child` - Card content
- `Elevation` - Shadow elevation (default: 1.0)
- `Color` - Background color (default: surface color)
- `ShadowColor` - Shadow color
- `BorderRadius` - Corner radius (default: 4.0)
- `Margin` - External spacing (default: 8px all sides)
- `Padding` - Internal spacing (default: 16px all sides)
- `Width` - Card width
- `Height` - Card height

### Basic Card

```go
card := material.NewCard(&widgets.Text{
    Data: "Hello, Material!",
})
```

### Styled Card

```go
card := material.NewCard(content)
card.Elevation = 4.0
card.BorderRadius = 8.0
card.Color = goflow.NewColor(255, 255, 255, 255)
card.Margin = goflow.NewEdgeInsets(16, 8, 16, 8)
```

### Example: Profile Card

```go
func buildProfileCard(name, email, avatarURL string) goflow.Widget {
    avatar := widgets.NewImage(&widgets.NetworkImage{
        URL: avatarURL,
    })
    avatar.Fit = widgets.ImageFitCover
    size := 64.0
    avatar.Width = &size
    avatar.Height = &size

    return material.NewCard(&widgets.Row{
        Children: []goflow.Widget{
            &widgets.Container{
                Width: &size,
                Height: &size,
                Decoration: &goflow.BoxDecoration{
                    BorderRadius: size / 2,
                },
                Child: avatar,
            },
            &widgets.Expanded{
                Child: &widgets.Column{
                    Children: []goflow.Widget{
                        &widgets.Text{Data: name},
                        &widgets.Text{Data: email},
                    },
                    CrossAxisAlign: widgets.CrossAxisStart,
                },
            },
        },
    })
}
```

### Example: Stats Card

```go
func buildStatsCard(title, value, change string) goflow.Widget {
    card := material.NewCard(&widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{Data: title},
            &widgets.Text{Data: value},
            &widgets.Text{Data: change},
        },
        CrossAxisAlign: widgets.CrossAxisStart,
    })

    width := 200.0
    card.Width = &width

    return card
}
```

## ListTile

Fixed-height row typically used in lists.

```go
material.NewListTile(&widgets.Text{Data: "Title"})
```

### Properties

- `Leading` - Leading widget (icon or avatar)
- `Title` - Title widget (required)
- `Subtitle` - Subtitle widget
- `Trailing` - Trailing widget (icon or action)
- `IsThreeLine` - Three-line layout (height: 88)
- `Dense` - Compact layout (height: 48)
- `Enabled` - Whether enabled (default: true)
- `Selected` - Whether selected (default: false)
- `OnTap` - Tap callback
- `OnLongPress` - Long press callback
- `ContentPadding` - Internal padding (default: 16h, 8v)

### Basic ListTile

```go
tile := material.NewListTile(&widgets.Text{
    Data: "List item",
})
```

### ListTile with Icon

```go
tile := material.NewListTile(&widgets.Text{Data: "Home"})
tile.Leading = widgets.NewIcon(widgets.IconHome)
tile.OnTap = func() {
    navigateHome()
}
```

### ListTile with All Elements

```go
tile := material.NewListTile(&widgets.Text{Data: "Message"})
tile.Leading = widgets.NewIcon(widgets.IconEmail)
tile.Subtitle = &widgets.Text{Data: "This is a message preview"}
tile.Trailing = &widgets.Text{Data: "2:30 PM"}
tile.OnTap = openMessage
```

### Example: Contact List

```go
type Contact struct {
    Name  string
    Email string
    Phone string
}

func buildContactTile(contact Contact) goflow.Widget {
    tile := material.NewListTile(&widgets.Text{
        Data: contact.Name,
    })

    tile.Leading = &widgets.Container{
        Width: floatPtr(40.0),
        Height: floatPtr(40.0),
        Child: &widgets.Center{
            Child: widgets.NewIcon(widgets.IconPerson),
        },
    }

    tile.Subtitle = &widgets.Text{
        Data: contact.Email,
    }

    tile.Trailing = widgets.NewIconButton(
        widgets.IconPhone,
        func() {
            callContact(contact)
        },
    )

    tile.OnTap = func() {
        showContactDetails(contact)
    }

    return tile
}
```

### Example: Settings List

```go
func buildSettingsList() goflow.Widget {
    settings := []struct {
        Icon  widgets.IconData
        Title string
        Value string
    }{
        {widgets.IconPerson, "Account", "John Doe"},
        {widgets.IconSettings, "Preferences", "Configured"},
        {widgets.IconWarning, "Privacy", "Enhanced"},
    }

    items := make([]goflow.Widget, len(settings))
    for i, setting := range settings {
        tile := material.NewListTile(&widgets.Text{
            Data: setting.Title,
        })
        tile.Leading = widgets.NewIcon(setting.Icon)
        tile.Subtitle = &widgets.Text{Data: setting.Value}
        tile.Trailing = widgets.NewIcon(widgets.IconArrowForward)

        items[i] = tile
    }

    return &widgets.ListView{
        Children: items,
    }
}
```

## Dialog

Material Design dialog for important information.

```go
material.NewDialog(&widgets.Text{Data: "Dialog content"})
```

### Properties

- `Child` - Dialog content
- `BackgroundColor` - Background color
- `Elevation` - Shadow elevation (default: 24.0)
- `BorderRadius` - Corner radius (default: 4.0)
- `InsetPadding` - Padding from screen edges

### Basic Dialog

```go
dialog := material.NewDialog(&widgets.Column{
    Children: []goflow.Widget{
        &widgets.Text{Data: "Important message"},
        material.NewButton(
            &widgets.Text{Data: "OK"},
            closeDialog,
        ),
    },
})
```

### Example: Confirmation Dialog

```go
func buildConfirmDialog(message string, onConfirm func()) goflow.Widget {
    return material.NewDialog(&widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{Data: message},
            &widgets.Row{
                Children: []goflow.Widget{
                    material.NewTextButton("Cancel", closeDialog),
                    material.NewButton(
                        &widgets.Text{Data: "Confirm"},
                        onConfirm,
                    ),
                },
                MainAxisAlignment: widgets.MainAxisEnd,
            },
        },
    })
}
```

## AlertDialog

Material Design alert dialog with title, content, and actions.

```go
dialog := material.NewAlertDialog()
dialog.Title = &widgets.Text{Data: "Alert"}
dialog.Content = &widgets.Text{Data: "This is an alert message"}
dialog.Actions = []goflow.Widget{
    material.NewTextButton("Cancel", onCancel),
    material.NewTextButton("OK", onOK),
}
```

### Properties

- `Title` - Title widget
- `Content` - Content widget
- `Actions` - Action buttons
- `BackgroundColor` - Background color

### Basic Alert

```go
alert := material.NewAlertDialog()
alert.Title = &widgets.Text{Data: "Delete Item?"}
alert.Content = &widgets.Text{Data: "This action cannot be undone."}
alert.Actions = []goflow.Widget{
    material.NewTextButton("Cancel", closeDialog),
    material.NewTextButton("Delete", deleteItem),
}
```

### Example: Error Dialog

```go
func buildErrorDialog(errorMsg string) goflow.Widget {
    alert := material.NewAlertDialog()

    titleStyle := goflow.NewTextStyle()
    titleStyle.FontSize = 20.0
    titleStyle.FontWeight = goflow.FontWeightBold
    titleStyle.Color = goflow.NewColor(244, 67, 54, 255) // Red

    alert.Title = &widgets.Row{
        Children: []goflow.Widget{
            widgets.NewIcon(widgets.IconError),
            &widgets.Text{
                Data: "Error",
                Style: titleStyle,
            },
        },
    }

    alert.Content = &widgets.Text{Data: errorMsg}

    alert.Actions = []goflow.Widget{
        material.NewTextButton("OK", closeDialog),
    }

    return alert
}
```

### Example: Form Dialog

```go
func buildInputDialog() goflow.Widget {
    nameController := material.NewTextEditingController()

    alert := material.NewAlertDialog()
    alert.Title = &widgets.Text{Data: "Enter Your Name"}

    alert.Content = &widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{Data: "Please provide your name:"},
            material.NewTextField(nameController),
        },
        CrossAxisAlign: widgets.CrossAxisStretch,
    }

    alert.Actions = []goflow.Widget{
        material.NewTextButton("Cancel", closeDialog),
        material.NewButton(
            &widgets.Text{Data: "Submit"},
            func() {
                submitName(nameController.Text.Get())
                closeDialog()
            },
        ),
    }

    return alert
}
```

## DrawerHeader

Header section for navigation drawer.

```go
material.NewDrawerHeader(&widgets.Column{
    Children: []goflow.Widget{
        &widgets.Text{Data: "John Doe"},
        &widgets.Text{Data: "john@example.com"},
    },
})
```

### Properties

- `Child` - Header content
- `Decoration` - Background decoration
- `Padding` - Internal padding (default: 16px)
- `Margin` - External margin

### Basic Header

```go
header := material.NewDrawerHeader(&widgets.Text{
    Data: "My App",
})
```

### Styled Header

```go
header := material.NewDrawerHeader(content)
header.Decoration = &material.BoxDecoration{
    Color: goflow.NewColor(33, 150, 243, 255),
}
header.Padding = goflow.NewEdgeInsets(24, 16, 24, 16)
```

### Example: User Profile Header

```go
func buildDrawerHeader(name, email, avatarURL string) goflow.Widget {
    size := 64.0
    avatar := widgets.NewImage(&widgets.NetworkImage{
        URL: avatarURL,
    })
    avatar.Width = &size
    avatar.Height = &size
    avatar.Fit = widgets.ImageFitCover

    nameStyle := goflow.NewTextStyle()
    nameStyle.FontSize = 18.0
    nameStyle.FontWeight = goflow.FontWeightBold
    nameStyle.Color = goflow.NewColor(255, 255, 255, 255)

    emailStyle := goflow.NewTextStyle()
    emailStyle.FontSize = 14.0
    emailStyle.Color = goflow.NewColor(255, 255, 255, 200)

    header := material.NewDrawerHeader(&widgets.Column{
        Children: []goflow.Widget{
            &widgets.Container{
                Width: &size,
                Height: &size,
                Decoration: &goflow.BoxDecoration{
                    BorderRadius: size / 2,
                },
                Child: avatar,
            },
            &widgets.Container{Height: floatPtr(16.0)},
            &widgets.Text{Data: name, Style: nameStyle},
            &widgets.Text{Data: email, Style: emailStyle},
        },
        CrossAxisAlign: widgets.CrossAxisStart,
    })

    header.Decoration = &material.BoxDecoration{
        Color: goflow.NewColor(33, 150, 243, 255),
    }

    return header
}
```

### Example: App Branding Header

```go
func buildBrandingHeader() goflow.Widget {
    logo := widgets.NewImage(&widgets.AssetImage{
        Path: "assets/logo.png",
    })
    size := 80.0
    logo.Width = &size
    logo.Height = &size

    titleStyle := goflow.NewTextStyle()
    titleStyle.FontSize = 24.0
    titleStyle.FontWeight = goflow.FontWeightBold
    titleStyle.Color = goflow.NewColor(255, 255, 255, 255)

    header := material.NewDrawerHeader(&widgets.Column{
        Children: []goflow.Widget{
            logo,
            &widgets.Text{
                Data: "MyApp",
                Style: titleStyle,
            },
            &widgets.Text{
                Data: "Version 1.0.0",
            },
        },
        MainAxisAlign: widgets.MainAxisCenter,
        CrossAxisAlign: widgets.CrossAxisCenter,
    })

    header.Decoration = &material.BoxDecoration{
        Gradient: &material.Gradient{
            Colors: []goflow.Color{
                *goflow.NewColor(33, 150, 243, 255),
                *goflow.NewColor(21, 101, 192, 255),
            },
        },
    }

    return header
}
```

## Best Practices

### 1. Use Cards for Grouped Content

```go
// ✅ Good - card groups related information
material.NewCard(&widgets.Column{
    Children: []goflow.Widget{
        title,
        content,
        actions,
    },
})
```

### 2. Keep ListTile Content Concise

```go
// ✅ Good - clear and concise
tile := material.NewListTile(&widgets.Text{Data: "Settings"})
tile.Leading = widgets.NewIcon(widgets.IconSettings)

// ❌ Avoid - too much content
tile.Subtitle = &widgets.Text{
    Data: "Very long subtitle that wraps multiple lines...",
}
```

### 3. Limit Dialog Actions

```go
// ✅ Good - 1-3 actions
alert.Actions = []goflow.Widget{
    material.NewTextButton("Cancel", onCancel),
    material.NewTextButton("OK", onOK),
}

// ❌ Avoid - too many actions
alert.Actions = []goflow.Widget{
    // 5+ buttons - overwhelming
}
```

### 4. Use Appropriate Card Elevation

```go
// Resting elevation: 1-2
card.Elevation = 1.0

// Raised/selected: 4-8
card.Elevation = 6.0

// Dialog/modal: 16-24
dialog.Elevation = 24.0
```

### 5. Provide Visual Hierarchy in Dialogs

```go
alert := material.NewAlertDialog()

// Title should be prominent
titleStyle := goflow.NewTextStyle()
titleStyle.FontSize = 20.0
titleStyle.FontWeight = goflow.FontWeightBold
alert.Title = &widgets.Text{Data: "Important", Style: titleStyle}

// Content is smaller
alert.Content = &widgets.Text{Data: "Message"}

// Actions aligned right
alert.Actions = []goflow.Widget{
    material.NewTextButton("Cancel", onCancel),
    material.NewButton(&widgets.Text{Data: "Confirm"}, onConfirm),
}
```

### 6. Use Dense ListTiles for Compact Lists

```go
// Normal - 56px height
tile := material.NewListTile(title)

// Dense - 48px height (for lists with many items)
tile.Dense = true
```

### 7. Indicate Selected State

```go
currentRoute := signals.New("home")

func buildNavTile(route string, title string) goflow.Widget {
    tile := material.NewListTile(&widgets.Text{Data: title})
    tile.Selected = currentRoute.Get() == route
    tile.OnTap = func() {
        currentRoute.Set(route)
        navigateTo(route)
    }
    return tile
}
```

### 8. Add Padding to Card Content

```go
// ✅ Good - content has breathing room
card := material.NewCard(content)
card.Padding = goflow.NewEdgeInsets(16, 16, 16, 16)

// ❌ Avoid - content touches edges
card.Padding = goflow.ZeroEdgeInsets()
```

## Common Patterns

### Card List

```go
func buildCardList(items []Item) goflow.Widget {
    cards := make([]goflow.Widget, len(items))

    for i, item := range items {
        cards[i] = material.NewCard(&widgets.Column{
            Children: []goflow.Widget{
                &widgets.Text{Data: item.Title},
                &widgets.Text{Data: item.Description},
            },
            CrossAxisAlign: widgets.CrossAxisStart,
        })
    }

    return &widgets.ListView{
        Children: cards,
    }
}
```

### Drawer with Header and Items

```go
func buildDrawer() goflow.Widget {
    return material.NewDrawer(&widgets.Column{
        Children: []goflow.Widget{
            buildDrawerHeader("User", "user@example.com", "avatar.jpg"),
            material.NewListTile(&widgets.Text{Data: "Home"}),
            material.NewListTile(&widgets.Text{Data: "Settings"}),
            &widgets.Container{
                Height: floatPtr(1.0),
                Color: goflow.NewColor(224, 224, 224, 255),
            },
            material.NewListTile(&widgets.Text{Data: "Logout"}),
        },
    })
}
```

### Confirmation Dialog Pattern

```go
func showConfirmDialog(title, message string, onConfirm func()) {
    alert := material.NewAlertDialog()
    alert.Title = &widgets.Text{Data: title}
    alert.Content = &widgets.Text{Data: message}
    alert.Actions = []goflow.Widget{
        material.NewTextButton("Cancel", closeDialog),
        material.NewButton(
            &widgets.Text{Data: "Confirm"},
            func() {
                onConfirm()
                closeDialog()
            },
        ),
    }

    showDialog(alert)
}
```

## Helper Functions

```go
// Float pointer helper
func floatPtr(f float64) *float64 {
    return &f
}

// Build separator
func buildDivider() goflow.Widget {
    return &widgets.Container{
        Height: floatPtr(1.0),
        Color: goflow.NewColor(224, 224, 224, 255),
    }
}

// Build section header
func buildSectionHeader(title string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 14.0
    style.FontWeight = goflow.FontWeightBold
    style.Color = goflow.NewColor(100, 100, 100, 255)

    return &widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
        Child: &widgets.Text{
            Data: title,
            Style: style,
        },
    }
}
```
