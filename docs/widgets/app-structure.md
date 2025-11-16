# App Structure Widgets

Widgets that provide the basic structure for applications.

## Scaffold

Material Design app structure with AppBar, Body, FAB, Drawer, and more.

```go
&material.Scaffold{
    AppBar: material.NewAppBar(&widgets.Text{Data: "My App"}),
    Body: &widgets.Center{
        Child: &widgets.Text{Data: "Hello, World!"},
    },
    FloatingActionButton: material.NewFloatingActionButton(
        widgets.NewIcon(widgets.IconAdd),
        onAdd,
    ),
    Drawer: buildDrawer(),
    BottomNavigationBar: buildBottomNav(),
}
```

### Properties

- `AppBar` - Top app bar
- `Body` - Main content
- `FloatingActionButton` - FAB widget
- `FloatingActionButtonLocation` - FAB position
- `Drawer` - Side navigation drawer
- `BottomNavigationBar` - Bottom navigation
- `BottomSheet` - Bottom sheet widget
- `BackgroundColor` - Background color
- `ResizeToAvoidBottomInset` - Resize when keyboard appears

### FAB Locations

- `FABLocationEndFloat` - Bottom right (default)
- `FABLocationCenterFloat` - Bottom center
- `FABLocationStartFloat` - Bottom left
- `FABLocationEndDocked` - Docked to bottom bar, right
- `FABLocationCenterDocked` - Docked to bottom bar, center

### Example: Full Scaffold

```go
&material.Scaffold{
    AppBar: material.NewAppBar(&widgets.Text{
        Data: "Dashboard",
    }),
    Body: &widgets.ListView{
        Children: buildDashboardItems(),
    },
    FloatingActionButton: material.NewFloatingActionButton(
        widgets.NewIcon(widgets.IconAdd),
        func() {
            showCreateDialog()
        },
    ),
    FloatingActionButtonLocation: material.FABLocationEndFloat,
    Drawer: buildNavigationDrawer(),
    BottomNavigationBar: buildBottomNavigationBar(),
    BackgroundColor: goflow.NewColor(250, 250, 250, 255),
}
```

## CupertinoPageScaffold

iOS-style page structure.

```go
&cupertino.CupertinoPageScaffold{
    NavigationBar: cupertino.NewNavigationBar(&widgets.Text{
        Data: "My App",
    }),
    Child: &widgets.Center{
        Child: &widgets.Text{Data: "Hello, iOS!"},
    },
}
```

### Properties

- `NavigationBar` - Top navigation bar
- `Child` - Main content
- `BackgroundColor` - Background color
- `ResizeToAvoidBottomInset` - Resize when keyboard appears

### Example: iOS Page

```go
&cupertino.CupertinoPageScaffold{
    NavigationBar: cupertino.NewNavigationBar(&widgets.Text{
        Data: "Settings",
    }),
    Child: &widgets.ListView{
        Children: buildSettingsItems(),
    },
    BackgroundColor: cupertino.DefaultLightTheme().BackgroundColor,
}
```

## CupertinoTabScaffold

iOS-style tabbed interface.

```go
&cupertino.CupertinoTabScaffold{
    TabBar: buildTabBar(),
    TabBuilder: func(ctx goflow.BuildContext, index int) goflow.Widget {
        switch index {
        case 0:
            return buildHomeTab()
        case 1:
            return buildSearchTab()
        case 2:
            return buildProfileTab()
        default:
            return &widgets.Container{}
        }
    },
    CurrentIndex: 0,
}
```

### Properties

- `TabBar` - Tab bar widget
- `TabBuilder` - Function to build tab content
- `CurrentIndex` - Currently active tab
- `BackgroundColor` - Background color

### Example: Tabbed App

```go
currentTab := signals.New(0)

&cupertino.CupertinoTabScaffold{
    TabBar: buildTabBar(currentTab),
    TabBuilder: func(ctx goflow.BuildContext, index int) goflow.Widget {
        tabs := []goflow.Widget{
            buildHomeTab(),
            buildSearchTab(),
            buildProfileTab(),
        }
        if index < len(tabs) {
            return tabs[index]
        }
        return &widgets.Container{}
    },
    CurrentIndex: currentTab.Get(),
}
```

