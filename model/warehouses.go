package model

import (
	"app/container"
)

//Item holds the details of an item
type Warehouse struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

//WarehouseListResponse holds the response of a warehouse list get request
type WarehouseListResponse struct {
	Warehouses []Warehouse `json:"warehouse"`
	Count      int         `json:"count"`
}

//ReadWarehouseList queries warehouse list from db
func ReadWarehouseList() (WarehouseListResponse, error) {
	dbr := container.GetReader()
	warehouses := make([]Warehouse, 0)
	err := dbr.Select(&warehouses, "SELECT id, name FROM warehouses")
	if err != nil {
		return WarehouseListResponse{}, err
	}
	warehouseListResponse := WarehouseListResponse{
		Warehouses: warehouses,
		Count:      len(warehouses),
	}
	return warehouseListResponse, nil
}

//CreateWarehouse creates a new warehouse entry in db
func CreateWarehouse(warehouse Warehouse) error {
	dbw := container.GetWriter()
	_, err := dbw.Exec("INSERT INTO warehouses (name) VALUES (?)", warehouse.Name)
	if err != nil {
		return err
	}
	return nil
}
