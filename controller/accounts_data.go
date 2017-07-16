package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/pei0804/goa-simple-sample/app"
	"github.com/pei0804/goa-simple-sample/models"
)

// AccountsDataController implements the accounts resource.
type AccountsDataController struct {
	*goa.Controller
	db *gorm.DB
}

// NewAccountsDataController creates a accounts controller.
func NewAccountsDataController(service *goa.Service, db *gorm.DB) *AccountsDataController {
	return &AccountsDataController{
		Controller: service.NewController("AccountsDataController"),
		db:         db,
	}
}

// Add runs the add action.
func (c *AccountsDataController) Add(ctx *app.AddAccountsDataContext) error {
	// AccountsDataController_Add: start_implement

	// Put your logic here
	a := &models.Account{}
	a.Name = ctx.Name
	a.Email = ctx.Email
	adb := models.NewAccountDB(c.db)
	err := adb.Add(ctx.Context, a)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// AccountsDataController_Add: end_implement
	data := make([]*app.Account, 1)
	data[0] = a.AccountToAccount()
	res := &app.Accountmedia{Status: 200, Data: data}
	return ctx.OK(res)
}

// Delete runs the delete action.
func (c *AccountsDataController) Delete(ctx *app.DeleteAccountsDataContext) error {
	// AccountsDataController_Delete: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	err := adb.Delete(ctx.Context, ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// AccountsDataController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *AccountsDataController) List(ctx *app.ListAccountsDataContext) error {
	// AccountsDataController_List: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	a := adb.ListAccount(ctx.Context)

	// AccountsDataController_List: end_implement
	res := &app.Accountmedia{Status: 200, Data: a}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *AccountsDataController) Show(ctx *app.ShowAccountsDataContext) error {
	// AccountsDataController_Show: start_implement

	// Put your logic here
	adb := models.NewAccountDB(c.db)
	a, err := adb.OneAccount(ctx, ctx.ID)
	if err != nil {
		return ctx.NotFound()
	}

	// AccountsDataController_Show: end_implement
	data := make([]*app.Account, 1)
	data[0] = a
	res := &app.Accountmedia{Status: 200, Data: data}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountsDataController) Update(ctx *app.UpdateAccountsDataContext) error {
	// AccountsDataController_Update: start_implement

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
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// AccountsDataController_Update: end_implement
	return nil
}
