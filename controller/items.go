package controller

import "app/model"

func GetInventoryList() (model.ItemListResponse, error) {
	inventoryList, err := model.ReadInventoryList()
	if err != nil {
		return model.ItemListResponse{}, err
	}
	return inventoryList, nil
}

func CreateItem(item model.Item) error {
	return model.CreateItem(item)
}

func EditItem(item model.Item) (int64, error) {
	return model.EditItem(item)
}

func DeleteItem(id int) (int64, error) {
	return model.DeleteItem(id)
}