## AppBar

Material Design top app bar.

```go
appBar := material.NewAppBar(&widgets.Text{Data: "Title"})
appBar.Leading = widgets.NewIconButton(widgets.IconMenu, onMenuPressed)
appBar.Actions = []goflow.Widget{
    widgets.NewIconButton(widgets.IconSearch, onSearch),
    widgets.NewIconButton(widgets.IconSettings, onSettings),
}
```

### Properties

- `Title` - Title widget (usually Text)
- `Leading` - Leading widget (usually menu or back button)
- `Actions` - Action widgets (usually icon buttons)
- `BackgroundColor` - Background color
- `ForegroundColor` - Text/icon color
- `Elevation` - Shadow elevation
- `Height` - AppBar height (default 56.0)

### Example: App Bar with Search

```go
isSearching := signals.New(false)

appBar := material.NewAppBar(
    func() goflow.Widget {
        if isSearching.Get() {
            return buildSearchField()
        }
        return &widgets.Text{Data: "My App"}
    }(),
)
appBar.Actions = []goflow.Widget{
    widgets.NewIconButton(widgets.IconSearch, func() {
        isSearching.Set(!isSearching.Get())
    }),
}
```

## CupertinoNavigationBar

iOS-style navigation bar.

```go
navBar := cupertino.NewNavigationBar(&widgets.Text{Data: "Title"})
navBar.Leading = widgets.NewIconButton(widgets.IconArrowBack, onBack)
navBar.Trailing = widgets.NewIconButton(widgets.IconEdit, onEdit)
```

### Properties

- `Middle` - Middle widget (title)
- `Leading` - Leading widget (back button)
- `Trailing` - Trailing widget (action button)
- `BackgroundColor` - Background color
- `BorderColor` - Border color
- `Padding` - Internal padding
- `Height` - Bar height (default 44.0)

## Drawer

Side navigation drawer.

```go
drawer := material.NewDrawer(&widgets.Column{
    Children: []goflow.Widget{
        material.NewDrawerHeader(&widgets.Column{
            Children: []goflow.Widget{
                &widgets.Text{Data: "John Doe"},
                &widgets.Text{Data: "john@example.com"},
            },
        }),
        material.NewListTile(&widgets.Text{Data: "Home"}),
        material.NewListTile(&widgets.Text{Data: "Settings"}),
        material.NewListTile(&widgets.Text{Data: "Logout"}),
    },
})
```

### Properties

- `Child` - Drawer content
- `BackgroundColor` - Background color
- `Elevation` - Shadow elevation
- `Width` - Drawer width (default 304.0)

### Example: Navigation Drawer

```go
func buildDrawer() goflow.Widget {
    return material.NewDrawer(&widgets.Column{
        Children: []goflow.Widget{
            // Header
            material.NewDrawerHeader(&widgets.Container{
                Padding: goflow.NewEdgeInsets(16, 16, 16, 16),
                Child: &widgets.Column{
                    Children: []goflow.Widget{
                        &widgets.Icon{
                            Icon: widgets.IconPerson,
                            Size: 48.0,
                        },
                        &widgets.Text{Data: "John Doe"},
                        &widgets.Text{Data: "john@example.com"},
                    },
                    MainAxisAlign: widgets.MainAxisCenter,
                },
            }),

            // Navigation items
            material.NewListTile(&widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconHome),
                    &widgets.Text{Data: "Home"},
                },
            }),
            material.NewListTile(&widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconSettings),
                    &widgets.Text{Data: "Settings"},
                },
            }),

            // Divider
            &widgets.Container{
                Height: floatPtr(1.0),
                Color: goflow.NewColor(224, 224, 224, 255),
            },

            material.NewListTile(&widgets.Row{
                Children: []goflow.Widget{
                    widgets.NewIcon(widgets.IconLogout),
                    &widgets.Text{Data: "Logout"},
                },
            }),
        },
    })
}
```

