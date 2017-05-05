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

var ViewType = MediaType("application/vnd.viewType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("value", String, "値", func() {
			Example("hoge")
		})
		Required("ID", "value")
	})
	View("default", func() {
		Attribute("ID")
		Attribute("value")
	})
	View("tiny", func() {
		Attribute("ID")
	})
})

var ArrayType = MediaType("application/vnd.arrayType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("value", String, "値", func() {
			Example("hoge")
		})
		Required("ID", "value")
	})
	View("default", func() {
		Attribute("ID")
		Attribute("value")
	})
})

var ValidationType = MediaType("application/vnd.validationType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("integer", Integer, "数字（1〜10）", func() {
			Example(5)
		})
		Attribute("string", String, "文字（1~10文字）", func() {
			Example("あいうえお")
		})
		Attribute("email", String, "メールアドレス", func() {
			Example("example@gmail.com")
		})
		Attribute("enum", String, "列挙型", func() {
			Example("A")
		})
		Attribute("default", String, "デフォルト値", func() {
			Example("でふぉ")
		})
	})
	Required("ID", "integer", "string", "email", "enum", "default")
	View("default", func() {
		Attribute("ID")
		Attribute("integer")
		Attribute("string")
		Attribute("email")
		Attribute("enum")
		Attribute("default")
	})
})
