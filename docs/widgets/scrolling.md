# Scrolling Widgets

Widgets that provide scrollable content areas.

## ListView

Scrollable list of widgets arranged linearly.

```go
&widgets.ListView{
    Children: []goflow.Widget{
        item1,
        item2,
        item3,
    },
}
```

### Properties

- `Children` - List of child widgets
- `ScrollDirection` - Vertical or horizontal scroll (default: Vertical)
- `Padding` - Padding around list items
- `ShrinkWrap` - Whether to size to content (default: false)
- `Physics` - Scroll behavior

### Basic ListView

```go
&widgets.ListView{
    Children: []goflow.Widget{
        &widgets.Text{Data: "Item 1"},
        &widgets.Text{Data: "Item 2"},
        &widgets.Text{Data: "Item 3"},
    },
}
```

### Horizontal ListView

```go
&widgets.ListView{
    ScrollDirection: widgets.AxisHorizontal,
    Children: []goflow.Widget{
        buildCard("Card 1"),
        buildCard("Card 2"),
        buildCard("Card 3"),
    },
}
```

### ListView with Padding

```go
listView := widgets.NewListView(items)
listView.Padding = goflow.NewEdgeInsets(16, 8, 16, 8)
```

### Scroll Physics

- `ScrollPhysicsAlwaysScrollable` - Always allow scrolling (default)
- `ScrollPhysicsNeverScrollable` - Disable scrolling
- `ScrollPhysicsBouncing` - iOS-style bouncing
- `ScrollPhysicsClamping` - Android-style clamping

```go
listView := widgets.NewListView(items)
listView.Physics = widgets.ScrollPhysicsBouncing
```

### Example: Simple List

```go
func buildSimpleList() goflow.Widget {
    items := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
    children := make([]goflow.Widget, len(items))

    for i, item := range items {
        children[i] = material.NewListTile(&widgets.Text{
            Data: item,
        })
    }

    return widgets.NewListView(children)
}
```

### Example: Heterogeneous List

```go
func buildMixedList() goflow.Widget {
    return &widgets.ListView{
        Children: []goflow.Widget{
            // Header
            &widgets.Container{
                Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
                Child: &widgets.Text{Data: "Categories"},
            },

            // Items
            material.NewListTile(&widgets.Text{Data: "Technology"}),
            material.NewListTile(&widgets.Text{Data: "Science"}),
            material.NewListTile(&widgets.Text{Data: "Art"}),

            // Divider
            &widgets.Container{
                Height: floatPtr(1.0),
                Color: goflow.NewColor(224, 224, 224, 255),
            },

            // More items
            material.NewListTile(&widgets.Text{Data: "Music"}),
            material.NewListTile(&widgets.Text{Data: "Sports"}),
        },
    }
}
```

## ListView.builder

Lazily builds list items for efficient rendering of large lists.

```go
&widgets.ListViewBuilder{
    ItemCount: 1000,
    ItemBuilder: func(ctx goflow.BuildContext, index int) goflow.Widget {
        return &widgets.Text{
            Data: fmt.Sprintf("Item %d", index),
        }
    },
}
```

### Properties

- `ItemCount` - Number of items
- `ItemBuilder` - Function to build each item
- `ScrollDirection` - Scroll direction (default: Vertical)
- `Padding` - Padding around items

### Basic Builder

```go
builder := widgets.NewListViewBuilder(
    100,
    func(ctx goflow.BuildContext, i int) goflow.Widget {
        return &widgets.Text{
            Data: fmt.Sprintf("Item %d", i),
        }
    },
)
```

### Example: Large List

```go
func buildLargeList(count int) goflow.Widget {
    return widgets.NewListViewBuilder(
        count,
        func(ctx goflow.BuildContext, index int) goflow.Widget {
            return material.NewListTile(&widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconPerson),
                    &widgets.Expanded{
                        Child: &widgets.Column{
                            Children: []goflow.Widget{
                                &widgets.Text{
                                    Data: fmt.Sprintf("User %d", index),
                                },
                                &widgets.Text{
                                    Data: fmt.Sprintf("user%d@example.com", index),
                                },
                            },
                            CrossAxisAlign: widgets.CrossAxisStart,
                        },
                    },
                },
            })
        },
    )
}
```

