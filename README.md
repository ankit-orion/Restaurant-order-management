# Restaurant Order Management API

This repository contains an API for managing orders in a restaurant setting. The API is built using Go (Golang) and utilizes Gin framework for routing and MongoDB for data storage.

## Getting Started

To run the API locally, follow these steps:

1. Clone the repository: git clone <https://github.com/TusharPachouri/Video-Streaming-Application.git>

2. Install dependencies:

3. Set up MongoDB and configure the connection details in the `config.go` file.

4. Run the server:

The server will start running on `localhost:8080` by default.

## Endpoints

### Add Order

- **URL:** `/order/create`
- **Method:** `POST`
- **Function:** `routes.AddOrder`
- **Description:** Add a new order.

### Get Orders

- **URL:** `/orders`
- **Method:** `GET`
- **Function:** `routes.GetOrders`
- **Description:** Retrieve all orders.

### Get Order by ID

- **URL:** `/order/:id/`
- **Method:** `GET`
- **Function:** `routes.GetById`
- **Description:** Retrieve an order by its ID.

### Get Orders by Waiter Name

- **URL:** `/waiter/:waiter`
- **Method:** `GET`
- **Function:** `routes.GetByWaiterName`
- **Description:** Retrieve orders served by a specific waiter.

### Update Order

- **URL:** `/order/update/:id/`
- **Method:** `PUT`
- **Function:** `routes.UpdateOrder`
- **Description:** Update an order by its ID.

### Update Waiter Name by ID

- **URL:** `/waiter/update/:id`
- **Method:** `PUT`
- **Function:** `routes.UpdateWaiterNameById`
- **Description:** Update the name of the waiter associated with an order.

### Delete Order by ID

- **URL:** `/order/delete/:id`
- **Method:** `DELETE`
- **Function:** `routes.DeleteById`
- **Description:** Delete an order by its ID.

### Delete Orders by Waiter

- **URL:** `/waiter/delete/:waiter`
- **Method:** `DELETE`
- **Function:** `routes.DeleteByWaiter`
- **Description:** Delete all orders served by a specific waiter.

## Models

### Order

type Order struct {
    ID     primitive.ObjectID `bson:"_id"`
    Dish   string             `json:"dish" validate:"required"`
    Price  float32            `json:"price"`
    Server string             `json:"server" validate:"required"`
    Table  string             `json:"table"`
}

type Waiter struct {
    Server string `json:"server"`
}

## Dependencies

- **Gin**: Web framework for Go.
- **MongoDB Go Driver**: MongoDB driver for Go.

## Contributing

Contributions are welcome. Before contributing, please open an issue to discuss the changes you would like to make.
