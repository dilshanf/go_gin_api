// Golang REST API with API versioning and Basic Auth
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURL = "mongodb://localhost:27017"
var dbName = "mydb"

func GetMethod(c *gin.Context) {
	fmt.Println("\n'GetMethod' called")
	IdValue := c.Params.ByName("IdValue")
	message := "GetMethod Called With Param: " + IdValue
	c.JSON(http.StatusOK, message)

	ReqPayload := make([]byte, 1024)
	ReqPayload, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Request Payload Data: ", string(ReqPayload))
}

type Person struct {
	Name string
	Age  int
	City string
}

func insertDoc() {
	collection := client.Database("mydb").Collection("persons")

	ruan := Person{"Ruan", 34, "Cape Town"}

	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}

func main() {
	router := gin.Default()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Create Sub Router for  customised API version
	subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
		"admin": "adminpass",
	}))

	subRouterAuthenticated.GET("/:IdValue", GetMethod)

	listenPort := "1357"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}
