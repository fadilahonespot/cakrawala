package handler

import (
	"net/http"
	"time"

	"github.com/fadilahonespot/cakrawala/utils/response"
	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {}

func NewHealthCheckHandler() *HealthCheckHandler {
    return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) ServeHTTP(c echo.Context) (err error) {
	data := map[string]interface{}{
		"time": time.Now().UTC().Format(time.RFC1123),
	}
	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}