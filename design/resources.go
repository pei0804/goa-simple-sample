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

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "/js/")
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
	Action("id", func() {
		Description("複数アクション（:id）")
		Routing(
			// エンドポイントにリソースを指定出来る
			// GET http://localhost:8080/api/v1/actions/1になる
			GET("/:id"),
		)
		Params(func() {
			// :idはIntegert型でなければならない。
			Param("id", Integer, "id")
			// Requiredはリソースを含めたエンドポイントになるので、定義しなくても良い
			//Required("id")
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
	Action("list", func() {
		Description("リストを返す")
		Routing(
			GET("/list"),
			GET("/list/new"),
			GET("/list/topic"),
		)
		Response(OK, CollectionOf(UserType))
		Response(BadRequest, ErrorMedia)
	})
	Action("follow", func() {
		Description("フォロー操作")
		Routing(
			PUT("/users/follow"),
			DELETE("/users/follow"),
		)
		Response(OK, MessageType)
		Response(BadRequest, ErrorMedia)
	})
	Action("etc", func() {
		Routing(GET("/users/:id/follow/:type"))
		Description("ちょっと特殊ケース")
		Params(func() {
			Param("id", Integer, "id")
			Param("type", Integer, "タイプ", func() {
				Enum(1, 2, 3)
			})
		})
		Response(OK, "plain/text")
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("response", func() {
	BasePath("/response")
	Action("list", func() {
		Description("ユーザー（複数）")
		Routing(
			GET("/users"),
		)
		// 複数返す
		Response(OK, CollectionOf(UserType))
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("ユーザー（単数）")
		Routing(
			GET("/users/:id"),
		)
		// 単一
		Response(OK, UserType)
		Response(BadRequest, ErrorMedia)
	})
	Action("hash", func() {
		Description("ユーザー（ハッシュ）")
		Routing(
			GET("/users/hash"),
		)
		// 連想配列
		Response(OK, HashOf(String, Integer))
		Response(BadRequest, ErrorMedia)
	})
	Action("array", func() {
		Description("ユーザー（配列）")
		Routing(
			GET("/users/array"),
		)
		// 配列
		Response(OK, ArrayOf(Integer))
		Response(BadRequest, ErrorMedia)
	})
	Action("nested", func() {
		Description("ネストしたMediaType")
		Routing(
			GET("/users/nested"),
		)
		Params(func() {
			Param("test", String, func() {
				MinLength(1)
			})
			Required("test")
		})
		// ネストした要素を返す
		Response(OK, ArticleType)
		Response(BadRequest, CustomeErrorMedia)
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
			// Integer型
			Param("id", Integer, "id", func() {
				Example(1)
			})
			// Integer型かつ1〜10以下
			Param("integerType", Integer, "数字（1〜10）", func() {
				Minimum(0)
				Maximum(10)
				Example(2)
			})
			// String型かつ1〜10文字以下
			Param("stringType", String, "文字（1~10文字）", func() {
				MinLength(1)
				MaxLength(10)
				Example("あいうえお")
			})
			// String型かつemailフォーマット
			Param("email", String, "メールアドレス", func() {
				Format("email")
				Example("example@gmail.com")
			})
			// String型でEnumで指定されたいずれかの文字列
			Param("enumType", String, "列挙型", func() {
				Enum("A", "B", "C")
				Example("A")
			})
			// String型で何も指定が無ければ”でふぉ”という文字列が自動でセットされる
			Param("defaultType", String, "デフォルト値", func() {
				Default("でふぉ")
				Example("でふぉ")
			})
			// String型で正規表現で指定したパターンの文字列
			Param("reg", String, "正規表現", func() {
				Pattern("^[a-z0-9]{5}$")
				Example("12abc")
			})

			// 全て必須パラメーター
			Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
		})
		Response(OK, ValidationType)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("accounts", func() {
	BasePath("/accounts")
	Action("list", func() {
		Description("複数")
		Routing(
			GET("/"),
		)
		Response(OK, CollectionOf(Account))
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("単数")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK, Account)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("追加")
		Routing(
			POST("/"),
		)
		Params(func() {
			Param("name", String, "名前", func() {
				Example("山田 太郎")
			})
			Param("email", String, "名前", func() {
				Format("email")
				Example("example@gmail.com")
			})
			Required("name", "email")
		})
		Response(OK, Account)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("削除")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "名前", func() {
				Example(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("更新")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("name", String, "名前", func() {
				Default("")
				Example("山田 太郎")
			})
			Param("email", String, "名前", func() {
				Format("email")
				Default("")
				Example("example@gmail.com")
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("bottles", func() {
	BasePath("/bottles")
	Action("list", func() {
		Description("複数")
		Routing(
			GET("/"),
		)
		Response(OK, CollectionOf(Bottle))
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("単数")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK, Bottle)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("追加")
		Routing(
			POST("/"),
		)
		Params(func() {
			Param("name", String, "ボトル名", func() {
				Default("")
				Example("赤ワインなにか")
			})
			Param("quantity", Integer, "数量", func() {
				Example(0)
			})
			Required("name", "quantity")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("削除")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "名前", func() {
				Example(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("更新")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("name", String, "ボトル名", func() {
				Default("")
				Example("赤ワインなにか")
			})
			Param("quantity", Integer, "数量", func() {
				Example(0)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
