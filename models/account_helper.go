// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": Model Helpers
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

// MediaType Retrieval Functions

// ListAccount returns an array of view: default.
func (m *AccountDB) ListAccount(ctx context.Context) []*app.Account {
	defer goa.MeasureSince([]string{"goa", "db", "account", "listaccount"}, time.Now())

	var native []*Account
	var objs []*app.Account
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Account", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountToAccount())
	}

	return objs
}

// AccountToAccount loads a Account and builds the default view of media type Account.
func (m *Account) AccountToAccount() *app.Account {
	account := &app.Account{}
	account.Email = m.Email
	account.ID = m.ID
	account.Name = m.Name

	return account
}

// OneAccount loads a Account and builds the default view of media type Account.
func (m *AccountDB) OneAccount(ctx context.Context, id int) (*app.Account, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "oneaccount"}, time.Now())

	var native Account
	err := m.Db.Scopes().Table(m.TableName()).Preload("Bottles").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Account", "error", err.Error())
		return nil, err
	}

	view := *native.AccountToAccount()
	return &view, err
}
