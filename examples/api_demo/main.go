package main

import (
	"context"
	"fmt"

	"github.com/base-go/GoFlow/pkg/api"
	"github.com/base-go/GoFlow/pkg/api/middleware"
	"github.com/base-go/GoFlow/pkg/api/services"
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

func main() {
	// Create API client with middleware
	apiClient := api.NewClient(api.ClientConfig{
		BaseURL: "https://jsonplaceholder.typicode.com",
	}).Use(middleware.Logging(nil))

	// Create and run the app
	app := &APIDemo{
		apiClient: apiClient,
	}

	goflow.RunApp(app)
}

// APIDemo is the main application widget
type APIDemo struct {
	goflow.BaseWidget
	apiClient *api.Client
}

func (a *APIDemo) Build(ctx goflow.BuildContext) goflow.Widget {
	return &api.APIProvider{
		Client: a.apiClient,
		Child:  &HomePage{},
	}
}

// HomePage displays a list of users
type HomePage struct {
	goflow.BaseWidget
	userResource *api.ResourceList[services.User]
	initialized  *signals.Signal[bool]
}

func (h *HomePage) Build(ctx goflow.BuildContext) goflow.Widget {
	// Initialize on first build
	if h.initialized == nil {
		h.initialized = signals.NewSignal(false)
		h.loadUsers(ctx)
	}

	// Get the resource state
	state := h.userResource.Get()

	// Show loading state
	if state.Loading {
		return &widgets.Center{
			Child: &widgets.Column{
				MainAxisAlignment: widgets.MainAxisAlignmentCenter,
				Children: []goflow.Widget{
					&widgets.Text{
						Text: "Loading users...",
						Style: &widgets.TextStyle{
							FontSize: goflow.Float64(20),
							Color:    goflow.NewColor(100, 100, 100, 255),
						},
					},
				},
			},
		}
	}

	// Show error state
	if state.Error != nil {
		return &widgets.Center{
			Child: &widgets.Column{
				MainAxisAlignment: widgets.MainAxisAlignmentCenter,
				Children: []goflow.Widget{
					&widgets.Text{
						Text: "Error loading users",
						Style: &widgets.TextStyle{
							FontSize: goflow.Float64(20),
							Color:    goflow.NewColor(200, 50, 50, 255),
						},
					},
					&widgets.SizedBox{Height: goflow.Float64(10)},
					&widgets.Text{
						Text: state.Error.Error(),
						Style: &widgets.TextStyle{
							FontSize: goflow.Float64(14),
							Color:    goflow.NewColor(150, 50, 50, 255),
						},
					},
				},
			},
		}
	}

	// Show user list
	users := state.Data
	if users == nil || len(*users) == 0 {
		return &widgets.Center{
			Child: &widgets.Text{
				Text: "No users found",
				Style: &widgets.TextStyle{
					FontSize: goflow.Float64(20),
					Color:    goflow.NewColor(100, 100, 100, 255),
				},
			},
		}
	}

	// Build user list
	userWidgets := []goflow.Widget{
		&widgets.Text{
			Text: "Users from API",
			Style: &widgets.TextStyle{
				FontSize:   goflow.Float64(28),
				FontWeight: widgets.FontWeightBold,
				Color:      goflow.NewColor(0, 0, 0, 255),
			},
		},
		&widgets.SizedBox{Height: goflow.Float64(20)},
	}

	for _, user := range *users {
		userWidgets = append(userWidgets, &UserCard{User: user})
		userWidgets = append(userWidgets, &widgets.SizedBox{Height: goflow.Float64(10)})
	}

	return &widgets.Padding{
		Padding: goflow.NewEdgeInsets(20, 20, 20, 20),
		Child: &widgets.SingleChildScrollView{
			Child: &widgets.Column{
				CrossAxisAlignment: widgets.CrossAxisAlignmentStart,
				Children:           userWidgets,
			},
		},
	}
}

func (h *HomePage) loadUsers(ctx goflow.BuildContext) {
	apiClient := api.GetClient(ctx)
	if apiClient == nil {
		return
	}

	userService := services.NewUserService(apiClient)
	h.userResource = userService.GetUsers(context.Background())
	h.initialized.Set(true)
}

// UserCard displays a single user
type UserCard struct {
	goflow.BaseWidget
	User services.User
}

func (u *UserCard) Build(ctx goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Padding: goflow.NewEdgeInsets(15, 15, 15, 15),
		Color:   goflow.NewColor(245, 245, 245, 255),
		Child: &widgets.Column{
			CrossAxisAlignment: widgets.CrossAxisAlignmentStart,
			Children: []goflow.Widget{
				&widgets.Text{
					Text: u.User.Name,
					Style: &widgets.TextStyle{
						FontSize:   goflow.Float64(18),
						FontWeight: widgets.FontWeightBold,
						Color:      goflow.NewColor(0, 0, 0, 255),
					},
				},
				&widgets.SizedBox{Height: goflow.Float64(5)},
				&widgets.Text{
					Text: fmt.Sprintf("@%s • %s", u.User.Username, u.User.Email),
					Style: &widgets.TextStyle{
						FontSize: goflow.Float64(14),
						Color:    goflow.NewColor(100, 100, 100, 255),
					},
				},
				&widgets.SizedBox{Height: goflow.Float64(5)},
				&widgets.Text{
					Text: u.User.Company.Name,
					Style: &widgets.TextStyle{
						FontSize: goflow.Float64(14),
						Color:    goflow.NewColor(50, 100, 200, 255),
					},
				},
			},
		},
	}
}
