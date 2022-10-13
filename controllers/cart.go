package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/khvnkay/ecommerce-yt/database"
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
