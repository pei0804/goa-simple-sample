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
	if ctx.ID == 0 {
		return ctx.NotFound()
	}

	// ActionsController_ID: end_implement
	res := &app.Integertype{}
	res.ID = ctx.ID
	return ctx.OK(res)
}

// Hello runs the hello action.
func (c *ActionsController) Hello(ctx *app.HelloActionsContext) error {
	// ActionsController_Ping: start_implement

	// Put your logic here
	name := ctx.Name

	// ActionsController_Ping: end_implement
	res := &app.Messagetype{}
	res.Message = "Hello " + name
	return ctx.OK(res)
}

// Ping runs the ping action.
func (c *ActionsController) Ping(ctx *app.PingActionsContext) error {
	// ActionsController_Ping: start_implement

	// Put your logic here
	message := "pong"

	// ActionsController_Ping: end_implement
	res := &app.Messagetype{}
	res.Message = message
	return ctx.OK(res)
}
