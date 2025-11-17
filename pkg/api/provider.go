package api

import (
	goflow "github.com/base-go/GoFlow/pkg/core/framework"
)

// contextKey is used to store the API client in BuildContext
type contextKey string

const apiClientKey contextKey = "api_client"

// APIProvider provides an API client to its descendants
type APIProvider struct {
	goflow.BaseWidget
	Client *Client
	Child  goflow.Widget
}

// Build returns the child widget with API client in context
func (p *APIProvider) Build(ctx goflow.BuildContext) goflow.Widget {
	// TODO: Store client in context when BuildContext supports SetValue/GetValue
	// ctx.SetValue(apiClientKey, p.Client)
	return p.Child
}

// GetClient retrieves the API client from context
func GetClient(ctx goflow.BuildContext) *Client {
	// TODO: Implement when BuildContext supports GetValue
	// if client, ok := ctx.GetValue(apiClientKey).(*Client); ok {
	// 	return client
	// }
	return nil
}
