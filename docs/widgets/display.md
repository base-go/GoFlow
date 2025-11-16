# Display Widgets

Widgets for displaying text, icons, and images.

## Text

Display styled text content.

```go
&widgets.Text{
    Data: "Hello, World!",
}
```

### Properties

- `Data` - Text content to display
- `Style` - Text style (font, size, color, weight, etc.)

### Basic Text

```go
&widgets.Text{
    Data: "Simple text",
}
```

### Styled Text

```go
style := goflow.NewTextStyle()
style.FontSize = 24.0
style.Color = goflow.NewColor(33, 150, 243, 255) // Blue
style.FontWeight = goflow.FontWeightBold

&widgets.Text{
    Data: "Styled text",
    Style: style,
}
```

### Text Styles

```go
// Create text style
style := goflow.NewTextStyle()

// Font properties
style.FontSize = 16.0           // Font size in pixels
style.FontWeight = goflow.FontWeightNormal  // or FontWeightBold
style.FontStyle = goflow.FontStyleNormal    // or FontStyleItalic

// Color
style.Color = goflow.NewColor(0, 0, 0, 255)

// Alignment
style.TextAlign = goflow.TextAlignCenter  // Left, Center, Right, Justify

// Decoration
style.Decoration = goflow.TextDecorationUnderline  // Underline, Overline, LineThrough
```

### Example: Headings

```go
func buildHeading(text string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 32.0
    style.FontWeight = goflow.FontWeightBold
    style.Color = goflow.NewColor(0, 0, 0, 255)

    return &widgets.Text{
        Data: text,
        Style: style,
    }
}

func buildSubheading(text string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 20.0
    style.FontWeight = goflow.FontWeightBold
    style.Color = goflow.NewColor(100, 100, 100, 255)

    return &widgets.Text{
        Data: text,
        Style: style,
    }
}

func buildBody(text string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 16.0
    style.Color = goflow.NewColor(0, 0, 0, 255)

    return &widgets.Text{
        Data: text,
        Style: style,
    }
}
```

### Example: Colored Text

```go
errorStyle := goflow.NewTextStyle()
errorStyle.Color = goflow.NewColor(244, 67, 54, 255) // Red

&widgets.Text{
    Data: "Error message",
    Style: errorStyle,
}

successStyle := goflow.NewTextStyle()
successStyle.Color = goflow.NewColor(76, 175, 80, 255) // Green

&widgets.Text{
    Data: "Success!",
    Style: successStyle,
}
```

### Example: Dynamic Text with Signals

```go
counter := signals.New(0)

&widgets.Text{
    Data: fmt.Sprintf("Count: %d", counter.Get()),
}

// Update text reactively
button := material.NewButton(
    &widgets.Text{Data: "Increment"},
    func() {
        counter.Set(counter.Get() + 1)
    },
)
```

## Icon

Display icons from the Material Icons set.

```go
&widgets.Icon{
    Icon: widgets.IconHome,
    Size: 24.0,
    Color: goflow.NewColor(33, 150, 243, 255),
}
```

### Properties

- `Icon` - Icon data (from predefined icons)
- `Size` - Icon size in pixels (default 24.0)
- `Color` - Icon color
- `SemanticLabel` - Accessibility label

### Available Material Icons

```go
widgets.IconHome
widgets.IconMenu
widgets.IconSearch
widgets.IconSettings
widgets.IconFavorite
widgets.IconAdd
widgets.IconRemove
widgets.IconClose
widgets.IconCheck
widgets.IconArrowBack
widgets.IconArrowForward
widgets.IconEdit
widgets.IconDelete
widgets.IconShare
widgets.IconPerson
widgets.IconEmail
widgets.IconPhone
widgets.IconInfo
widgets.IconWarning
widgets.IconError
```

### Basic Icon

```go
icon := widgets.NewIcon(widgets.IconHome)
```

### Colored Icon

```go
icon := widgets.NewIcon(widgets.IconFavorite)
icon.Size = 32.0
icon.Color = goflow.NewColor(244, 67, 54, 255) // Red
```

### Example: Icon with Text

```go
&widgets.Row{
    Children: []goflow.Widget{
        widgets.NewIcon(widgets.IconPerson),
        &widgets.Text{Data: "Profile"},
    },
    MainAxisAlignment: widgets.MainAxisCenter,
}
```

