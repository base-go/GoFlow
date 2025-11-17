package main

import (
	"context"
	"fmt"

	"github.com/base-go/GoFlow/pkg/api"
	"github.com/base-go/GoFlow/pkg/api/services"
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
	"github.com/base-go/GoFlow/pkg/core/signals"
	"github.com/base-go/GoFlow/pkg/core/widgets"
)

// PostsPage displays posts for a specific user
type PostsPage struct {
	goflow.BaseWidget
	UserID          int
	postsResource   *api.ResourceList[services.Post]
	selectedPostID  *signals.Signal[int]
	commentsVisible *signals.Signal[bool]
	initialized     *signals.Signal[bool]
}

func (p *PostsPage) Build(ctx goflow.BuildContext) goflow.Widget {
	// Initialize on first build
	if p.initialized == nil {
		p.initialized = signals.New(false)
		p.selectedPostID = signals.New(-1)
		p.commentsVisible = signals.New(false)
		p.loadPosts(ctx)
	} else {
		p.initialized.Set(true)
	}

	state := p.postsResource.Get()

	// Show loading state
	if state.Loading {
		return &widgets.Center{
			Child: widgets.NewTextWithStyle(
				"Loading posts...",
				&goflow.TextStyle{
					FontSize: 18,
					Color:    goflow.NewColor(100, 100, 100, 255),
				},
			),
		}
	}

	// Show error state
	if state.Error != nil {
		return &widgets.Center{
			Child: &widgets.Column{
				MainAxisAlign: widgets.MainAxisCenter,
				Children: []goflow.Widget{
					widgets.NewTextWithStyle(
						"Error loading posts",
						&goflow.TextStyle{
							FontSize: 18,
							Color:    goflow.NewColor(200, 50, 50, 255),
						},
					),
					&widgets.SizedBox{Height: floatPtr(10)},
					widgets.NewTextWithStyle(
						state.Error.Error(),
						&goflow.TextStyle{
							FontSize: 12,
							Color:    goflow.NewColor(150, 50, 50, 255),
						},
					),
				},
			},
		}
	}

	// Build posts list
	posts := state.Data
	if posts == nil || len(*posts) == 0 {
		return &widgets.Center{
			Child: widgets.NewTextWithStyle(
				"No posts found",
				&goflow.TextStyle{
					FontSize: 18,
					Color:    goflow.NewColor(100, 100, 100, 255),
				},
			),
		}
	}

	postWidgets := []goflow.Widget{
		widgets.NewTextWithStyle(
			fmt.Sprintf("Posts by User %d", p.UserID),
			&goflow.TextStyle{
				FontSize:   24,
				FontWeight: goflow.FontWeightBold,
				Color:      goflow.NewColor(0, 0, 0, 255),
			},
		),
		&widgets.SizedBox{Height: floatPtr(20)},
	}

	selectedID := p.selectedPostID.Get()
	showComments := p.commentsVisible.Get()

	for _, post := range *posts {
		isSelected := post.ID == selectedID
		postWidgets = append(postWidgets, &PostCard{
			Post:       post,
			IsSelected: isSelected,
			OnTap: func() {
				if isSelected {
					p.commentsVisible.Set(!showComments)
				} else {
					p.selectedPostID.Set(post.ID)
					p.commentsVisible.Set(true)
				}
			},
		})

		// Show comments if this post is selected and comments are visible
		if isSelected && showComments {
			postWidgets = append(postWidgets, &CommentsWidget{
				PostID: post.ID,
			})
		}

		postWidgets = append(postWidgets, &widgets.SizedBox{Height: floatPtr(10)})
	}

	return &widgets.Padding{
		Padding: goflow.NewEdgeInsets(20, 20, 20, 20),
		Child: &widgets.SingleChildScrollView{
			Child: &widgets.Column{
				CrossAxisAlign: widgets.CrossAxisStart,
				Children:       postWidgets,
			},
		},
	}
}

func (p *PostsPage) loadPosts(ctx goflow.BuildContext) {
	apiClient := api.GetClient(ctx)
	if apiClient == nil {
		return
	}

	postService := services.NewPostService(apiClient)
	p.postsResource = postService.GetUserPosts(context.Background(), p.UserID)
	p.initialized.Set(true)
}

// PostCard displays a single post
type PostCard struct {
	goflow.BaseWidget
	Post       services.Post
	IsSelected bool
	OnTap      func()
}

