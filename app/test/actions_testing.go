// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": actions TestHelpers
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.3.1

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
)

// HelloActionsBadRequest runs the method Hello of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HelloActionsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController, name string) (http.ResponseWriter, error) {
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
		sliceVal := []string{name}
		query["name"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/actions/hello"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{name}
		prms["name"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	helloCtx, _err := app.NewHelloActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Hello(helloCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// HelloActionsOK runs the method Hello of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HelloActionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController, name string) (http.ResponseWriter, *app.Message) {
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
		sliceVal := []string{name}
		query["name"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/v1/actions/hello"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{name}
		prms["name"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	helloCtx, _err := app.NewHelloActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Hello(helloCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Message", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// IDActionsBadRequest runs the method ID of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IDActionsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController, id int) (http.ResponseWriter, error) {
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
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/actions/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	idCtx, _err := app.NewIDActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.ID(idCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// IDActionsNotFound runs the method ID of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IDActionsNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController, id int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
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
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/actions/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	idCtx, _err := app.NewIDActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.ID(idCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// IDActionsOK runs the method ID of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func IDActionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController, id int) (http.ResponseWriter, *app.Integer) {
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
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/actions/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	idCtx, _err := app.NewIDActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.ID(idCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Integer
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Integer)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Integer", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// PingActionsBadRequest runs the method Ping of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PingActionsBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController) (http.ResponseWriter, error) {
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
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/actions/ping"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	pingCtx, _err := app.NewPingActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Ping(pingCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// PingActionsOK runs the method Ping of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func PingActionsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ActionsController) (http.ResponseWriter, *app.Message) {
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
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/actions/ping"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ActionsTest"), rw, req, prms)
	pingCtx, _err := app.NewPingActionsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Ping(pingCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Message
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.Message)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.Message", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
