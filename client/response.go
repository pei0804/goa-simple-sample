// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": response Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// ArrayResponsePath computes a request path to the array action of response.
func ArrayResponsePath() string {

	return fmt.Sprintf("/api/v1/response/users/array")
}

// ユーザー（配列）
func (c *Client) ArrayResponse(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewArrayResponseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewArrayResponseRequest create the request corresponding to the array action endpoint of the response resource.
func (c *Client) NewArrayResponseRequest(ctx context.Context, path string) (*http.Request, error) {
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

// HashResponsePath computes a request path to the hash action of response.
func HashResponsePath() string {

	return fmt.Sprintf("/api/v1/response/users/hash")
}

// ユーザー（ハッシュ）
func (c *Client) HashResponse(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewHashResponseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewHashResponseRequest create the request corresponding to the hash action endpoint of the response resource.
func (c *Client) NewHashResponseRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListResponsePath computes a request path to the list action of response.
func ListResponsePath() string {

	return fmt.Sprintf("/api/v1/response/users")
}

// ユーザー（複数）
func (c *Client) ListResponse(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListResponseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListResponseRequest create the request corresponding to the list action endpoint of the response resource.
func (c *Client) NewListResponseRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowResponsePath computes a request path to the show action of response.
func ShowResponsePath(id string) string {
	param0 := id

	return fmt.Sprintf("/api/v1/response/users/%s", param0)
}

// ユーザー（単数）
func (c *Client) ShowResponse(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowResponseRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowResponseRequest create the request corresponding to the show action endpoint of the response resource.
func (c *Client) NewShowResponseRequest(ctx context.Context, path string) (*http.Request, error) {
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
