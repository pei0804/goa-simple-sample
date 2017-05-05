package controller

import (
	"github.com/goadesign/goa"
)

// JsController implements the js resource.
type JsController struct {
	*goa.Controller
}

// NewJsController creates a js controller.
func NewJsController(service *goa.Service) *JsController {
	return &JsController{Controller: service.NewController("JsController")}
}
