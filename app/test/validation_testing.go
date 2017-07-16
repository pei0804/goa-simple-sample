// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": validation TestHelpers
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/pei0804/goa-simple-sample/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// ValidationValidationBadRequest runs the method Validation of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ValidationValidationBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ValidationController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/validation"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ValidationTest"), rw, req, prms)
	validationCtx, _err := app.NewValidationValidationContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Validation(validationCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ValidationValidationOK runs the method Validation of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ValidationValidationOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ValidationController, defaultType string, email string, enumType string, id int, integerType int, reg string, stringType string) (http.ResponseWriter, *app.Validation) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{defaultType}
		query["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		query["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		query["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		query["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		query["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		query["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		query["stringType"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/validation"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{defaultType}
		prms["defaultType"] = sliceVal
	}
	{
		sliceVal := []string{email}
		prms["email"] = sliceVal
	}
	{
		sliceVal := []string{enumType}
		prms["enumType"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(id)}
		prms["id"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(integerType)}
		prms["integerType"] = sliceVal
	}
	{
		sliceVal := []string{reg}
		prms["reg"] = sliceVal
	}
	{
		sliceVal := []string{stringType}
		prms["stringType"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ValidationTest"), rw, req, prms)
	validationCtx, _err := app.NewValidationValidationContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Validation(validationCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Validation
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Validation)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Validation", resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
