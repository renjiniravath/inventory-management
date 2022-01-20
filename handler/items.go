package handler

import (
	"app/controller"
	"app/core/logger"
	"app/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//GetInventoryList handles the inventory list endpoint and returns a list of inventories as response
func GetInventoryList(c echo.Context) error {
	response, err := controller.GetInventoryList()
	if err != nil {
		logger.Error.Println("Error while getting inventory list - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while getting inventory list")
	}
	return c.JSON(http.StatusOK, response)
}

//CreateItem handles the create item request
func CreateItem(c echo.Context) error {
	request := &model.Item{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input data received")
	}
	if request.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Item name not provided")
	}
	if request.WarehouseID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Warehouse id not provided")
	}
	err := controller.CreateItem(*request)
	if err != nil {
		logger.Error.Println("Error while creating item - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while creating item")
	}
	return c.String(http.StatusOK, "Item successfully created")
}

//EditItem handles the edit item request
func EditItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item id")
	}
	request := &model.Item{}
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input data received")
	}
	if request.Name == "" && request.WarehouseID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Provide either a new name or warehouse id")
	}
	request.ID = id
	rowsAffected, err := controller.EditItem(*request)
	if err != nil {
		logger.Error.Println("Error while editing item - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while editing item")
	}
	if rowsAffected == 0 {
		logger.Error.Println("Error while editing item - Item not found")
		return echo.NewHTTPError(http.StatusBadRequest, "Error while editing item - Item not found")
	}
	return c.String(http.StatusOK, "Item successfully edited")
}

//DeleteItem handles the delete item request
func DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item id")
	}
	rowsAffected, err := controller.DeleteItem(id)
	if err != nil {
		logger.Error.Println("Error while deleting item - " + err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Error while deleting item")
	}
	if rowsAffected == 0 {
		logger.Error.Println("Error while deleting item - Item not found")
		return echo.NewHTTPError(http.StatusBadRequest, "Error while deleting item - Item not found")
	}
	return c.String(http.StatusOK, "Item successfully deleted")
}
