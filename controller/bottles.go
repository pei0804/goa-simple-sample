package controller

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/goa-simple-sample/app"
	"github.com/tikasan/goa-simple-sample/models"
)

// BottlesController implements the bottles resource.
type BottlesController struct {
	*goa.Controller
	db *gorm.DB
}

// NewBottlesController creates a bottles controller.
func NewBottlesController(service *goa.Service, db *gorm.DB) *BottlesController {
	return &BottlesController{
		Controller: service.NewController("BottlesController"),
		db:         db,
	}
}

// Add runs the add action.
func (c *BottlesController) Add(ctx *app.AddBottlesContext) error {
	// BottlesController_Add: start_implement

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
	// BottlesController_Add: end_implement
	return ctx.Created()
}

// Delete runs the delete action.
func (c *BottlesController) Delete(ctx *app.DeleteBottlesContext) error {
	// BottlesController_Delete: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	err := bdb.Delete(ctx.Context, ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// BottlesController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *BottlesController) List(ctx *app.ListBottlesContext) error {
	// BottlesController_List: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	b := bdb.ListBottle(ctx.Context, 0)

	// BottlesController_List: end_implement
	res := app.BottleCollection{}
	res = b
	return ctx.OK(res)
}

// ListRelation runs the listRelation action.
func (c *BottlesController) ListRelation(ctx *app.ListRelationBottlesContext) error {
	// BottlesController_ListRelation: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	b := bdb.ListBottleFullScan(ctx.Context, 0)

	// BottlesController_ListRelation: end_implement
	res := app.BottleRelationCollection{}
	res = b
	return ctx.OKRelation(res)
}

// Show runs the show action.
func (c *BottlesController) Show(ctx *app.ShowBottlesContext) error {
	// BottlesController_Show: start_implement

	// Put your logic here
	bdb := models.NewBottleDB(c.db)
	b, err := bdb.OneBottle(ctx.Context, ctx.ID, 0)
	if err != nil {
		return ctx.NotFound()
	}

	// BottlesController_Show: end_implement
	res := &app.Bottle{}
	res = b
	return ctx.OK(res)
}

// Update runs the update action.
func (c *BottlesController) Update(ctx *app.UpdateBottlesContext) error {
	// BottlesController_Update: start_implement

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

	// BottlesController_Update: end_implement
	return nil
}
