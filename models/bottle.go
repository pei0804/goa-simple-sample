// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": Models
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.1.0

package models

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/pei0804/goa-simple-sample/app"
	"time"
)

// celler bottle
type Bottle struct {
	ID               int              `gorm:"primary_key"` // primary key
	AccountID        int              // Belongs To Account
	BottleCategories []BottleCategory // has many BottleCategories
	Categories       []Category       `gorm:"many2many:bottle_categories"` // many to many Bottle/Category
	Name             string
	Quantity         int
	CreatedAt        time.Time  // timestamp
	DeletedAt        *time.Time // nullable timestamp (soft delete)
	UpdatedAt        time.Time  // timestamp
	Account          Account
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Bottle) TableName() string {
	return "bottles"

}

// BottleDB is the implementation of the storage interface for
// Bottle.
type BottleDB struct {
	Db *gorm.DB
}

// NewBottleDB creates a new storage type.
func NewBottleDB(db *gorm.DB) *BottleDB {
	return &BottleDB{Db: db}
}

// DB returns the underlying database.
func (m *BottleDB) DB() interface{} {
	return m.Db
}

// BottleStorage represents the storage interface.
type BottleStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Bottle, error)
	Get(ctx context.Context, id int) (*Bottle, error)
	Add(ctx context.Context, bottle *Bottle) error
	Update(ctx context.Context, bottle *Bottle) error
	Delete(ctx context.Context, id int) error

	ListBottle(ctx context.Context, accountID int) []*app.Bottle
	OneBottle(ctx context.Context, id int, accountID int) (*app.Bottle, error)

	ListBottleRelation(ctx context.Context, accountID int) []*app.BottleRelation
	OneBottleRelation(ctx context.Context, id int, accountID int) (*app.BottleRelation, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *BottleDB) TableName() string {
	return "bottles"

}

// Belongs To Relationships

// BottleFilterByAccount is a gorm filter for a Belongs To relationship.
func BottleFilterByAccount(accountID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if accountID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("account_id = ?", accountID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single Bottle as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *BottleDB) Get(ctx context.Context, id int) (*Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "get"}, time.Now())

	var native Bottle
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Bottle
func (m *BottleDB) List(ctx context.Context) ([]*Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "list"}, time.Now())

	var objs []*Bottle
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *BottleDB) Add(ctx context.Context, model *Bottle) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Bottle", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *BottleDB) Update(ctx context.Context, model *Bottle) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Bottle", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *BottleDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "delete"}, time.Now())

	var obj Bottle

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Bottle", "error", err.Error())
		return err
	}

	return nil
}
