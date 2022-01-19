package controller

import "app/model"

func GetWarehouseList() (model.WarehouseListResponse, error) {
	warehouseListResponse, err := model.ReadWarehouseList()
	if err != nil {
		return model.WarehouseListResponse{}, err
	}
	return warehouseListResponse, nil
}

func CreateWarehouse(warehouse model.Warehouse) error {
	return model.CreateWarehouse(warehouse)
}
