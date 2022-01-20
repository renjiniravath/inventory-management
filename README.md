# inventory-management

An inventory management application for a logistics company

## Setting up

- Install Docker

  Follow the steps in this link to install docker in your system
  https://docs.docker.com/get-docker/

## Running

### Run the following command in the root folder to run the program

- docker-compose up

### Safely stopping the system

- Ctrl+C

- docker-compose down

## Endpoints

### Root URL - `http://localhost:8080`

- This is the root url where the API is hosted. Prepend this root before all requests

### GET `/item`

Lists all items in the inventory with name and warehouse details

#### Response

```json
{
  "items": [
    {
      "id": 1,
      "name": "Item 2",
      "warehouseId": 1,
      "warehouseName": "RDU"
    },
    {
      "id": 2,
      "name": "Item 1",
      "warehouseId": 2,
      "warehouseName": "JFK"
    }
  ],
  "count": 2
}
```

### POST `/item`

Adds a item to the inventory

#### Request

```json
{
  "name": "Item 1",
  "warehouseId": 2
}
```

### Response

```
Item successfully created
```

### PUT `/item/:id`

Updates the details of an item in the inventory

#### Request

Provide atleast one of the fields

```json
{
  "name": "Item 1",
  "warehouseId": 2
}
```

### Response

```
Item successfully edited
```

### DELETE `/item/:id`

Deletes an item from inventory

### Response

```
Item successfully deleted
```

### GET `/warehouse`

Lists all warehouses

#### Response

```json
{
  "warehouses": [
    {
      "id": 1,
      "name": "RDU"
    },
    {
      "id": 2,
      "name": "JFK"
    }
  ],
  "count": 2
}
```

### POST `/warehouse`

Adds a warehouse

#### Request

```json
{
  "name": "JFK"
}
```

### Response

```
Warehouse successfully created
```

### PUT `/warehouse/:id`

Updates the name of a warehouse

#### Request

```json
{
  "name": "JFK"
}
```

### Response

```
Warehouse successfully edited
```

### DELETE `/warehouse/:id`

Deletes an warehouse

### Response

```
Warehouse successfully deleted
```
