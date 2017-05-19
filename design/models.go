package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("celler", func() {
	Description("celler Model")
	// Mysqlを使う
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQLのリレーションナルデータベース")
		// accountsテーブルのModelなら、Account
		Model("Account", func() {
			// MediaTypeで作成したAccountにマッピングする
			RendersTo(AccountData)
			Description("celler account")
			// PrimaryKeyの設定
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("email", gorma.String)
			// timestamp系の定義
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			// HasMany(複数形名, 単数形名)
			HasMany("Bottles", "Bottle")
		})
		Model("Bottle", func() {
			RendersTo(BottleData)
			Description("celler bottle")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("quantity", gorma.Integer)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			// BelongTo(単数形)
			BelongsTo("Account")
			HasMany("BottleCategories", "BottleCategory")
			ManyToMany("Category", "bottle_categories")
		})
		Model("Category", func() {
			RendersTo(CategoryData)
			Description("celler category")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			HasMany("BottleCategories", "BottleCategory")
			ManyToMany("Bottle", "Bottles")
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
