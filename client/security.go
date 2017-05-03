// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "tikasan/goa-simple-sample": security Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// SecuritySecurityPath computes a request path to the security action of security.
func SecuritySecurityPath() string {

	return fmt.Sprintf("/securiy")
}

// セキュリティの例です
func (c *Client) SecuritySecurity(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSecuritySecurityRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSecuritySecurityRequest create the request corresponding to the security action endpoint of the security resource.
func (c *Client) NewSecuritySecurityRequest(ctx context.Context, path string) (*http.Request, error) {
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
