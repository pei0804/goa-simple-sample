package controller

import (
	"github.com/goadesign/goa"
	"github.com/pei0804/goa-simple-sample/app"
)

// SecurityController implements the security resource.
type SecurityController struct {
	*goa.Controller
}

// NewSecurityController creates a security controller.
func NewSecurityController(service *goa.Service) *SecurityController {
	return &SecurityController{Controller: service.NewController("SecurityController")}
}

// Security runs the security action.
func (c *SecurityController) Security(ctx *app.SecuritySecurityContext) error {
	// SecurityController_Security: start_implement

	// Put your logic here

	// SecurityController_Security: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}
