package handler

import (
	"app/controller"
	"app/core/logger"
	"app/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GetWarehouseList handles the list warehouse request and returns the list of warehouses as response
func GetWarehouseList(c echo.Context) error {
	response, err := controller.GetWarehouseList()
	if err != nil {
		logger.Error.Println("Error while getting warehouse list - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while getting warehouse list")
	}
	return c.JSON(http.StatusOK, response)
}

//CreateWarehouse handles the create warehouse request
func CreateWarehouse(c echo.Context) error {
	request := &model.Warehouse{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input data received")
	}
	if request.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Warehouse name not provided")
	}
	err := controller.CreateWarehouse(*request)
	if err != nil {
		logger.Error.Println("Error while creating warehouse - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while creating warehouse")
	}
	return c.String(http.StatusOK, "Warehouse successfully created")
}

//EditWarehouse handles the edit warehouse request
func EditWarehouse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid warehouse id")
	}
	request := &model.Warehouse{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input data received")
	}
	if request.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Warehouse name not provided")
	}
	request.ID = id
	rowsAffected, err := controller.EditWarehouse(*request)
	if err != nil {
		logger.Error.Println("Error while editing warehouse - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while editing warehouse")
	}
	if rowsAffected == 0 {
		logger.Error.Println("Error while editing warehouse - Warehouse not found")
		return echo.NewHTTPError(http.StatusBadRequest, "Error while editing warehouse - Warehouse not found")
	}
	return c.String(http.StatusOK, "Warehouse successfully edited")
}

//DeleteWarehouse handles the delete warehouse request
func DeleteWarehouse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid warehouse id")
	}
	rowsAffected, err := controller.DeleteWarehouse(id)
	if err != nil {
		logger.Error.Println("Error while deleting warehouse - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while deleting warehouse")
	}
	if rowsAffected == 0 {
		logger.Error.Println("Error while deleting warehouse - Warehouse not found")
		return echo.NewHTTPError(http.StatusBadRequest, "Error while deleting warehouse - Warehouse not found")
	}
	return c.String(http.StatusOK, "Warehouse successfully deleted")
}