### Example: Colored Icons

```go
func buildStatusIcon(status string) goflow.Widget {
    var icon widgets.IconData
    var color *goflow.Color

    switch status {
    case "success":
        icon = widgets.IconCheck
        color = goflow.NewColor(76, 175, 80, 255) // Green
    case "warning":
        icon = widgets.IconWarning
        color = goflow.NewColor(255, 152, 0, 255) // Orange
    case "error":
        icon = widgets.IconError
        color = goflow.NewColor(244, 67, 54, 255) // Red
    default:
        icon = widgets.IconInfo
        color = goflow.NewColor(33, 150, 243, 255) // Blue
    }

    iconWidget := widgets.NewIcon(icon)
    iconWidget.Size = 48.0
    iconWidget.Color = color

    return iconWidget
}
```

### Example: Icon in AppBar

```go
appBar := material.NewAppBar(&widgets.Text{Data: "My App"})
appBar.Leading = widgets.NewIconButton(widgets.IconMenu, onMenuPressed)
appBar.Actions = []goflow.Widget{
    widgets.NewIconButton(widgets.IconSearch, onSearch),
    widgets.NewIconButton(widgets.IconSettings, onSettings),
}
```

## IconButton

Button that displays an icon.

See [buttons.md#iconbutton](buttons.md#iconbutton) for complete documentation.

```go
button := widgets.NewIconButton(
    widgets.IconFavorite,
    func() {
        toggleFavorite()
    },
)
button.Color = goflow.NewColor(244, 67, 54, 255)
button.IconSize = 28.0
```

## Image

Display images from various sources.

```go
&widgets.Image{
    Source: &widgets.AssetImage{Path: "assets/logo.png"},
    Width: floatPtr(200.0),
    Height: floatPtr(150.0),
    Fit: widgets.ImageFitCover,
}
```

### Properties

- `Source` - Image source (Asset, Network, or File)
- `Fit` - How to inscribe the image
- `Width` - Image width (optional)
- `Height` - Image height (optional)
- `SemanticLabel` - Accessibility label

### Image Sources

**AssetImage** - Load from bundled assets:
```go
&widgets.AssetImage{
    Path: "assets/images/logo.png",
}
```

**NetworkImage** - Load from URL:
```go
&widgets.NetworkImage{
    URL: "https://example.com/image.jpg",
}
```

**FileImage** - Load from file system:
```go
&widgets.FileImage{
    Path: "/path/to/image.png",
}
```

### Image Fit Modes

- `ImageFitFill` - Fill the box, distorting aspect ratio
- `ImageFitContain` - Contain within box, preserving aspect ratio
- `ImageFitCover` - Cover the box, preserving aspect ratio (crop if needed)
- `ImageFitFitWidth` - Fit width, height may overflow
- `ImageFitFitHeight` - Fit height, width may overflow
- `ImageFitNone` - No scaling
- `ImageFitScaleDown` - Scale down if needed, never scale up

### Example: Asset Image

```go
image := widgets.NewImage(&widgets.AssetImage{
    Path: "assets/profile.jpg",
})
image.Width = floatPtr(100.0)
image.Height = floatPtr(100.0)
image.Fit = widgets.ImageFitCover
```

### Example: Network Image

```go
image := widgets.NewImage(&widgets.NetworkImage{
    URL: "https://example.com/avatar.jpg",
})
image.Fit = widgets.ImageFitContain
```

### Example: Avatar

```go
func buildAvatar(imageURL string, size float64) goflow.Widget {
    image := widgets.NewImage(&widgets.NetworkImage{
        URL: imageURL,
    })
    image.Width = &size
    image.Height = &size
    image.Fit = widgets.ImageFitCover

    return &widgets.Container{
        Width: &size,
        Height: &size,
        Decoration: &goflow.BoxDecoration{
            BorderRadius: size / 2, // Circular
        },
        Child: image,
    }
}
```

### Example: Image with Placeholder

```go
isLoading := signals.New(true)

func buildImageWithLoading(url string) goflow.Widget {
    if isLoading.Get() {
        return &widgets.Center{
            Child: &widgets.Container{
                Width: floatPtr(50.0),
                Height: floatPtr(50.0),
                Child: &widgets.Text{Data: "Loading..."},
            },
        }
    }

    return widgets.NewImage(&widgets.NetworkImage{
        URL: url,
    })
}
```

### Example: Image Gallery

```go
func buildGallery(imageURLs []string) goflow.Widget {
    images := make([]goflow.Widget, len(imageURLs))

    for i, url := range imageURLs {
        image := widgets.NewImage(&widgets.NetworkImage{
            URL: url,
        })
        image.Fit = widgets.ImageFitCover

        images[i] = &widgets.Container{
            Width: floatPtr(150.0),
            Height: floatPtr(150.0),
            Margin: goflow.NewEdgeInsets(8, 8, 8, 8),
            Child: image,
        }
    }

    return &widgets.GridView{
        Children: images,
        CrossAxisCount: 3,
        ChildAspectRatio: 1.0,
        MainAxisSpacing: 8.0,
        CrossAxisSpacing: 8.0,
    }
}
```

## Best Practices

### 1. Use Semantic Text Styles

```go
// ✅ Good - semantic style names
func buildTitle(text string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 24.0
    style.FontWeight = goflow.FontWeightBold
    return &widgets.Text{Data: text, Style: style}
}

// ❌ Avoid - hardcoded styles everywhere
&widgets.Text{
    Data: "Title",
    Style: &goflow.TextStyle{
        FontSize: 24.0,
        FontWeight: goflow.FontWeightBold,
    },
}
```

### 2. Provide Semantic Labels for Accessibility

```go
// ✅ Good - accessible
icon := widgets.NewIcon(widgets.IconHome)
icon.SemanticLabel = "Home"

image := widgets.NewImage(source)
image.SemanticLabel = "User profile picture"
```

### 3. Use Appropriate Image Fit

```go
// Profile pictures - cover
image.Fit = widgets.ImageFitCover

// Logos - contain
image.Fit = widgets.ImageFitContain

// Banners - fit width
image.Fit = widgets.ImageFitFitWidth
```

### 4. Size Icons Consistently

```go
// Standard sizes
const (
    IconSizeSmall  = 18.0
    IconSizeMedium = 24.0  // Default
    IconSizeLarge  = 32.0
    IconSizeHuge   = 48.0
)
```

### 5. Use Color Constants

```go
var (
    ColorPrimary = goflow.NewColor(33, 150, 243, 255)
    ColorAccent = goflow.NewColor(255, 64, 129, 255)
    ColorError = goflow.NewColor(244, 67, 54, 255)
    ColorSuccess = goflow.NewColor(76, 175, 80, 255)
)

&widgets.Text{
    Data: "Success!",
    Style: &goflow.TextStyle{
        Color: ColorSuccess,
    },
}
```

### 6. Optimize Image Sizes

```go
// ✅ Good - specify dimensions
image := widgets.NewImage(source)
image.Width = floatPtr(200.0)
image.Height = floatPtr(150.0)

// ❌ Avoid - unbounded images can cause layout issues
image := widgets.NewImage(source)
```

### 7. Handle Missing Images

```go
func buildImageWithFallback(url string) goflow.Widget {
    if url == "" {
        return widgets.NewIcon(widgets.IconPerson)
    }

    return widgets.NewImage(&widgets.NetworkImage{
        URL: url,
    })
}
```

### 8. Use Icons for Actions

```go
// ✅ Good - clear icon for action
widgets.NewIconButton(widgets.IconDelete, onDelete)

// ✅ Also good - icon with label
&widgets.Row{
    Children: []goflow.Widget{
        widgets.NewIcon(widgets.IconDelete),
        &widgets.Text{Data: "Delete"},
    },
}
```

## Helper Functions

```go
// Float pointer helper
func floatPtr(f float64) *float64 {
    return &f
}

// Build text with style
func styledText(text string, fontSize float64, color *goflow.Color) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = fontSize
    style.Color = color

    return &widgets.Text{
        Data: text,
        Style: style,
    }
}

// Build circular avatar
func circularImage(source widgets.ImageSource, size float64) goflow.Widget {
    image := widgets.NewImage(source)
    image.Width = &size
    image.Height = &size
    image.Fit = widgets.ImageFitCover

    return &widgets.Container{
        Width: &size,
        Height: &size,
        Decoration: &goflow.BoxDecoration{
            BorderRadius: size / 2,
        },
        Child: image,
    }
}
```
