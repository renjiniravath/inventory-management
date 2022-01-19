package handler

import (
	"app/controller"
	"app/core/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

//GetInventoryList handler
func GetInventoryList(c echo.Context) error {
	response, err := controller.GetInventoryList()
	if err != nil {
		logger.Error.Println("Error while getting inventory list - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while getting inventory list")
	}
	return c.JSON(http.StatusOK, response)
}
