package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})

var _ = Resource("actions", func() {
	BasePath("/actions")
	Action("main", func() {
		Description("複数アクション（main）")
		Routing(
			GET("/main"),
		)
		Response(OK, MessageType)
		Response(BadRequest, ErrorMedia)
	})
	Action("sub", func() {
		Description("複数アクション（sub）")
		Routing(
			GET("/sub"),
		)
		Params(func() {
			Param("name", String, "名前", func() {
				Default("")
			})
			Required("name")
		})
		Response(OK, MessageType)
		Response(BadRequest, ErrorMedia)
	})
	Action("ID", func() {
		Description("複数アクション（:ID）")
		Routing(
			GET("/:ID"),
		)
		Params(func() {
			Param("ID", Integer, "ID")
		})
		Response(OK, IntegerType)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("security", func() {
	BasePath("/securiy")
	Action("security", func() {
		Description("セキュリティの例です")
		Routing(
			GET("/"),
		)
		Response(OK, MessageType)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("method", func() {
	BasePath("/method")
	Action("method", func() {
		Description("HTTPメソッド")
		Routing(
			GET("/get"),
			POST("/post"),
			DELETE("/delete"),
			PUT("/put"),
		)
		Response(OK, MessageType)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("view", func() {
	BasePath("/view")
	Action("view", func() {
		Description("MediaTypeのバリエーション")
		Routing(
			GET("/default"),
			GET("/tiny"),
		)
		Response(OK, ViewType)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("array", func() {
	BasePath("/array")
	Action("array", func() {
		Description("複数値")
		Routing(
			GET("/"),
		)
		Response(OK, CollectionOf(ArrayType))
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("validation", func() {
	BasePath("/validation")
	Action("validation", func() {
		Description("Validation")
		Routing(
			GET("/"),
		)
		Params(func() {
			Param("ID", Integer, "ID", func() {
				Example(1)
			})
			Param("integer", Integer, "数字（1〜10）", func() {
				Minimum(0)
				Maximum(10)
				Example(2)
			})
			Param("string", String, "文字（1~10文字）", func() {
				MinLength(1)
				MaxLength(10)
				Example("あいうえお")
			})
			Param("email", String, "メールアドレス", func() {
				Format("email")
				Example("example@gmail.com")
			})
			Param("enum", String, "列挙型", func() {
				Enum("A", "B", "C")
				Example("A")
			})
			Param("default", String, "デフォルト値", func() {
				Default("でふぉ")
				Example("でふぉ")
			})
		})
		Response(OK, ValidationType)
		Response(BadRequest, ErrorMedia)
	})
})