func (p *PostCard) Build(ctx goflow.BuildContext) goflow.Widget {
	bgColor := goflow.NewColor(245, 245, 245, 255)
	if p.IsSelected {
		bgColor = goflow.NewColor(230, 240, 255, 255)
	}

	return &widgets.GestureDetector{
		OnTap: p.OnTap,
		Child: &widgets.Container{
			Padding: goflow.NewEdgeInsets(15, 15, 15, 15),
			Color:   bgColor,
			Child: &widgets.Column{
				CrossAxisAlign: widgets.CrossAxisStart,
				Children: []goflow.Widget{
					widgets.NewTextWithStyle(
						p.Post.Title,
						&goflow.TextStyle{
							FontSize:   16,
							FontWeight: goflow.FontWeightBold,
							Color:      goflow.NewColor(0, 0, 0, 255),
						},
					),
					&widgets.SizedBox{Height: floatPtr(8)},
					widgets.NewTextWithStyle(
						p.Post.Body,
						&goflow.TextStyle{
							FontSize: 14,
							Color:    goflow.NewColor(60, 60, 60, 255),
						},
					),
				},
			},
		},
	}
}

// CommentsWidget displays comments for a post
type CommentsWidget struct {
	goflow.BaseWidget
	PostID           int
	commentsResource *api.ResourceList[services.Comment]
	initialized      *signals.Signal[bool]
}

func (c *CommentsWidget) Build(ctx goflow.BuildContext) goflow.Widget {
	// Initialize on first build
	if c.initialized == nil {
		c.initialized = signals.New(false)
		c.loadComments(ctx)
	}

	state := c.commentsResource.Get()

	// Show loading state
	if state.Loading {
		return &widgets.Padding{
			Padding: goflow.NewEdgeInsets(15, 15, 15, 15),
			Child: widgets.NewTextWithStyle(
				"Loading comments...",
				&goflow.TextStyle{
					FontSize: 14,
					Color:    goflow.NewColor(100, 100, 100, 255),
				},
			),
		}
	}

	// Show error state
	if state.Error != nil {
		return &widgets.Padding{
			Padding: goflow.NewEdgeInsets(15, 15, 15, 15),
			Child: widgets.NewTextWithStyle(
				fmt.Sprintf("Error: %v", state.Error),
				&goflow.TextStyle{
					FontSize: 14,
					Color:    goflow.NewColor(200, 50, 50, 255),
				},
			),
		}
	}

	// Build comments
	comments := state.Data
	if comments == nil || len(*comments) == 0 {
		return &widgets.Padding{
			Padding: goflow.NewEdgeInsets(15, 15, 15, 15),
			Child: widgets.NewTextWithStyle(
				"No comments",
				&goflow.TextStyle{
					FontSize: 14,
					Color:    goflow.NewColor(100, 100, 100, 255),
				},
			),
		}
	}

	commentWidgets := []goflow.Widget{
		widgets.NewTextWithStyle(
			"Comments:",
			&goflow.TextStyle{
				FontSize:   14,
				FontWeight: goflow.FontWeightBold,
				Color:      goflow.NewColor(50, 50, 50, 255),
			},
		),
		&widgets.SizedBox{Height: floatPtr(10)},
	}

	for _, comment := range *comments {
		commentWidgets = append(commentWidgets, &CommentCard{Comment: comment})
		commentWidgets = append(commentWidgets, &widgets.SizedBox{Height: floatPtr(8)})
	}

	return &widgets.Container{
		Padding: goflow.NewEdgeInsets(15, 15, 15, 15),
		Color:   goflow.NewColor(255, 255, 255, 255),
		Child: &widgets.Column{
			CrossAxisAlign: widgets.CrossAxisStart,
			Children:       commentWidgets,
		},
	}
}

func (c *CommentsWidget) loadComments(ctx goflow.BuildContext) {
	apiClient := api.GetClient(ctx)
	if apiClient == nil {
		return
	}

	postService := services.NewPostService(apiClient)
	c.commentsResource = postService.GetPostComments(context.Background(), c.PostID)
	c.initialized.Set(true)
}

// CommentCard displays a single comment
type CommentCard struct {
	goflow.BaseWidget
	Comment services.Comment
}

func (c *CommentCard) Build(ctx goflow.BuildContext) goflow.Widget {
	return &widgets.Container{
		Padding: goflow.NewEdgeInsets(10, 10, 10, 10),
		Color:   goflow.NewColor(250, 250, 250, 255),
		Child: &widgets.Column{
			CrossAxisAlign: widgets.CrossAxisStart,
			Children: []goflow.Widget{
				widgets.NewTextWithStyle(
					fmt.Sprintf("%s (%s)", c.Comment.Name, c.Comment.Email),
					&goflow.TextStyle{
						FontSize:   12,
						FontWeight: goflow.FontWeightBold,
						Color:      goflow.NewColor(50, 100, 200, 255),
					},
				),
				&widgets.SizedBox{Height: floatPtr(5)},
				widgets.NewTextWithStyle(
					c.Comment.Body,
					&goflow.TextStyle{
						FontSize: 12,
						Color:    goflow.NewColor(60, 60, 60, 255),
					},
				),
			},
		},
	}
}
