//go:generate goagen bootstrap -d github.com/tikasan/goa-simple-sample/design

package main

import (
	"flag"
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/goa-simple-sample/app"
	"github.com/tikasan/goa-simple-sample/controller"
	"github.com/tikasan/goa-simple-sample/database"
)

func main() {
	// Create service
	service := goa.New("goa simple sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	var (
		port = flag.String("port", ":8080", "addr to bind")
		env  = flag.String("env", "development", "application envirionment (production, development etc.)")
	)
	flag.Parse()

	cs, err := database.NewConfigsFromFile("dbconfig.yml")
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	dbcon, err := cs.Open(*env)
	if err != nil {
		log.Fatalf("database initialization failed: %s", err)
	}

	// Mount "accounts" controller
	c := controller.NewAccountsController(service, dbcon)
	app.MountAccountsController(service, c)
	// Mount "actions" controller
	c2 := controller.NewActionsController(service)
	app.MountActionsController(service, c2)
	// Mount "js" controller
	c3 := controller.NewJsController(service)
	app.MountJsController(service, c3)
	// Mount "method" controller
	c4 := controller.NewMethodController(service)
	app.MountMethodController(service, c4)
	// Mount "response" controller
	c5 := controller.NewResponseController(service)
	app.MountResponseController(service, c5)
	// Mount "security" controller
	c6 := controller.NewSecurityController(service)
	app.MountSecurityController(service, c6)
	// Mount "swagger" controller
	c7 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c7)
	// Mount "validation" controller
	c8 := controller.NewValidationController(service)
	app.MountValidationController(service, c8)

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}
