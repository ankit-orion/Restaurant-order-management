package main

import (
	"os"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT") //PORT is present in env file
	// if port == "" {
	// 	port = "5000"
	// }
	//create -->
	router := gin.New()                           // created a router (server for restaurant)
	router.POST("/order/create", routes.AddOrder) // POST --> and go to AddOrder function

	//read the data -->
	router.GET("/orders", routes.GetOrders) // GET  --> all the orders at once
	router.GET("/order/:id/", routes.GetById)
	router.GET("/waiter/:waiter", routes.GetByWaiterName) // it will find all the orders having same waiter name

	//Update
	router.PUT("/order/update/:id/", routes.UpdateOrder) //will update an order using its ID
	router.PUT("/waiter/update/:id", routes.UpdateWaiterNameById)

	//delete
	router.DELETE("/order/delete/:id", routes.DeleteById)

	router.Run(":" + port) // port --> that is equal to the PORT given in .env file
}
