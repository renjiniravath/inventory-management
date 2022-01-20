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

func EditWarehouse(warehouse model.Warehouse) (int64, error) {
	return model.EditWarehouse(warehouse)
}

func DeleteWarehouse(id int) (int64, error) {
	return model.DeleteWarehouse(id)
}
