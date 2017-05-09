package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
// MediaTypeに名前をつけます
var IntegerType = MediaType("application/vnd.integerType+json", func() {
	// 説明
	Description("example")
	// どのような値があるか（複数定義出来る）
	Attributes(func() {
		// IDはInteger型
		Attribute("ID", Integer, "ID", func() {
			// 返すレスポンスの例
			Example(1)
		})
		// レスポンスに必須な要素（基本は全て必須にした方が楽）
		Required("ID")
	})
	// 返すレスポンスのフォーマット（別記事で紹介予定）
	View("default", func() {
		Attribute("ID")
	})
})

var MessageType = MediaType("application/vnd.messageType+json", func() {
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

var UserType = MediaType("application/vnd.userType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("name", String, "名前", func() {
			Example("hoge")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("satak47cpc@gmail.com")
		})
		Required("ID", "name", "email")
	})
	// 特別な指定がない場合はdefaultのMediaType
	View("default", func() {
		Attribute("ID")
		Attribute("name")
		Attribute("email")
	})
	// tinyという名前の場合は、簡潔なレスポンスフォーマットにすることが出来る
	View("tiny", func() {
		Attribute("ID")
		Attribute("name")
	})
})

var ValidationType = MediaType("application/vnd.validationType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
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
	Required("ID", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
	View("default", func() {
		Attribute("ID")
		Attribute("integerType")
		Attribute("stringType")
		Attribute("email")
		Attribute("enumType")
		Attribute("defaultType")
		Attribute("reg")
	})
})

var ResponseFormat = Type("response", func() {
	Attribute("status", Integer, func() {
		Default(200)
	})
	Required("status")
})

var Data = Type("data", func() {
	Attribute("title", String)
	Attribute("body", String)
	Required("title", "body")
})

var ArticleType = MediaType("application/vnd.articleType+json", func() {
	Description("example")
	Attribute("data", ArrayOf(Data))
	Attribute("response", ResponseFormat)
	View("default", func() {
		Attribute("data")
		Attribute("response")
	})
	Required("data", "response")
})

//-------gorma--------

var Account = MediaType("application/vnd.account+json", func() {
	Description("celler account")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("name", String, "名前", func() {
			Example("山田　太郎")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("example@gmail.com")
		})
	})
	Required("ID", "name", "email")
	View("default", func() {
		Attribute("ID")
		Attribute("name")
		Attribute("email")
	})
})

var Bottle = MediaType("application/vnd.bottle+json", func() {
	Description("celler bottles")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("name", String, "ボトル名", func() {
			Example("シャルドネ")
		})
		Attribute("quantity", Integer, "数量", func() {
			Example(4)
		})
	})
	Required("ID", "name", "quantity")
	View("default", func() {
		Attribute("ID")
		Attribute("name")
		Attribute("quantity")
	})
})

var Category = MediaType("application/vnd.category+json", func() {
	Description("celler account")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("name", String, "名前", func() {
			Example("ワイン")
		})
	})
	Required("ID", "name")
	View("default", func() {
		Attribute("ID")
		Attribute("name")
	})
})
