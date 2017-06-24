package database

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	testfixtures "gopkg.in/testfixtures.v2"
	"gopkg.in/yaml.v1"
)

// Configsは環境ごとの設定情報をもつ
type Configs map[string]*Config

// Openは指定された環境についてDBに接続します。
func (cs Configs) Open(env string) (*gorm.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, nil
	}
	return config.Open()
}

// Configはsql-migrateの設定ファイルと同じ形式を想定している
type Config struct {
	Datasource string `yaml:"datasource"`
}

// DSNは設定されているDSNを返します
func (c *Config) DSN() string {
	return c.Datasource
}

// OpenはConfigで指定されている接続先に接続する。
// MySQL固定
func (c *Config) Open() (*gorm.DB, error) {
	return gorm.Open("mysql", c.DSN())
}

// NewConfigsFromFileはConfigから設定を読み取る
func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return NewConfigs(f)
}

// NewConfigsはio.ReaderからDB用設定を読み取る
func NewConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func PrepareTestDatabase() (*gorm.DB, *testfixtures.Context) {
	db, err := gorm.Open("mysql", "celler:celler@tcp(localhost:3306)/celler_test?parseTime=true&collation=utf8_general_ci&interpolateParams=true&loc=Asia%2FTokyo&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	// creating the context that hold the fixtures
	// see about all compatible databases in this page below
	testfixtures.SkipDatabaseNameCheck(true)
	fixtures, err := testfixtures.NewFiles(db.DB(), &testfixtures.MySQL{},
		"../fixtures/accounts.yml",
		"../fixtures/bottles.yml",
		"../fixtures/categories.yml",
		"../fixtures/bottle_categories.yml",
	)
	if err != nil {
		log.Fatal(err)
	}
	return db, fixtures
}

func PrepareTestData(fixtures *testfixtures.Context) {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
