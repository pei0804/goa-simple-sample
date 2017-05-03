package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-simple-sample/app"
)

// ActionsController implements the actions resource.
type ActionsController struct {
	*goa.Controller
}

// NewActionsController creates a actions controller.
func NewActionsController(service *goa.Service) *ActionsController {
	return &ActionsController{Controller: service.NewController("ActionsController")}
}

// ID runs the id action.
func (c *ActionsController) ID(ctx *app.IDActionsContext) error {
	// ActionsController_ID: start_implement

	// Put your logic here

	// ActionsController_ID: end_implement
	res := &app.Integertype{}
	return ctx.OK(res)
}

// Main runs the main action.
func (c *ActionsController) Main(ctx *app.MainActionsContext) error {
	// ActionsController_Main: start_implement

	// Put your logic here

	// ActionsController_Main: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}

// Sub runs the sub action.
func (c *ActionsController) Sub(ctx *app.SubActionsContext) error {
	// ActionsController_Sub: start_implement

	// Put your logic here

	// ActionsController_Sub: end_implement
	res := &app.Messagetype{}
	return ctx.OK(res)
}
