package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-simple-sample/app"
)

// ArrayController implements the array resource.
type ArrayController struct {
	*goa.Controller
}

// NewArrayController creates a array controller.
func NewArrayController(service *goa.Service) *ArrayController {
	return &ArrayController{Controller: service.NewController("ArrayController")}
}

// Array runs the array action.
func (c *ArrayController) Array(ctx *app.ArrayArrayContext) error {
	// ArrayController_Array: start_implement

	// Put your logic here

	// ArrayController_Array: end_implement
	res := app.ArraytypeCollection{}
	return ctx.OK(res)
}
