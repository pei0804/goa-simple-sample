package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/goa-simple-sample/app"
)

// AccountsController implements the accounts resource.
type AccountsController struct {
	*goa.Controller
	db *gorm.DB
}

// NewAccountsController creates a accounts controller.
func NewAccountsController(service *goa.Service, db *gorm.DB) *AccountsController {
	return &AccountsController{
		Controller: service.NewController("AccountsController"),
		db:         db,
	}
}

// List runs the list action.
func (c *AccountsController) List(ctx *app.ListAccountsContext) error {
	// AccountsController_List: start_implement

	// Put your logic here

	// AccountsController_List: end_implement
	res := app.AccountCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AccountsController) Show(ctx *app.ShowAccountsContext) error {
	// AccountsController_Show: start_implement

	// Put your logic here

	// AccountsController_Show: end_implement
	res := &app.Account{}
	return ctx.OK(res)
}
