package routes

import (
	"context"
	"net/http"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New() //global veriable declaration

var orderCollection *mongo.Collection = openCollection(Client, "order") //it will create a collection in mongoDb "order"

func AddOrder(c *gin.Context) {
	// order := models.Order{
	// 	Dish: "Rajma chawal",
	// 	Price: 69.69,
	// 	Server: "ram",
	// 	Table: "1",
	// }
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel() //at last we have to canel the function
	var order models.Order

	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error is": "Compelete the order :)"})
		return //if order is empty or have no value other wise it will go for validation (given below )
	}

	// if validationErr := validate.Struct(order); validationErr != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error is": " Validation Error "})
	// }
	validatioonErr := validate.Struct(order) // it will check which values have "validation required"
	if validatioonErr != nil {
		mg := "Validatioon error is : " + validatioonErr.Error()
		c.JSON(http.StatusBadRequest, mg) //if the value of server and dish isn't given then it will show validation error
		return
	}
	order.ID = primitive.NewObjectID()                     // creating new premitive id of
	inserted, err := orderCollection.InsertOne(ctx, order) //now we need the data entry that we want to enter in our database
	//it will insert the data in out new collection "order"
	if err != nil {
		msg := "order item not created" + err.Error() //basic error handling
		c.JSON(http.StatusInternalServerError, gin.H{"Error ": msg})
		return
	}

	c.JSON(http.StatusOK, inserted) //status OK and data is inserted in collection
}

func GetOrders(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var orders []bson.M // creating a bson which have all the orders in it

	cursor, err := orderCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error :": err})
	}
	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error : ": err.Error()})
		return
	}
	//fmt.Println(orders)  //to print the data in terminal as well:)
	c.JSON(http.StatusOK, orders) //at last print or show all the orders
}

func GetById(c *gin.Context) {
	orderId := c.Params.ByName("id")
	docId, err := primitive.ObjectIDFromHex(orderId)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERRor: ": err.Error()})
		return
	}
	var order bson.M

	if err := orderCollection.FindOne(ctx, bson.M{"_id": docId}).Decode(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error : ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)

}

func GetByWaiterName(c *gin.Context) {
	orderWaiter := c.Params.ByName("waiter") //taking the value that user give to execess the name of waiter
	// docId, err := primitive.ObjectIDFromHex(orderWaiter) //since waiter or server isn't empty cause it was in valitator
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	// }
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var order []bson.M                                                      //whenever we have to GET more values then we need to use this   // [] it is important because it is a map
	cursor, err := orderCollection.Find(ctx, bson.M{"server": orderWaiter}) //checking the value
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: ": err.Error()}) //checking the error
		return
	}
	if err = cursor.All(ctx, &order); err != nil { // taking all the value which have the value "orderWaiter"
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order) //Finally printing to GETting the values :+)

}

func UpdateOrder(c *gin.Context) {
	orderID := c.Params.ByName("id")               //what id is pased by the client /:id this type
	docId, _ := primitive.ObjectIDFromHex(orderID) //no error because waiter name isn't empty here and it is always some value or name

	var orderNew models.Order                     //creating new order
	if err := c.BindJSON(&orderNew); err != nil { // taking from the body and checking that there is some value
		c.JSON(http.StatusBadRequest, gin.H{"Error is : ": err.Error()})
		return
	}
	validationErr := validate.Struct(orderNew) // now checking that all the required values like server and dish is compolsory and if not given then error will occur
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error is : ": validationErr.Error()}) //badrequest from client
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	res, err := orderCollection.ReplaceOne(ctx, bson.M{"_id": docId}, bson.M{ //replacing all the values from old order to orderNew
		"dish":   orderNew.Dish,
		"price":  orderNew.Price,
		"server": orderNew.Server,
		"table":  orderNew.Table,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error : ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.ModifiedCount) //if everything ok it will return count of modification like "1", "2"

}
func UpdateWaiterNameById(c *gin.Context) {
	orderId := c.Params.ByName("id")
	docId, _ := primitive.ObjectIDFromHex(orderId)

	var waiter models.Waiter

	if err := c.BindJSON(&waiter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error is : ": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	res, err := orderCollection.UpdateOne(ctx, bson.M{"_id": docId}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "server", Value: waiter.Server},
		}},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error is : ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.ModifiedCount)

}
func DeleteById(c *gin.Context) {
	orderId := c.Params.ByName("id")
	docId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	deleteResult, err := orderCollection.DeleteOne(ctx, bson.M{"_id": docId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error : ": err.Error()})
		return
	}
	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted: ": deleteResult.DeletedCount})

}
