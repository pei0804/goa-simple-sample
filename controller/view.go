package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-simple-sample/app"
)

// ViewController implements the view resource.
type ViewController struct {
	*goa.Controller
}

// NewViewController creates a view controller.
func NewViewController(service *goa.Service) *ViewController {
	return &ViewController{Controller: service.NewController("ViewController")}
}

// View runs the view action.
func (c *ViewController) View(ctx *app.ViewViewContext) error {
	// ViewController_View: start_implement

	// Put your logic here

	// ViewController_View: end_implement
	res := &app.Viewtype{}
	return ctx.OK(res)
}
