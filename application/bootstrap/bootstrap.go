package bootstrap

import (
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"strconv"

	SandboxLoadUserNameByIdHandlerPkg "github.com/screwyprof/gosandbox/handlers/sandbox"
	SandboxServicePkg "github.com/screwyprof/gosandbox/services/sandbox"
	SandboxMysqlStoragePkg "github.com/screwyprof/gosandbox/storage/sandbox/mysql"
)

type Config struct {
	DbHost string
	DbPort int
	DbName string
	DbUser string
	DBPass string
}

type appContext struct {
	db *sql.DB
	config *Config
}

func Bootstrap() *appContext {
	log.Println("Bootstrapping application")

	app := &appContext{}

	app.initConfig()
	app.initDb()

	return app
}

func (app *appContext) Run() {
	sandboxService := SandboxServicePkg.NewInstance(SandboxMysqlStoragePkg.NewInstance(app.db))

	http.Handle(
		"/load-user-name-by-id",
		SandboxLoadUserNameByIdHandlerPkg.NewSandboxLoadUserNameByIdHandler(sandboxService))

	http.ListenAndServe("localhost:3000", nil)
}

func(app *appContext) initConfig() {

	file, err := os.Open("config/config.json")
	if (err != nil) {
		panic("Cannot open config.json!")
	}

	decoder := json.NewDecoder(file)
	config := Config{}

	err = decoder.Decode(&config)
	if (err != nil) {
		panic("Cannot read config.json: " + err.Error())
	}

	app.config = &config
}

func (app *appContext) initDb() {
	db, err := sql.Open(
		"mysql", app.config.DbUser + ":" + app.config.DBPass +
		"@tcp(" + app.config.DbHost + ":" + strconv.Itoa(app.config.DbPort) + ")/" + app.config.DbName)

	if err != nil {
		panic("Cannot connect to mysql" + err.Error())
	}
	//defer db.Close()
	app.db = db
}
