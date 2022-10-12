package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancal()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("oail to connect mongi")
		return nil
	}

	fmt.Println("suceessfullly connnect")

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, colllectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("Ecommerce").Collection(colllectionName)

	return collection

}

func ProductData(client *mongo.Client, colllectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(colllectionName)
	return productCollection

}
