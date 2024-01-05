package handler

import (
	"net/http"
	"time"

	"github.com/Blackmocca/wedding-ui/backend/repository"
	"github.com/Blackmocca/wedding-ui/models"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

type Handler struct {
	repo *repository.PsqlClient
}

func NewHandler(repo *repository.PsqlClient) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Fetch(c echo.Context) error {
	var ctx = c.Request().Context()

	celebrates, err := h.repo.Fetch(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"celebrates": celebrates,
	}
	return echo.NewHTTPError(http.StatusOK, resp)
}

func (h *Handler) Create(c echo.Context) error {
	var ctx = c.Request().Context()
	var params = c.Get("params").(map[string]interface{})
	celebrate := models.NewCelebrate(cast.ToString(params["celebrate_text"]), cast.ToString(params["celebrate_from"]))
	celebrate.CreatedAt = time.Now()
	celebrate.UpdatedAt = time.Now()
	if err := h.repo.Create(ctx, celebrate); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := map[string]interface{}{
		"message": "Successful",
	}
	return echo.NewHTTPError(http.StatusOK, resp)
}
