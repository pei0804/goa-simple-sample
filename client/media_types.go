// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "tikasan/goa-simple-sample": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// example (default view)
//
// Identifier: application/vnd.arraytype+json; view=default
type Arraytype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// 値
	Value string `form:"value" json:"value" xml:"value"`
}

// Validate validates the Arraytype media type instance.
func (mt *Arraytype) Validate() (err error) {

	if mt.Value == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	return
}

// DecodeArraytype decodes the Arraytype instance encoded in resp body.
func (c *Client) DecodeArraytype(resp *http.Response) (*Arraytype, error) {
	var decoded Arraytype
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ArraytypeCollection is the media type for an array of Arraytype (default view)
//
// Identifier: application/vnd.arraytype+json; type=collection; view=default
type ArraytypeCollection []*Arraytype

// Validate validates the ArraytypeCollection media type instance.
func (mt ArraytypeCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeArraytypeCollection decodes the ArraytypeCollection instance encoded in resp body.
func (c *Client) DecodeArraytypeCollection(resp *http.Response) (ArraytypeCollection, error) {
	var decoded ArraytypeCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// example (default view)
//
// Identifier: application/vnd.integertype+json; view=default
type Integertype struct {
	// 数値
	ID int `form:"ID" json:"ID" xml:"ID"`
}

// DecodeIntegertype decodes the Integertype instance encoded in resp body.
func (c *Client) DecodeIntegertype(resp *http.Response) (*Integertype, error) {
	var decoded Integertype
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// example (default view)
//
// Identifier: application/vnd.messagetype+json; view=default
type Messagetype struct {
	// メッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// Validate validates the Messagetype media type instance.
func (mt *Messagetype) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// DecodeMessagetype decodes the Messagetype instance encoded in resp body.
func (c *Client) DecodeMessagetype(resp *http.Response) (*Messagetype, error) {
	var decoded Messagetype
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// example (default view)
//
// Identifier: application/vnd.validationtype+json; view=default
type Validationtype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// デフォルト値
	Default string `form:"default" json:"default" xml:"default"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 列挙型
	Enum string `form:"enum" json:"enum" xml:"enum"`
	// 数字（1〜10）
	Integer int `form:"integer" json:"integer" xml:"integer"`
	// 文字（1~10文字）
	String string `form:"string" json:"string" xml:"string"`
}

// Validate validates the Validationtype media type instance.
func (mt *Validationtype) Validate() (err error) {

	if mt.String == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "string"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.Enum == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "enum"))
	}
	if mt.Default == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "default"))
	}
	return
}

// DecodeValidationtype decodes the Validationtype instance encoded in resp body.
func (c *Client) DecodeValidationtype(resp *http.Response) (*Validationtype, error) {
	var decoded Validationtype
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// example (default view)
//
// Identifier: application/vnd.viewtype+json; view=default
type Viewtype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// 値
	Value string `form:"value" json:"value" xml:"value"`
}

// Validate validates the Viewtype media type instance.
func (mt *Viewtype) Validate() (err error) {

	if mt.Value == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	return
}

// example (tiny view)
//
// Identifier: application/vnd.viewtype+json; view=tiny
type ViewtypeTiny struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
}

// DecodeViewtype decodes the Viewtype instance encoded in resp body.
func (c *Client) DecodeViewtype(resp *http.Response) (*Viewtype, error) {
	var decoded Viewtype
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeViewtypeTiny decodes the ViewtypeTiny instance encoded in resp body.
func (c *Client) DecodeViewtypeTiny(resp *http.Response) (*ViewtypeTiny, error) {
	var decoded ViewtypeTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
