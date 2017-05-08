package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("celler", func() {
	Description("celler Model")
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQLのリレーションナルデータベース")
		Model("Account", func() {
			RendersTo(Account)
			Description("celler account")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("email", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			HasMany("Bottles", "Bottle")
		})
		Model("Bottle", func() {
			RendersTo(Bottle)
			Description("celler bottle")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			BelongsTo("Account")
			HasMany("BottleCategories", "BottleCategory")
		})
		Model("Category", func() {
			RendersTo(Category)
			Description("celler category")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			HasMany("BottleCategories", "BottleCategory")
		})
		Model("BottleCategory", func() {
			Description("celler bottle category")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			BelongsTo("Category")
			BelongsTo("Bottle")
		})
	})
})