## BottomNavigationBar

Material Design bottom navigation.

```go
currentIndex := signals.New(0)

bottomNav := material.NewBottomNavigationBar([]material.BottomNavigationBarItem{
    {Icon: widgets.IconHome, Label: "Home"},
    {Icon: widgets.IconSearch, Label: "Search"},
    {Icon: widgets.IconPerson, Label: "Profile"},
})
bottomNav.CurrentIndex = currentIndex.Get()
bottomNav.OnTap = func(index int) {
    currentIndex.Set(index)
}
```

### Properties

- `Items` - Navigation items
- `CurrentIndex` - Currently selected index
- `OnTap` - Callback when item tapped
- `BackgroundColor` - Background color
- `SelectedItemColor` - Selected item color
- `UnselectedItemColor` - Unselected item color
- `ShowSelectedLabels` - Show labels for selected items
- `ShowUnselectedLabels` - Show labels for unselected items
- `Type` - Navigation bar type
- `Elevation` - Shadow elevation

### Bottom Navigation Bar Item

```go
material.BottomNavigationBarItem{
    Icon: widgets.IconHome,
    ActiveIcon: &widgets.IconHomeFilled, // Optional
    Label: "Home",
    BackgroundColor: nil, // Optional
    Tooltip: "Go to home", // Optional
}
```

### Example: Complete Bottom Nav

```go
func buildBottomNavigationBar(currentIndex *signals.Signal[int]) goflow.Widget {
    items := []material.BottomNavigationBarItem{
        {Icon: widgets.IconHome, Label: "Home"},
        {Icon: widgets.IconSearch, Label: "Search"},
        {Icon: widgets.IconFavorite, Label: "Favorites"},
        {Icon: widgets.IconPerson, Label: "Profile"},
    }

    bottomNav := material.NewBottomNavigationBar(items)
    bottomNav.CurrentIndex = currentIndex.Get()
    bottomNav.OnTap = func(index int) {
        currentIndex.Set(index)
    }
    bottomNav.SelectedItemColor = goflow.NewColor(33, 150, 243, 255)

    return bottomNav
}
```

## Best Practices

### 1. Use Appropriate Structure for Platform

```go
// Material (Android, Web, Desktop)
&material.Scaffold{
    AppBar: appBar,
    Body: body,
}

// Cupertino (iOS, macOS)
&cupertino.CupertinoPageScaffold{
    NavigationBar: navBar,
    Child: body,
}

// Adaptive (Recommended)
func buildPage() goflow.Widget {
    if goflow.GetPlatform().IsApple() {
        return &cupertino.CupertinoPageScaffold{...}
    }
    return &material.Scaffold{...}
}
```

### 2. Limit Bottom Navigation Items

```go
// ✅ Good - 3-5 items
items := []material.BottomNavigationBarItem{
    {Icon: widgets.IconHome, Label: "Home"},
    {Icon: widgets.IconSearch, Label: "Search"},
    {Icon: widgets.IconPerson, Label: "Profile"},
}

// ❌ Too many - use drawer instead
items := []material.BottomNavigationBarItem{
    // 7 items - overwhelming
}
```

### 3. Provide Clear Navigation

```go
appBar := material.NewAppBar(&widgets.Text{Data: "Details"})
appBar.Leading = widgets.NewIconButton(
    widgets.IconArrowBack,
    func() {
        navigator.Pop()
    },
)
```

### 4. Use FAB for Primary Action

```go
// ✅ Good - one primary action
fab := material.NewFloatingActionButton(
    widgets.NewIcon(widgets.IconAdd),
    createNew,
)

// ❌ Avoid - multiple FABs confuse users
```

### 5. Keep Drawer Organized

```go
func buildDrawer() goflow.Widget {
    return material.NewDrawer(&widgets.Column{
        Children: []goflow.Widget{
            // 1. Header
            buildDrawerHeader(),

            // 2. Primary navigation
            buildNavigationItems(),

            // 3. Divider
            buildDivider(),

            // 4. Secondary actions
            buildSecondaryActions(),
        },
    })
}
```
