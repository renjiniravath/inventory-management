package handler

import (
	"app/controller"
	"app/core/logger"
	"app/model"
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

func CreateWarehouse(c echo.Context) error {
	request := &model.Warehouse{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input data received")
	}
	if request.Name == "" {
		return c.JSON(http.StatusBadRequest, "Warehouse name not provided")
	}
	err := controller.CreateWarehouse(*request)
	if err != nil {
		logger.Error.Println("Error while getting warehouse list - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while getting warehouse list")
	}
	return c.String(http.StatusOK, "Warehouse successfully created")
}
