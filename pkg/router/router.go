package router

import (
	"pricingapi/pkg/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var loggerFormat = `[${time_rfc3339}]  ${status}  ${method} ${host}${path} ${latency_human}` + "\n"

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: loggerFormat}))

	e.POST("/apply", handler.Apply)

	return e
}
