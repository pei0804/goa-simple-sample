package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Swaggerをローカルで実行するめの定義
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})

// /actionsの定義をする
var _ = Resource("actions", func() {
	// actionsリソースのベースパス
	BasePath("/actions")
	/*
		リソースに対してどういった操作を行うか定義する
		add	リソースを追加する
		list	リソースをリストで取得する
		delete	リソースを削除する
		上記のような感じで定義すればおｋです。
	*/
	Action("ping", func() {
		// アクションの説明
		Description("サーバーとの導通確認")
		Routing(
			// エンドポイント -> GET http://localhost/api/v1/actions/pingになる
			GET("/ping"),
		)
		// 返したいレスポンス
		// 200 OK + MessageTypeで定義しているMediaType
		Response(OK, MessageType)
		// 400 BadRequest + ErrorMediaというデフォルトで容易されているMediaType
		// 足りないパラメーターなどがあれば自動的に返される
		Response(BadRequest, ErrorMedia)
	})
	Action("hello", func() {
		Description("挨拶する")
		Routing(
			GET("/hello"),
		)
		// リクエストで付加出来るパラメーター
		Params(func() {
			// nameという名前でパラメーターをStringでなげれる
			Param("name", String, "名前", func() {
				// もし、空だった場合は空文字を格納する
				Default("")
			})
			// 必ず設定されるべきパラメーター（デフォルト値があるので、存在しなければ空になる）
			Required("name")
		})
		Response(OK, MessageType)
		Response(BadRequest, ErrorMedia)
	})
	Action("ID", func() {
		Description("複数アクション（:ID）")
		Routing(
			// エンドポイントにリソースを指定出来る
			// GET http://localhost:8080/api/v1/actions/1になる
			GET("/:ID"),
		)
		Params(func() {
			// :IDはIntegert型でなければならない。
			Param("ID", Integer, "ID")
			// Requiredはリソースを含めたエンドポイントになるので、定義しなくても良い
			//Required("ID")
		})
		Response(OK, IntegerType)
		// 指定したリソースが無ければNotFoundを返す可能生がある
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
