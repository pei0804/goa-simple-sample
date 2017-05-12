package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/goa-simple-sample/app"
	"github.com/tikasan/goa-simple-sample/models"
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

// Add runs the add action.
func (c *AccountsController) Add(ctx *app.AddAccountsContext) error {
	// AccountsController_Add: start_implement

	// Put your logic here
	a := &models.Account{}
	a.Name = ctx.Name
	a.Email = ctx.Email
	adb := models.NewAccountDB(c.db)
	err := adb.Add(ctx.Context, a)
	if err != nil {
		ctx.BadRequest(err)
	}

	// AccountsController_Add: end_implement
	res := &app.Account{}
	res = a.AccountToAccount()
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *AccountsController) Delete(ctx *app.DeleteAccountsContext) error {
	// AccountsController_Delete: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	err := adb.Delete(ctx.Context, ctx.ID)
	if err != nil {
		ctx.BadRequest(err)
	}

	// AccountsController_Delete: end_implement
	return ctx.OK([]byte{})
}

// List runs the list action.
func (c *AccountsController) List(ctx *app.ListAccountsContext) error {
	// AccountsController_List: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	a := adb.ListAccount(ctx.Context)

	// AccountsController_List: end_implement
	res := app.AccountCollection{}
	res = a
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AccountsController) Show(ctx *app.ShowAccountsContext) error {
	// AccountsController_Show: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	a, err := adb.OneAccount(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	// AccountsController_Show: end_implement
	res := &app.Account{}
	res = a
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountsController) Update(ctx *app.UpdateAccountsContext) error {
	// AccountsController_Update: start_implement

	// Put your logic here
	a := &models.Account{}
	a.ID = ctx.ID
	a.Name = ctx.Name
	a.Email = ctx.Email
	adb := models.NewAccountDB(c.db)
	err := adb.Update(ctx.Context, a)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ctx.BadRequest(err)
	}

	// AccountsController_Update: end_implement
	return ctx.OK([]byte{})
}
