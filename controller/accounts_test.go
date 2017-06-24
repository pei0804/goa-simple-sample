package controller

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	_ "github.com/lib/pq"
	"github.com/tikasan/goa-simple-sample/app/test"
	"github.com/tikasan/goa-simple-sample/models"
)

func TestAccountsController_Show(t *testing.T) {
	prepareTestDatabase()
	var (
		service = goa.New("goa simple sample")
		ctrl    = NewAccountsController(service, db)
	)

	cases := []models.Account{
		{
			ID:    1,
			Name:  "user1",
			Email: "example@gmail.com",
		},
		{
			ID:    2,
			Name:  "user2",
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
