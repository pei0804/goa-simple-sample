package controller

import (
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/ory-am/dockertest.v3"
	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	db       *gorm.DB
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=celler_test"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = gorm.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/celler_test?parseTime=true&collation=utf8_general_ci&interpolateParams=true&loc=Asia%%2FTokyo", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.DB().Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}

	n, err := migrate.Exec(db.DB(), "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	testfixtures.SkipDatabaseNameCheck(true)
	fixtures, err = testfixtures.NewFiles(db.DB(), &testfixtures.MySQL{},
		"../fixtures/accounts.yml",
		"../fixtures/bottles.yml",
		"../fixtures/categories.yml",
		"../fixtures/bottle_categories.yml",
	)

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