### Example: Dynamic List with Data

```go
type User struct {
    Name  string
    Email string
}

func buildUserList(users []User) goflow.Widget {
    return widgets.NewListViewBuilder(
        len(users),
        func(ctx goflow.BuildContext, index int) goflow.Widget {
            user := users[index]

            return material.NewListTile(&widgets.Column{
                Children: []goflow.Widget{
                    &widgets.Text{Data: user.Name},
                    &widgets.Text{Data: user.Email},
                },
                CrossAxisAlign: widgets.CrossAxisStart,
            })
        },
    )
}
```

### Example: Horizontal Carousel

```go
func buildCarousel(images []string) goflow.Widget {
    builder := widgets.NewListViewBuilder(
        len(images),
        func(ctx goflow.BuildContext, index int) goflow.Widget {
            image := widgets.NewImage(&widgets.NetworkImage{
                URL: images[index],
            })
            width := 300.0
            height := 200.0
            image.Width = &width
            image.Height = &height
            image.Fit = widgets.ImageFitCover

            return &widgets.Container{
                Width: &width,
                Height: &height,
                Margin: goflow.NewEdgeInsets(0, 8, 0, 8),
                Child: image,
            }
        },
    )
    builder.ScrollDirection = widgets.AxisHorizontal

    return builder
}
```

## GridView

Scrollable 2D grid of widgets.

```go
&widgets.GridView{
    Children: items,
    CrossAxisCount: 2,
    ChildAspectRatio: 1.0,
    MainAxisSpacing: 8.0,
    CrossAxisSpacing: 8.0,
}
```

### Properties

- `Children` - Grid items
- `CrossAxisCount` - Number of columns
- `ChildAspectRatio` - Width/height ratio of items (default: 1.0)
- `MainAxisSpacing` - Vertical spacing between rows
- `CrossAxisSpacing` - Horizontal spacing between columns
- `Padding` - Padding around grid

### Basic Grid

```go
grid := widgets.NewGridView(items, 3)
```

### Grid with Spacing

```go
grid := widgets.NewGridView(items, 2)
grid.MainAxisSpacing = 16.0
grid.CrossAxisSpacing = 16.0
grid.ChildAspectRatio = 0.75 // Portrait items
```

### Example: Image Grid

```go
func buildImageGrid(imageURLs []string) goflow.Widget {
    images := make([]goflow.Widget, len(imageURLs))

    for i, url := range imageURLs {
        image := widgets.NewImage(&widgets.NetworkImage{
            URL: url,
        })
        image.Fit = widgets.ImageFitCover

        images[i] = &widgets.Container{
            Child: image,
            Decoration: &goflow.BoxDecoration{
                BorderRadius: 8.0,
            },
        }
    }

    grid := widgets.NewGridView(images, 3)
    grid.MainAxisSpacing = 8.0
    grid.CrossAxisSpacing = 8.0
    grid.Padding = goflow.NewEdgeInsets(16, 16, 16, 16)

    return grid
}
```

### Example: Product Grid

```go
type Product struct {
    Name  string
    Price string
    Image string
}

func buildProductGrid(products []Product) goflow.Widget {
    items := make([]goflow.Widget, len(products))

    for i, product := range products {
        image := widgets.NewImage(&widgets.AssetImage{
            Path: product.Image,
        })
        image.Fit = widgets.ImageFitCover

        items[i] = material.NewCard(&widgets.Column{
            Children: []goflow.Widget{
                image,
                &widgets.Container{
                    Padding: goflow.NewEdgeInsets(8, 8, 8, 8),
                    Child: &widgets.Column{
                        Children: []goflow.Widget{
                            &widgets.Text{Data: product.Name},
                            &widgets.Text{Data: product.Price},
                        },
                        CrossAxisAlign: widgets.CrossAxisStart,
                    },
                },
            },
        })
    }

    grid := widgets.NewGridView(items, 2)
    grid.ChildAspectRatio = 0.8
    grid.MainAxisSpacing = 16.0
    grid.CrossAxisSpacing = 16.0
    grid.Padding = goflow.NewEdgeInsets(16, 16, 16, 16)

    return grid
}
```

