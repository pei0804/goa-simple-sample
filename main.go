//go:generate goagen bootstrap -d github.com/tikasan/goa-simple-sample/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/goa-simple-sample/app"
	"github.com/tikasan/goa-simple-sample/controller"
)

func main() {
	// Create service
	service := goa.New("goa simple sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "actions" controller
	c := controller.NewActionsController(service)
	app.MountActionsController(service, c)
	// Mount "array" controller
	c2 := controller.NewArrayController(service)
	app.MountArrayController(service, c2)
	// Mount "js" controller
	c3 := controller.NewJsController(service)
	app.MountJsController(service, c3)
	// Mount "method" controller
	c4 := controller.NewMethodController(service)
	app.MountMethodController(service, c4)
	// Mount "security" controller
	c5 := controller.NewSecurityController(service)
	app.MountSecurityController(service, c5)
	// Mount "swagger" controller
	c6 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c6)
	// Mount "validation" controller
	c7 := controller.NewValidationController(service)
	app.MountValidationController(service, c7)
	// Mount "view" controller
	c8 := controller.NewViewController(service)
	app.MountViewController(service, c8)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
