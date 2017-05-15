package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
// MediaTypeに名前をつけます
var IntegerMedia = MediaType("application/vnd.integer+json", func() {
	// 説明
	Description("example")
	// どのような値があるか（複数定義出来る）
	Attributes(func() {
		// idはInteger型
		Attribute("id", Integer, "id", func() {
			// 返すレスポンスの例
			Example(1)
		})
		// レスポンスに必須な要素（基本は全て必須にした方が楽）
		Required("id")
	})
	// 返すレスポンスのフォーマット（別記事で紹介予定）
	View("default", func() {
		Attribute("id")
	})
})

var MessageMedia = MediaType("application/vnd.message+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("message", String, "メッセージ", func() {
			Example("ok")
		})
		Required("message")
	})
	View("default", func() {
		Attribute("message")
	})
})

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("name", String, "名前", func() {
			Example("hoge")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("satak47cpc@gmail.com")
		})
		Required("id", "name", "email")
	})
	// 特別な指定がない場合はdefaultのMediaType
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
	// tinyという名前の場合は、簡潔なレスポンスフォーマットにすることが出来る
	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})
})

var ValidationMedia = MediaType("application/vnd.validation+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("integerType", Integer, "数字（1〜10）", func() {
			Example(5)
		})
		Attribute("stringType", String, "文字（1~10文字）", func() {
			Example("あいうえお")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("example@gmail.com")
		})
		Attribute("enumType", String, "列挙型", func() {
			Example("A")
		})
		Attribute("defaultType", String, "デフォルト値", func() {
			Example("でふぉ")
		})
		Attribute("reg", String, "デフォルト値", func() {
			Example("12abc")
		})
	})
	Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
	View("default", func() {
		Attribute("id")
		Attribute("integerType")
		Attribute("stringType")
		Attribute("email")
		Attribute("enumType")
		Attribute("defaultType")
		Attribute("reg")
	})
})

var OKMedia = Type("ok", func() {
	Attribute("status", Integer, func() {
		Example(200)
	})
	Required("status")
})

var ErrorValue = Type("errorValue", func() {
	Attribute("status", Integer, func() {
		Example(400)
	})
	Attribute("code", String, func() {
		Example("bad_request")
	})
	Attribute("detail", String, func() {
		Example("xxx is nil")
	})
	Required("status", "code", "detail")
})

var CustomeErrorMedia = MediaType("application/vnd.error+json", func() {
	Attribute("response", ErrorValue)
	Required("response")
	View("default", func() {
		Attribute("response")
	})
})

var ArticleData = Type("data", func() {
	Attribute("title", String)
	Attribute("body", String)
	Required("title", "body")
})

var ArticleMedia = MediaType("application/vnd.article+json", func() {
	Description("example")
	Attribute("data", ArrayOf(ArticleData))
	Attribute("response", OKMedia)
	Required("data", "response")
	View("default", func() {
		Attribute("data")
		Attribute("response")
	})
})

//--------------------------------
// gorma
//--------------------------------

var AccountData = MediaType("application/vnd.account+json", func() {
	Description("celler account")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "名前", func() {
		Example("山田　太郎")
	})
	Attribute("email", String, "メールアドレス", func() {
		Example("example@gmail.com")
	})
	Required("id", "name", "email")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
})

var AccountMedia = MediaType("application/vnd.accountMedia+json", func() {
	Description("example")
	Attribute("data", ArrayOf(AccountData))
	Attribute("status", Integer, func() {
		Example(200)
	})
	Required("data", "status")
	View("default", func() {
		Attribute("data")
		Attribute("status")
	})
})

// Bottle

var BottleData = MediaType("application/vnd.bottle+json", func() {
	Description("celler bottles")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "ボトル名", func() {
		Example("シャルドネ")
	})
	Attribute("quantity", Integer, "数量", func() {
		Example(4)
	})
	// accountのMediaTypeを入れ子構造にする
	Attribute("account", AccountData)
	Attribute("categories", ArrayOf(CategoryData))
	Required("id", "name", "quantity", "account", "categories")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("quantity")
		Attribute("account")
		Attribute("categories")
	})
})

var BottleMedia = MediaType("application/vnd.bottleMedia+json", func() {
	Description("example")
	Attribute("data", ArrayOf(BottleData))
	Attribute("status", Integer, func() {
		Example(200)
	})
	Required("data", "status")
	View("default", func() {
		Attribute("data")
		Attribute("status")
	})
})

// Category

var CategoryData = MediaType("application/vnd.category+json", func() {
	Description("celler account")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "名前", func() {
		Example("ワイン")
	})
	Required("id", "name")
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

var CategoryMedia = MediaType("application/vnd.categoryMedia+json", func() {
	Description("example")
	Attribute("data", ArrayOf(CategoryData))
	Attribute("status", Integer, func() {
		Example(200)
	})
	Required("data", "status")
	View("default", func() {
		Attribute("data")
		Attribute("status")
	})
})
