package controller

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-simple-sample/app"
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
	res := &app.Validationtype{}
	return ctx.OK(res)
}