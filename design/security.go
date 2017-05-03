package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var UserAuth = APIKeySecurity("userToken", func() {
	Description("ユーザートークン")
	Header("X-Authorization")
})
