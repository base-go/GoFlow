package api

import (
	"github.com/base-go/GoFlow/pkg/core/framework"
)

// contextKey is used to store the API client in BuildContext
type contextKey string

const apiClientKey contextKey = "api_client"

// APIProvider provides an API client to its descendants
type APIProvider struct {
	framework.BaseWidget
	Client *Client
	Child  framework.Widget
}

// Build returns the child widget with API client in context
func (p *APIProvider) Build(ctx framework.BuildContext) framework.Widget {
	// Store client in context
	ctx.SetValue(apiClientKey, p.Client)
	return p.Child
}

// GetClient retrieves the API client from context
func GetClient(ctx framework.BuildContext) *Client {
	if client, ok := ctx.GetValue(apiClientKey).(*Client); ok {
		return client
	}
	return nil
}
