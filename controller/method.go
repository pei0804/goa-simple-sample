package controller

import (
	"fmt"

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

// Etc runs the etc action.
func (c *MethodController) Etc(ctx *app.EtcMethodContext) error {
	// MethodController_Etc: start_implement

	// Put your logic here

	// MethodController_Etc: end_implement
	return ctx.OK([]byte(fmt.Sprintf("ID: %d, Type %d", ctx.ID, ctx.Type)))
}

// Follow runs the follow action.
func (c *MethodController) Follow(ctx *app.FollowMethodContext) error {
	// MethodController_Follow: start_implement

	// Put your logic here
	var message string
	//　何かのフォロー操作
	if "PUT" == ctx.Request.Method {
		message = "フォローした"
	} else if "DELETE" == ctx.Request.Method {
		message = "フォロー外した"
	}
	// MethodController_Follow: end_implement
	res := &app.Message{}
	res.Message = message
	return ctx.OK(res)
}

// List runs the list action.
func (c *MethodController) List(ctx *app.ListMethodContext) error {
	// MethodController_List: start_implement

	// Put your logic here
	var listType string
	switch ctx.RequestURI {
	//case client.ListMethodPath(): <---これでもいけるけど・・・って感じ　何かいい方法。
	case "/api/v1/method/list":
		listType = "ただの"
	case "/api/v1/method/list/new":
		listType = "新しい"
	case "/api/v1/method/list/topic":
		listType = "注目されている"
	}
	// MethodController_List: end_implement
	res := make(app.UserTinyCollection, 2)
	u1 := &app.UserTiny{
		ID:   1,
		Name: listType + "ユーザー1",
	}
	u2 := &app.UserTiny{
		ID:   2,
		Name: listType + "ユーザー2",
	}
	res[0] = u1
	res[1] = u2
	return ctx.OKTiny(res)
}

// Method runs the method action.
func (c *MethodController) Method(ctx *app.MethodMethodContext) error {
	// MethodController_Method: start_implement

	// Put your logic here
	message := ctx.RequestURI

	// MethodController_Method: end_implement
	res := &app.Message{}
	res.Message = message
	return ctx.OK(res)
}
