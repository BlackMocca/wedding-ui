package backend

import (
	"fmt"
	"log"

	"github.com/Blackmocca/wedding-ui/backend/handler"
	"github.com/Blackmocca/wedding-ui/backend/middleware"
	"github.com/Blackmocca/wedding-ui/backend/repository"
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

var (
	psqlConnectionStr = constants.GetEnv("PSQL_CONNECTION_URL", "")
	apiPort           = cast.ToInt(constants.GetEnv("API_PORT", "8081"))
)

func StartServer() {
	middL := middleware.NewRestAPIMiddleware()

	repo := repository.NewPsqlClient(psqlConnectionStr)
	h := handler.NewHandler(repo)

	e := echo.New()
	e.Use(middL.InitContext)
	e.Use(middL.InputForm)
	e.POST("/celebrate", h.Create)

	portStr := fmt.Sprintf(":%d", apiPort)
	log.Fatal(e.Start(portStr))
}
