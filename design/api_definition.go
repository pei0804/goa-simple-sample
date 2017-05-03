package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("tikasan/goa-simple-sample", func() {
	Title("tikasan/goa-simple-sample")
	Description("tikasan/goa-simple-sample")
	License(func() {
		Name("MIT")
		URL("https://github.com/tikasan/eventory/blob/master/LICENSE")
	})
	Docs(func() {
		Description("tikasan/goa-simple-sample")
		URL("https://github.com/tikasan/goa-simple-sample")
	})
	Host("localhost:8080")
	Scheme("http")
	Scheme("https")
	BasePath("/")

	Origin("*", func() {
		Methods("GET", "POST", "PUT", "DELETE")
		MaxAge(600)
		Credentials()
	})
})
