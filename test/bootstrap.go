package test

import (
	"database/sql"
	"encoding/json"
	"os"
	"strconv"
	"runtime"

	. "gopkg.in/check.v1"

	"github.com/screwyprof/gosandbox/services/sandbox"
	SandboxServicePkg "github.com/screwyprof/gosandbox/services/sandbox"
	SandboxMysqlStoragePkg "github.com/screwyprof/gosandbox/storage/sandbox/mysql"
	"path"
)


type Config struct {
	DbHost string
	DbPort int
	DbName string
	DbUser string
	DBPass string
}

type IntegrationTestSuite struct {
	config *Config
	SandboxService sandbox.ISandboxService
}

func (suite *IntegrationTestSuite) SetUpSuite(c *C) {

	config, err := suite.initConfig()
	if err != nil {
		c.Fatal(err)
	}

	suite.config = config

	db, err := suite.initDb()
	if err != nil {
		c.Fatal(err)
	}

	suite.SandboxService = SandboxServicePkg.NewInstance(SandboxMysqlStoragePkg.NewInstance(db))
}

func (suite *IntegrationTestSuite) initDb() (*sql.DB, error) {
	db, err := sql.Open(
		"mysql", suite.config.DbUser + ":" + suite.config.DBPass +
		"@tcp(" + suite.config.DbHost + ":" + strconv.Itoa(suite.config.DbPort) + ")/" + suite.config.DbName)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func(suite *IntegrationTestSuite) initConfig() (*Config, error) {

	file, err := os.Open(getAbsolutePath("../application/config/config.json"))
	if (err != nil) {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	config := Config{}

	err = decoder.Decode(&config)
	if (err != nil) {
		return nil, err
	}

	return &config, nil
}

func getAbsolutePath(filename string) string {
	_, basename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(basename), filename)
}

//func (suite *IntegrationTestSuite) TearDownSuite(c *C) {
//
//}
