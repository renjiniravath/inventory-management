package model

import (
	"app/container"
)

//Item holds the details of an item
type Item struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	WarehouseID   string `json:"warehouseId" db:"warehouse_id"`
	WarehouseName string `json:"warehouseName" db:"warehouse_name"`
}

//ItemListResponse holds the response of a items list get request
type ItemListResponse struct {
	Items []Item `json:"items"`
	Count int    `json:"count"`
}

//ReadInventoryList queries inventory list from db
func ReadInventoryList() (ItemListResponse, error) {
	dbr := container.GetReader()
	items := make([]Item, 0)
	err := dbr.Select(&items, "SELECT items.id, items.name, warehouse_id, warehouses.name AS warehouse_name FROM items JOIN warehouses ON warehouses.id=warehouse_id")
	if err != nil {
		return ItemListResponse{}, err
	}
	itemListResponse := ItemListResponse{
		Items: items,
		Count: len(items),
	}
	return itemListResponse, nil
}
