package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"pricingapi/pkg/api"
	"pricingapi/pkg/request"
	"pricingapi/pkg/response"

	"github.com/labstack/echo/v4"
)

func Apply(c echo.Context) error {
	defer c.Request().Body.Close()

	req := request.Apply{}
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		log.Printf("Failed processing apply request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	appliedRules := api.Instance().Apply(req.Price, req.UserType)

	return c.JSON(http.StatusOK, &response.Apply{
		Applied:      len(appliedRules) != 0,
		AppliedRules: appliedRules,
	})
}
