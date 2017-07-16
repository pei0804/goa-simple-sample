//go:generate goagen bootstrap -d github.com/pei0804/goa-simple-sample/design

package main

import (
	"flag"
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/pei0804/goa-simple-sample/app"
	"github.com/pei0804/goa-simple-sample/controller"
	"github.com/pei0804/goa-simple-sample/database"
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
		port  = flag.String("port", ":8080", "addr to bind")
		env   = flag.String("env", "development", "application envirionment (production, development etc.)")
		dbrun = flag.Bool("dbrun", false, "database run mode")
	)
	flag.Parse()

	if *dbrun {
		cs, err := database.NewConfigsFromFile("dbconfig.yml")
		if err != nil {
			log.Fatalf("cannot open database configuration. exit. %s", err)
		}
		dbcon, err := cs.Open(*env)
		if err != nil {
			log.Fatalf("database initialization failed: %s", err)
		}
		// Mount "accounts" controller
		a := controller.NewAccountsController(service, dbcon)
		app.MountAccountsController(service, a)
		// Mount "bottles" controller
		b := controller.NewBottlesController(service, dbcon)
		app.MountBottlesController(service, b)
		// Mount "accounts_data" controller
		ad := controller.NewAccountsDataController(service, dbcon)
		app.MountAccountsDataController(service, ad)
		// Mount "bottles_data" controller
		bd := controller.NewBottlesDataController(service, dbcon)
		app.MountBottlesDataController(service, bd)
	}

	// Mount "actions" controller
	action := controller.NewActionsController(service)
	app.MountActionsController(service, action)
	// Mount "js" controller
	js := controller.NewJsController(service)
	app.MountJsController(service, js)
	// Mount "method" controller
	method := controller.NewMethodController(service)
	app.MountMethodController(service, method)
	// Mount "response" controller
	response := controller.NewResponseController(service)
	app.MountResponseController(service, response)
	// Mount "security" controller
	security := controller.NewSecurityController(service)
	app.MountSecurityController(service, security)
	// Mount "swagger" controller
	swagger := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, swagger)
	// Mount "validation" controller
	validation := controller.NewValidationController(service)
	app.MountValidationController(service, validation)

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}
}
