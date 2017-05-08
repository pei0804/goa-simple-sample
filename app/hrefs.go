// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": Application Resource Href Factories
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"fmt"
	"strings"
)

// AccountsHref returns the resource href.
func AccountsHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/accounts/users/%v", paramid)
}

// ResponseHref returns the resource href.
func ResponseHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/response/users/%v", paramid)
}