### Example: Settings Grid

```go
func buildSettingsGrid() goflow.Widget {
    settings := []struct {
        Icon  widgets.IconData
        Label string
    }{
        {widgets.IconPerson, "Account"},
        {widgets.IconSettings, "Settings"},
        {widgets.IconWarning, "Privacy"},
        {widgets.IconInfo, "About"},
        {widgets.IconShare, "Share"},
        {widgets.IconPhone, "Contact"},
    }

    items := make([]goflow.Widget, len(settings))

    for i, setting := range settings {
        icon := widgets.NewIcon(setting.Icon)
        icon.Size = 48.0
        icon.Color = goflow.NewColor(33, 150, 243, 255)

        items[i] = &widgets.Container{
            Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
            Child: &widgets.Column{
                Children: []goflow.Widget{
                    icon,
                    &widgets.Text{Data: setting.Label},
                },
                MainAxisAlign: widgets.MainAxisCenter,
                CrossAxisAlign: widgets.CrossAxisCenter,
            },
        }
    }

    grid := widgets.NewGridView(items, 3)
    grid.MainAxisSpacing = 8.0
    grid.CrossAxisSpacing = 8.0

    return grid
}
```

## SingleChildScrollView

Scrollable container for a single widget.

```go
&widgets.SingleChildScrollView{
    Child: longContent,
}
```

### Properties

- `Child` - The scrollable child widget
- `ScrollDirection` - Scroll direction (default: Vertical)
- `Padding` - Padding around child
- `Physics` - Scroll behavior

### Basic Scrolling

```go
scrollView := widgets.NewSingleChildScrollView(longContent)
```

### Horizontal Scrolling

```go
scrollView := widgets.NewSingleChildScrollView(wideContent)
scrollView.ScrollDirection = widgets.AxisHorizontal
```

### Example: Long Form

```go
func buildLongForm() goflow.Widget {
    return widgets.NewSingleChildScrollView(&widgets.Column{
        Children: []goflow.Widget{
            &widgets.Text{Data: "User Information"},
            material.NewTextField(),
            material.NewTextField(),
            material.NewTextField(),
            &widgets.Text{Data: "Address"},
            material.NewTextField(),
            material.NewTextField(),
            material.NewTextField(),
            &widgets.Text{Data: "Additional Details"},
            material.NewTextField(),
            material.NewTextField(),
            material.NewButton(
                &widgets.Text{Data: "Submit"},
                onSubmit,
            ),
        },
    })
}
```

### Example: Scrollable Content

```go
func buildArticle(title, content string) goflow.Widget {
    return widgets.NewSingleChildScrollView(&widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
        Child: &widgets.Column{
            Children: []goflow.Widget{
                &widgets.Text{Data: title},
                &widgets.Container{
                    Height: floatPtr(16.0),
                },
                &widgets.Text{Data: content},
            },
            CrossAxisAlign: widgets.CrossAxisStart,
        },
    })
}
```

## Best Practices

### 1. Use ListView.builder for Large Lists

```go
// ✅ Good - efficient for large lists
&widgets.ListViewBuilder{
    ItemCount: 10000,
    ItemBuilder: func(ctx goflow.BuildContext, i int) goflow.Widget {
        return buildItem(i)
    },
}

// ❌ Avoid - creates all items upfront
&widgets.ListView{
    Children: make10000Items(),
}
```

