package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-simple-sample/app"
)

// ResponseController implements the response resource.
type ResponseController struct {
	*goa.Controller
}

// NewResponseController creates a response controller.
func NewResponseController(service *goa.Service) *ResponseController {
	return &ResponseController{Controller: service.NewController("ResponseController")}
}

// Array runs the array action.
func (c *ResponseController) Array(ctx *app.ArrayResponseContext) error {
	// ResponseController_Array: start_implement

	// Put your logic here

	// ResponseController_Array: end_implement
	res := make([]int, 3)
	res[0] = 1
	res[1] = 2
	res[2] = 3
	return ctx.OK(res)
}

// Hash runs the hash action.
func (c *ResponseController) Hash(ctx *app.HashResponseContext) error {
	// ResponseController_Hash: start_implement

	// Put your logic here

	// ResponseController_Hash: end_implement
	res := make(map[string]int, 3)
	res["japan"] = 2
	res["america"] = 3
	res["korea"] = 4
	return ctx.OK(res)
}

// List runs the list action.
func (c *ResponseController) List(ctx *app.ListResponseContext) error {
	// ResponseController_List: start_implement

	// Put your logic here

	// ResponseController_List: end_implement
	res := make(app.UserTinyCollection, 2)
	u1 := &app.UserTiny{
		ID:   1,
		Name: "ユーザー1",
	}
	u2 := &app.UserTiny{
		ID:   2,
		Name: "ユーザー2",
	}
	res[0] = u1
	res[1] = u2
	return ctx.OKTiny(res)
}

// Show runs the show action.
func (c *ResponseController) Show(ctx *app.ShowResponseContext) error {
	// ResponseController_Show: start_implement

	// Put your logic here

	// ResponseController_Show: end_implement
	res := &app.User{}
	res.ID = 1
	res.Name = "ユーザー1"
	res.Email = "satak47cpc@gmail.com"
	return ctx.OK(res)
}

// Nested runs the show action.
func (c *ResponseController) Nested(ctx *app.NestedResponseContext) error {
	// ResponseController_Show: start_implement

	// Put your logic here

	//ResponseController_Show: end_implement
	res := &app.Article{}
	return ctx.OK(res)
}
