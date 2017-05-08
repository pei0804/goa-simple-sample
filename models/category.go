// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": Models
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

// celler category
type Category struct {
	ID               int              `gorm:"primary_key"` // primary key
	BottleCategories []BottleCategory // has many BottleCategories
	Name             string
	CreatedAt        time.Time  // timestamp
	DeletedAt        *time.Time // nullable timestamp (soft delete)
	UpdatedAt        time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Category) TableName() string {
	return "categories"

}

// CategoryDB is the implementation of the storage interface for
// Category.
type CategoryDB struct {
	Db *gorm.DB
}

// NewCategoryDB creates a new storage type.
func NewCategoryDB(db *gorm.DB) *CategoryDB {
	return &CategoryDB{Db: db}
}

// DB returns the underlying database.
func (m *CategoryDB) DB() interface{} {
	return m.Db
}

// CategoryStorage represents the storage interface.
type CategoryStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Category, error)
	Get(ctx context.Context, id int) (*Category, error)
	Add(ctx context.Context, category *Category) error
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id int) error

	ListCategory(ctx context.Context) []*app.Category
	OneCategory(ctx context.Context, id int) (*app.Category, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *CategoryDB) TableName() string {
	return "categories"

}

// CRUD Functions

// Get returns a single Category as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *CategoryDB) Get(ctx context.Context, id int) (*Category, error) {
	defer goa.MeasureSince([]string{"goa", "db", "category", "get"}, time.Now())

	var native Category
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Category
func (m *CategoryDB) List(ctx context.Context) ([]*Category, error) {
	defer goa.MeasureSince([]string{"goa", "db", "category", "list"}, time.Now())

	var objs []*Category
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *CategoryDB) Add(ctx context.Context, model *Category) error {
	defer goa.MeasureSince([]string{"goa", "db", "category", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Category", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *CategoryDB) Update(ctx context.Context, model *Category) error {
	defer goa.MeasureSince([]string{"goa", "db", "category", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Category", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *CategoryDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "category", "delete"}, time.Now())

	var obj Category

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Category", "error", err.Error())
		return err
	}

	return nil
}
