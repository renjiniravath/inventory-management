package model

import (
	"app/container"
	"strings"
)

//Item holds the details of an item
type Item struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	WarehouseID   int    `json:"warehouseId" db:"warehouse_id"`
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

//CreateItem creates a new item entry in db
func CreateItem(item Item) error {
	dbw := container.GetWriter()
	_, err := dbw.Exec("INSERT INTO items (name, warehouse_id) VALUES (?,?)", item.Name, item.WarehouseID)
	if err != nil {
		return err
	}
	return nil
}

//EditItem updates the details of a item in db
func EditItem(item Item) (int64, error) {
	dbw := container.GetWriter()
	fields := []string{}
	var values []interface{}
	if item.Name != "" {
		fields = append(fields, " name = ? ")
		values = append(values, item.Name)
	}
	if item.WarehouseID != 0 {
		fields = append(fields, " warehouse_id = ? ")
		values = append(values, item.WarehouseID)
	}
	values = append(values, item.ID)
	query := "UPDATE items SET " + strings.Join(fields, " , ") + " WHERE id = ?"
	result, err := dbw.Exec(query, values...)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

//DeleteItem deletes a item from db
func DeleteItem(id int) (int64, error) {
	dbw := container.GetWriter()
	result, err := dbw.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
