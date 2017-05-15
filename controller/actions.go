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

// Action ID
func (c *ActionsController) ID(ctx *app.IDActionsContext) error {

	/*
		Paramsで定義したものはctx.hogeで受け取れる
		Params(func() {
			Param("ID", Integer, "ID")
		})
	*/
	if ctx.ID == 0 {
		// Response(NotFound)
		// 0の場合は、404エラーを返す
		return ctx.NotFound()
	}

	// Response(OK, IntegerType)
	// MediaType IntegerTypeでレスポンス
	res := &app.Integer{}
	/*
		IDはInteger型で
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
	*/
	res.ID = ctx.ID
	return ctx.OK(res)
}

// Hello runs the hello action.
func (c *ActionsController) Hello(ctx *app.HelloActionsContext) error {
	// ActionsController_Ping: start_implement

	// Put your logic here
	name := ctx.Name

	// ActionsController_Ping: end_implement
	res := &app.Message{}
	res.Message = "Hello " + name
	return ctx.OK(res)
}

// Ping runs the ping action.
func (c *ActionsController) Ping(ctx *app.PingActionsContext) error {
	// ActionsController_Ping: start_implement

	// Put your logic here
	message := "pong"

	// ActionsController_Ping: end_implement
	res := &app.Message{}
	res.Message = message
	return ctx.OK(res)
}
