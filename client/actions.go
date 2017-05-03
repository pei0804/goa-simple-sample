// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "tikasan/goa-simple-sample": actions Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// IDActionsPath computes a request path to the id action of actions.
func IDActionsPath(id string) string {
	param0 := id

	return fmt.Sprintf("/actions/%s", param0)
}

// 複数アクション（:id）
func (c *Client) IDActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewIDActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIDActionsRequest create the request corresponding to the id action endpoint of the actions resource.
func (c *Client) NewIDActionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// MainActionsPath computes a request path to the main action of actions.
func MainActionsPath() string {

	return fmt.Sprintf("/actions/main")
}

// 複数アクション（main）
func (c *Client) MainActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewMainActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMainActionsRequest create the request corresponding to the main action endpoint of the actions resource.
func (c *Client) NewMainActionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SubActionsPath computes a request path to the sub action of actions.
func SubActionsPath() string {

	return fmt.Sprintf("/actions/sub")
}

// 複数アクション（sub）
func (c *Client) SubActions(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSubActionsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubActionsRequest create the request corresponding to the sub action endpoint of the actions resource.
func (c *Client) NewSubActionsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
