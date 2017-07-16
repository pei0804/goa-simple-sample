package controller

import (
	"github.com/goadesign/goa"
	"github.com/pei0804/goa-simple-sample/app"
)

// ValidationController implements the validation resource.
type ValidationController struct {
	*goa.Controller
}

// NewValidationController creates a validation controller.
func NewValidationController(service *goa.Service) *ValidationController {
	return &ValidationController{Controller: service.NewController("ValidationController")}
}

// Validation runs the validation action.
func (c *ValidationController) Validation(ctx *app.ValidationValidationContext) error {
	// ValidationController_Validation: start_implement

	// Put your logic here

	// ValidationController_Validation: end_implement
	res := &app.Validation{}
	res.ID = ctx.ID
	res.IntegerType = ctx.IntegerType
	res.StringType = ctx.StringType
	res.Email = ctx.Email
	res.EnumType = ctx.EnumType
	res.DefaultType = ctx.DefaultType
	res.Reg = ctx.Reg
	return ctx.OK(res)
}
