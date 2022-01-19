package handler

import (
	"app/controller"
	"app/core/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

//GetWarehouseList handler
func GetWarehouseList(c echo.Context) error {
	response, err := controller.GetWarehouseList()
	if err != nil {
		logger.Error.Println("Error while getting warehouse list - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while getting warehouse list")
	}
	return c.JSON(http.StatusOK, response)
}