### 2. Set Explicit Sizes for Scrollable Content

```go
// ✅ Good - items have explicit size
&widgets.ListView{
    Children: []goflow.Widget{
        &widgets.Container{
            Height: floatPtr(100.0),
            Child: item1,
        },
        &widgets.Container{
            Height: floatPtr(100.0),
            Child: item2,
        },
    },
}
```

### 3. Use Appropriate Scroll Physics

```go
// iOS app - bouncing
listView.Physics = widgets.ScrollPhysicsBouncing

// Android app - clamping
listView.Physics = widgets.ScrollPhysicsClamping

// Embedded list - disable scrolling
listView.Physics = widgets.ScrollPhysicsNeverScrollable
```

### 4. Add Padding for Visual Breathing Room

```go
listView := widgets.NewListView(items)
listView.Padding = goflow.NewEdgeInsets(16, 8, 16, 8)
```

### 5. Use GridView for Uniform Items

```go
// ✅ Good - items are uniform
grid := widgets.NewGridView(photos, 3)

// ❌ Avoid - items vary in size, use ListView instead
```

### 6. Limit Grid Columns

```go
// ✅ Good - readable on mobile
grid := widgets.NewGridView(items, 2)

// ❌ Too many - items too small
grid := widgets.NewGridView(items, 6)
```

### 7. Use ShrinkWrap for Nested Scrolling

```go
&widgets.SingleChildScrollView{
    Child: &widgets.Column{
        Children: []goflow.Widget{
            header,
            &widgets.ListView{
                Children: items,
                ShrinkWrap: true, // Don't take infinite height
                Physics: widgets.ScrollPhysicsNeverScrollable,
            },
            footer,
        },
    },
}
```

### 8. Add Empty States

```go
func buildList(items []Item) goflow.Widget {
    if len(items) == 0 {
        return &widgets.Center{
            Child: &widgets.Column{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconInfo),
                    &widgets.Text{Data: "No items found"},
                },
                MainAxisAlign: widgets.MainAxisCenter,
            },
        }
    }

    return widgets.NewListView(buildItems(items))
}
```

## Performance Tips

### 1. Reuse Widgets

```go
// ✅ Good - builder creates widgets on demand
&widgets.ListViewBuilder{
    ItemCount: count,
    ItemBuilder: buildItem,
}
```

### 2. Avoid Heavy Computations in Builders

```go
// ✅ Good - compute data outside builder
data := processData(rawData)

&widgets.ListViewBuilder{
    ItemCount: len(data),
    ItemBuilder: func(ctx goflow.BuildContext, i int) goflow.Widget {
        return buildSimpleItem(data[i])
    },
}

// ❌ Avoid - heavy computation in builder
&widgets.ListViewBuilder{
    ItemBuilder: func(ctx goflow.BuildContext, i int) goflow.Widget {
        processed := heavyProcessing(i) // Bad!
        return buildItem(processed)
    },
}
```

### 3. Use Const Sizes

```go
const itemHeight = 80.0

&widgets.ListViewBuilder{
    ItemBuilder: func(ctx goflow.BuildContext, i int) goflow.Widget {
        return &widgets.Container{
            Height: floatPtr(itemHeight),
            Child: buildItem(i),
        }
    },
}
```

## Helper Functions

```go
// Float pointer helper
func floatPtr(f float64) *float64 {
    return &f
}

// Build separator
func buildSeparator() goflow.Widget {
    return &widgets.Container{
        Height: floatPtr(1.0),
        Color: goflow.NewColor(224, 224, 224, 255),
    }
}

// Build section header
func buildSectionHeader(title string) goflow.Widget {
    style := goflow.NewTextStyle()
    style.FontSize = 18.0
    style.FontWeight = goflow.FontWeightBold

    return &widgets.Container{
        Padding: goflow.NewEdgeInsets(16, 8, 16, 8),
        Child: &widgets.Text{
            Data: title,
            Style: style,
        },
    }
}
```
