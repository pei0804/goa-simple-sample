package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
var IntegerType = MediaType("application/vnd.integerType+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("ID", Integer, "数値", func() {
			Example(1)
		})
	})
	Required("ID")
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
	})
	Required("message")
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
	})
	Required("ID", "value")
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
	})
	Required("ID", "value")
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
