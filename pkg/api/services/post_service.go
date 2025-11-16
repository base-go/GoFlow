package services

import (
	"context"
	"fmt"

	"github.com/base-go/GoFlow/pkg/api"
)

// Post represents a blog post from the API
type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Comment represents a comment on a post
type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// PostService provides methods to interact with post API endpoints
type PostService struct {
	client *api.Client
}

// NewPostService creates a new PostService
func NewPostService(client *api.Client) *PostService {
	return &PostService{client: client}
}

// GetPost fetches a single post by ID
func (s *PostService) GetPost(ctx context.Context, id int) *api.Resource[Post] {
	path := fmt.Sprintf("/posts/%d", id)
	return api.Get[Post](ctx, s.client, path)
}

// GetPosts fetches all posts
func (s *PostService) GetPosts(ctx context.Context) *api.ResourceList[Post] {
	return api.FetchList[Post](ctx, s.client, "GET", "/posts", nil)
}

// GetUserPosts fetches posts by a specific user
func (s *PostService) GetUserPosts(ctx context.Context, userID int) *api.ResourceList[Post] {
	path := fmt.Sprintf("/posts?userId=%d", userID)
	return api.FetchList[Post](ctx, s.client, "GET", path, nil)
}

// GetPostComments fetches comments for a specific post
func (s *PostService) GetPostComments(ctx context.Context, postID int) *api.ResourceList[Comment] {
	path := fmt.Sprintf("/posts/%d/comments", postID)
	return api.FetchList[Comment](ctx, s.client, "GET", path, nil)
}

// CreatePost creates a new post
func (s *PostService) CreatePost(ctx context.Context, post Post) *api.Resource[Post] {
	return api.Post[Post](ctx, s.client, "/posts", post)
}

// UpdatePost updates an existing post
func (s *PostService) UpdatePost(ctx context.Context, id int, post Post) *api.Resource[Post] {
	path := fmt.Sprintf("/posts/%d", id)
	return api.Put[Post](ctx, s.client, path, post)
}

// DeletePost deletes a post
func (s *PostService) DeletePost(ctx context.Context, id int) *api.Resource[interface{}] {
	path := fmt.Sprintf("/posts/%d", id)
	return api.Delete[interface{}](ctx, s.client, path)
}
