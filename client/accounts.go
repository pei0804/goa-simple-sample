// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": accounts Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// AddAccountsPath computes a request path to the add action of accounts.
func AddAccountsPath() string {

	return fmt.Sprintf("/api/v1/accounts")
}

// 追加
func (c *Client) AddAccounts(ctx context.Context, path string, email *string, name *string) (*http.Response, error) {
	req, err := c.NewAddAccountsRequest(ctx, path, email, name)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddAccountsRequest create the request corresponding to the add action endpoint of the accounts resource.
func (c *Client) NewAddAccountsRequest(ctx context.Context, path string, email *string, name *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if email != nil {
		values.Set("email", *email)
	}
	if name != nil {
		values.Set("name", *name)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeleteAccountsPath computes a request path to the delete action of accounts.
func DeleteAccountsPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/accounts/users/%s", param0)
}

// 削除
func (c *Client) DeleteAccounts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteAccountsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAccountsRequest create the request corresponding to the delete action endpoint of the accounts resource.
func (c *Client) NewDeleteAccountsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListAccountsPath computes a request path to the list action of accounts.
func ListAccountsPath() string {

	return fmt.Sprintf("/api/v1/accounts")
}

// 複数
func (c *Client) ListAccounts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListAccountsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListAccountsRequest create the request corresponding to the list action endpoint of the accounts resource.
func (c *Client) NewListAccountsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ShowAccountsPath computes a request path to the show action of accounts.
func ShowAccountsPath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/api/v1/accounts/%s", param0)
}

// 単数
func (c *Client) ShowAccounts(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowAccountsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowAccountsRequest create the request corresponding to the show action endpoint of the accounts resource.
func (c *Client) NewShowAccountsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateAccountsPath computes a request path to the update action of accounts.
func UpdateAccountsPath(id string) string {
	param0 := id

	return fmt.Sprintf("/api/v1/accounts/users/%s", param0)
}

// 更新
func (c *Client) UpdateAccounts(ctx context.Context, path string, email *string, name *string) (*http.Response, error) {
	req, err := c.NewUpdateAccountsRequest(ctx, path, email, name)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAccountsRequest create the request corresponding to the update action endpoint of the accounts resource.
func (c *Client) NewUpdateAccountsRequest(ctx context.Context, path string, email *string, name *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if email != nil {
		values.Set("email", *email)
	}
	if name != nil {
		values.Set("name", *name)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
