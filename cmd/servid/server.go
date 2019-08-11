package servid

import (
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/rafaelcam/go-transaction-service/internal/platform/db"
	"github.com/rafaelcam/go-transaction-service/internal/transaction"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

const (
	configFilePath = "./configs"
	envUsage       = "environment for app, prod, dev, test"
	envDefault     = "dev"
	envFlagName    = "env"
)

var env string

// App Instance which contains router and dao
type App struct {
	*http.Server
	e                 *echo.Echo
	db                *sqlx.DB
	transactionRouter *transaction.Router
}

func config() {
	configLogger()
	configuration(configFilePath, env)
}

func configLogger() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func configuration(path string, env string) {
	flag.StringVar(&env, envFlagName, envDefault, envUsage)
	flag.Parse()

	if flag.Lookup("test.v") != nil {
		env = "test"
		path = "./../../configs"
	}

	log.Println("Environment is: " + env + " configFilePath is: " + path)

	viper.SetConfigName("conf_" + env)
	viper.AddConfigPath(path) // working directory

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}

}

// NewApp creates new App with db connection pool
func NewApp() *App {
	config()

	e := echo.New()
	e.Use(middleware.Logger()) // Log Requests

	database := setupDB(viper.GetString("database.url"))
	transactionRouter := transaction.NewRouter(e, database)

	server := &App{
		e:                 e,
		db:                database,
		transactionRouter: transactionRouter,
	}

	server.routes()

	return server
}

// Start launching the server
func (a *App) Start() {
	port := viper.GetString("server.port")
	a.e.Logger.Fatal(a.e.Start(port))
}

func (a *App) routes() {
	a.transactionRouter.Routes()
	//showRoutes(a.r)
}

func setupDB(dbURL string) *sqlx.DB {
	postgres, err := db.New(dbURL)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	return postgres
}
