package backend

import (
	"fmt"
	"log"

	"github.com/BlackMocca/sqlx"
	"github.com/Blackmocca/wedding-ui/backend/handler"
	"github.com/Blackmocca/wedding-ui/backend/middleware"
	"github.com/Blackmocca/wedding-ui/backend/repository"
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/labstack/echo/v4"
	echoMiddL "github.com/labstack/echo/v4/middleware"
	pg "github.com/lib/pq"
	"github.com/spf13/cast"
)

var (
	psqlConnectionStr = constants.GetEnv("PSQL_CONNECTION_URL", "")
	apiPort           = cast.ToInt(constants.GetEnv("API_PORT", "8081"))
)

func getPsqlClient(uri string) *sqlx.DB {
	addr, err := pg.ParseURL(uri)
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic(err)
	}

	return db
}

func StartServer() {
	middL := middleware.NewRestAPIMiddleware()
	client := getPsqlClient(psqlConnectionStr)
	defer client.Close()

	repo := repository.NewPsqlClient(client)
	h := handler.NewHandler(repo)

	e := echo.New()
	e.Use(echoMiddL.Recover())
	e.Use(echoMiddL.CORS())
	e.Use(middL.InitContext)
	e.Use(middL.InputForm)
	e.GET("/api/celebrate", h.Fetch)
	e.POST("/api/celebrate", h.Create)

	portStr := fmt.Sprintf(":%d", apiPort)
	log.Fatal(e.Start(portStr))
}
