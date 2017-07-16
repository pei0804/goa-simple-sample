package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/pei0804/goa-simple-sample/app"
	"github.com/pei0804/goa-simple-sample/models"
)

// BottlesDataController implements the bottles resource.
type BottlesDataController struct {
	*goa.Controller
	db *gorm.DB
}

// NewBottlesDataController creates a bottles controller.
func NewBottlesDataController(service *goa.Service, db *gorm.DB) *BottlesDataController {
	return &BottlesDataController{
		Controller: service.NewController("BottlesDataController"),
		db:         db,
	}
}

// Add runs the add action.
func (c *BottlesDataController) Add(ctx *app.AddBottlesDataContext) error {
	// BottlesDataController_Add: start_implement

	// Put your logic here
	b := &models.Bottle{}
	b.AccountID = ctx.AccountID
	b.Name = ctx.Name
	b.Quantity = ctx.Quantity
	bdb := models.NewBottleDB(c.db)
	err := bdb.Add(ctx.Context, b)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// BottlesDataController_Add: end_implement
	return ctx.Created()
}

// Delete runs the delete action.
func (c *BottlesDataController) Delete(ctx *app.DeleteBottlesDataContext) error {
	// BottlesDataController_Delete: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	err := bdb.Delete(ctx.Context, ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// BottlesDataController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *BottlesDataController) List(ctx *app.ListBottlesDataContext) error {
	// BottlesDataController_List: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	b := bdb.ListBottle(ctx.Context, 0)

	// BottlesDataController_List: end_implement
	res := &app.Bottlemedia{Status: 200, Data: b}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *BottlesDataController) Show(ctx *app.ShowBottlesDataContext) error {
	// BottlesDataController_Show: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	b, err := bdb.OneBottle(ctx.Context, ctx.ID, 0)
	if err != nil {
		return ctx.NotFound()
	}

	// BottlesDataController_Show: end_implement
	data := make([]*app.Bottle, 1)
	data[0] = b
	res := &app.Bottlemedia{Status: 200, Data: data}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *BottlesDataController) Update(ctx *app.UpdateBottlesDataContext) error {
	// BottlesDataController_Update: start_implement

	// Put your logic here
	b := &models.Bottle{}
	b.ID = ctx.ID
	b.Name = ctx.Name
	b.Quantity = ctx.Quantity
	bdb := models.NewBottleDB(c.db)
	err := bdb.Update(ctx.Context, b)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// BottlesDataController_Update: end_implement
	return nil
}
