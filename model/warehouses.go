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
	Warehouses []Warehouse `json:"warehouses"`
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

//EditWarehouse updates the name of a warehouse in db
func EditWarehouse(warehouse Warehouse) (int64, error) {
	dbw := container.GetWriter()
	result, err := dbw.Exec("UPDATE warehouses SET name = ? WHERE id = ?", warehouse.Name, warehouse.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

//DeleteWarehouse deletes a warehouse from db
func DeleteWarehouse(id int) (int64, error) {
	dbw := container.GetWriter()
	result, err := dbw.Exec("DELETE FROM warehouses WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
