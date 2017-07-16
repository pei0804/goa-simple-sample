// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.2.0-dirty

package app

import (
	"fmt"
	"strings"
)

// AccountsHref returns the resource href.
func AccountsHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/accounts/%v", paramid)
}

// AccountsDataHref returns the resource href.
func AccountsDataHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/accounts_data/%v", paramid)
}

// BottlesHref returns the resource href.
func BottlesHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/bottles/%v", paramid)
}

// BottlesDataHref returns the resource href.
func BottlesDataHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/bottles_data/%v", paramid)
}

// ResponseHref returns the resource href.
func ResponseHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/api/v1/response/users/%v", paramid)
}
