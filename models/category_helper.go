// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": Model Helpers
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/goa-simple-sample/app"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListCategory returns an array of view: default.
func (m *CategoryDB) ListCategory(ctx context.Context) []*app.Category {
	defer goa.MeasureSince([]string{"goa", "db", "category", "listcategory"}, time.Now())

	var native []*Category
	var objs []*app.Category
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Category", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.CategoryToCategory())
	}

	return objs
}

// CategoryToCategory loads a Category and builds the default view of media type Category.
func (m *Category) CategoryToCategory() *app.Category {
	category := &app.Category{}
	category.Name = m.Name

	return category
}

// OneCategory loads a Category and builds the default view of media type Category.
func (m *CategoryDB) OneCategory(ctx context.Context, id int) (*app.Category, error) {
	defer goa.MeasureSince([]string{"goa", "db", "category", "onecategory"}, time.Now())

	var native Category
	err := m.Db.Scopes().Table(m.TableName()).Preload("BottleCategories").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Category", "error", err.Error())
		return nil, err
	}

	view := *native.CategoryToCategory()
	return &view, err
}
