package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/pei0804/goa-simple-sample/app"
)

// ListBottle returns an array of view: default.
func (m *BottleDB) ListBottleFullScan(ctx context.Context, accountID int) []*app.BottleRelation {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottle"}, time.Now())

	var native []*Bottle
	var objs []*app.BottleRelation
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("BottleCategories").Preload("Account").Find(&native).Error

	for k, v := range native {
		for k2, v2 := range v.BottleCategories {
			native2 := []*BottleCategory{}
			m.Db.Preload("Category").Where("bottle_id = ?", v2.BottleID).Find(&native2)
			native[k].Categories = append(native[k].Categories, native2[k2].Category)
		}
	}

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottleRelation())
	}

	return objs
}
