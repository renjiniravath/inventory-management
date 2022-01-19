package controller

import "app/model"

func GetInventoryList() (model.ItemListResponse, error) {
	inventoryList, err := model.ReadInventoryList()
	if err != nil {
		return model.ItemListResponse{}, err
	}
	return inventoryList, nil
}
