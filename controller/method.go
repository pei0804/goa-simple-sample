package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-simple-sample/app"
)

// MethodController implements the method resource.
type MethodController struct {
	*goa.Controller
}

// NewMethodController creates a method controller.
func NewMethodController(service *goa.Service) *MethodController {
	return &MethodController{Controller: service.NewController("MethodController")}
}

// Follow runs the follow action.
func (c *MethodController) Follow(ctx *app.FollowMethodContext) error {
	// MethodController_Follow: start_implement

	// Put your logic here

	// MethodController_Follow: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}

// List runs the list action.
func (c *MethodController) List(ctx *app.ListMethodContext) error {
	// MethodController_List: start_implement

	// Put your logic here

	// MethodController_List: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}

// Method runs the method action.
func (c *MethodController) Method(ctx *app.MethodMethodContext) error {
	// MethodController_Method: start_implement

	// Put your logic here

	// MethodController_Method: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}
