package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khvnkay/ecommerce-yt/database"
	"github.com/khvnkay/ecommerce-yt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"errors"
	"log"
	"net/http"
	"time"
)
type Application struct {
	productCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewAppplication(prodCollaction, userCollaction, *mongo.Collection) *Application {
	return &Application {
		productCollection: prodCollaction,
		userCollection: userCollaction
	}
}

func (app *Application) AddtoCart() gin.Handler {
	return func (c *gin.Context)  {
		prodcyQueryID := c.Query("id")
		if prodcyQueryID == "" {
			log.Println("product Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product is empty"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user is empty"))
			return

		}
		productId, err := primitive.ObjectIDFromHex(prodcyQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.productCollection, app.userCollection, productId, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully added to thecart")

		
	}



}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func (c *gin.Context)  {
		prodcyQueryID := c.Query("id")
		if prodcyQueryID == "" {
			log.Println("product Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product is empty"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user is empty"))
			return

		}
		productId, err := primitive.ObjectIDFromHex(prodcyQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.RemoveItemFromCart(ctx, app.productCollection, app.userCollection, productId, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "remove item from  cart")

	}

}
func GetItemFromCart() gin.HandlerFunc {
	return func (c *gin.Context)  {
		user_id := c.Query("id")
		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
			c.Abort()
			return
		}

		usert_id, _ :=primitive.ObjectIDFromHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)

		defer cancel()

		var filledcar models.User
		err :=UserCollection.FindOne(ctx, bson.d{primitive.E{Key: "_id",Value: usert_id}}).Decode(&filledcar)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}


		filter_math := bson.D{{Key: "$math", Value: bson.D{primitive.E{Key: "_id", Value:   usert_id}}}}
		unwind := bson.D{{Key: "%unwind", Value: bson.D{primitive.E{Key: "path", Value: "$usercart"}}}}
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{primitive.E{Key: "$sum", Value: "$usercart.price"}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_math, unwind, grouping})
		if err != nil {
			log.Panicln(err)
		}
		var listing []bson.M
		if err  = pointcursor.All(ctx, &listing); err !=nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}


		for _,json := range ;liglisting {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledcart.UserCart)
		}
		ctx.Done()



	}

}

func GetItemByCart() gin.HandlerFunc {
	return func (c *gin.Context)  {
		prodcyQueryID := c.Query("id")
		if prodcyQueryID == "" {
			log.Println("product Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product is empty"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user is empty"))
			return

		}
		productId, err := primitive.ObjectIDFromHex(prodcyQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.g(ctx, app.productCollection, app.userCollection, productId, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully added to thecart")

	}

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func (c *gin.Context)  {
		prodcyQueryID := c.Query("id")
		if prodcyQueryID == "" {
			log.Println("product Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product is empty"))
			return
		}
		
		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully buy item from  cart ")

	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {

return func (c *gin.Context)  {
		prodcyQueryID := c.Query("id")
		if prodcyQueryID == "" {
			log.Println("product Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("Product is empty"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user Id is em,pty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user is empty"))
			return

		}
		productId, err := primitive.ObjectIDFromHex(prodcyQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()

		err = database.InstantBuy(ctx, app.productCollection, app.userCollection, productId, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully place the order")

	}
}
