package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primative"
)

type User struct {
	ID              primative.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name"  validate:"required,min=2, max=30" `
	Last_Name       *string            `json:"last_name"  validate:"required,min=2, max=30"`
	Password        *string            `json:"password"  validate:"required,min=2, max=30"`
	Email           *string            `json: "email" validate:"email ,required "`
	Phone           *string            `json: "phone"  validate:"required "`
	Token           *string            `json: "token"`
	Refresh_Token   *string            `json: "refresh_token"`
	Created_At      time.Time          `json: "created_at"`
	Updated_At      time.Time          `json: "updated_at"`
	User_ID         string             `json: "user_id"`
	UserCart        []ProductUser      `json: "usercart"`
	Address_Details []Address          `json: "address_details"`
	Order_Status    []Order            `json: "order_status"`
}

type Product struct {
	Product_ID   primative.ObjectID `bson:"_id`
	Product_Name *string            `json: "product_name"`
	Price        int                `json: "price"`
	Rating       *uint8             `json: "rating"`
	Images       *string            `json: "images"`
}

type ProductUser struct {
	Product_ID   primative.ObjectID `bson:"_id`
	Product_Name *string            `json: "product_name" bson:"product_name"`
	Price        *uint64            `json: "price" bson: "price"`
	Rating       *uint8             `json: "rating" bson: "rating"`
	Images       *string            `json: "images" bson: "images"`
}

type Address struct {
	Address_id primative.ObjectID `bson:"_id`
	Hosue      *string            `json: "house_name" bson:"house_name"`
	Street     *string            `json: "street_name" bson:"street_name"`
	City       *string            `json: "city_name" bson:"city_name"`
	Pincode    *string            `json: "pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primative.ObjectID `bson:"_id`
	Order_Cart     []ProductUser      `json: "order_list" bson:"order_list"`
	Ordered_At     time.Time          `json: "ordered_at" bson:"ordered_at"`
	Price          int                `json: "total_price" bson:"total_price"`
	Discount       *int               `json: "discount" bson:"discount"`
	Payment_Method Payment            `json: "payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
