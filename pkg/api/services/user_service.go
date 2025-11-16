package services

import (
	"context"
	"fmt"

	"github.com/base-go/GoFlow/pkg/api"
)

// User represents a user from the API
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
	Company  struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		BS          string `json:"bs"`
	} `json:"company"`
	Address struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
}

// UserService provides methods to interact with user API endpoints
type UserService struct {
	client *api.Client
}

// NewUserService creates a new UserService
func NewUserService(client *api.Client) *UserService {
	return &UserService{client: client}
}

// GetUser fetches a single user by ID
func (s *UserService) GetUser(ctx context.Context, id int) *api.Resource[User] {
	path := fmt.Sprintf("/users/%d", id)
	return api.Get[User](ctx, s.client, path)
}

// GetUsers fetches all users
func (s *UserService) GetUsers(ctx context.Context) *api.ResourceList[User] {
	return api.FetchList[User](ctx, s.client, "GET", "/users", nil)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user User) *api.Resource[User] {
	return api.Post[User](ctx, s.client, "/users", user)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(ctx context.Context, id int, user User) *api.Resource[User] {
	path := fmt.Sprintf("/users/%d", id)
	return api.Put[User](ctx, s.client, path, user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int) *api.Resource[interface{}] {
	path := fmt.Sprintf("/users/%d", id)
	return api.Delete[interface{}](ctx, s.client, path)
}
