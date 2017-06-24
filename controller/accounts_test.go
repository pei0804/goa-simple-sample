package controller

import (
	"context"
	"os"
	"testing"

	testfixtures "gopkg.in/testfixtures.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/tikasan/goa-simple-sample/app/test"
	"github.com/tikasan/goa-simple-sample/controller"
	"github.com/tikasan/goa-simple-sample/database"
	"github.com/tikasan/goa-simple-sample/models"
)

var (
	db       *gorm.DB
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	db, fixtures = database.PrepareTestDatabase()
	os.Exit(m.Run())
}

func TestAccountsController_Show(t *testing.T) {
	database.PrepareTestData(fixtures)
	var (
		service = goa.New("goa simple sample")
		ctrl    = controller.NewAccountsController(service, db)
	)

	cases := []models.Account{
		{
			ID:    1,
			Name:  "ユーザー1",
			Email: "example@gmail.com",
		},
		{
			ID:    2,
			Name:  "ユーザー2",
			Email: "example2@gmail.com",
		},
	}

	for k, tc := range cases {
		_, account := test.ShowAccountsOK(t, context.Background(), service, ctrl, tc.ID)

		if account == nil {
			t.Fatalf("%s: nil account", k)
		}
		if account.ID != tc.ID {
			t.Errorf("%v: invalid account ID, expected 1, got %v", k, tc.ID)
		}
		if account.Name != tc.Name {
			t.Errorf("%v: invalid account Name, expected 1, got %v", k, tc.Name)
		}
		if account.Email != tc.Email {
			t.Errorf("%v: invalid account Email, expected 1, got %v", k, tc.Email)
		}
	}
}
