package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"first_name" validate:"required,min=2,max=30"`
	LastName       *string            `json:"last_name" validate:"required,min=2,max=30"`
	Password       *string            `json:"password" validate:"required,min=6"`
	Email          *string            `json:"email" validate:"email, required"`
	Phone          *string            `json:"phone" validate:"required"`
	Token          *string            `json:"token"`
	RefreshToken   *string            `json:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	UserID         string             `json:"user_id"`
	UserCart       []ProductUser      `json:"user_cart" bson:"user_cart"`
	AddressDetails []Address          `json:"address_details" bson:"address"`
	OrderStatus    []Order            `json:"order_status" bson:"orders"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name"`
	Price       *uint64            `json:"price"`
	Rating      *uint              `json:"rating"`
	Image       *string            `json:"image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name"`
	Price       *int               `json:"price"`
	Rating      *uint              `json:"rating"`
	Image       *string            `json:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"_id"`
	House     *string            `json:"house_name" bson:"house_name"`
	Street    *string            `json:"street_name" bson:"street_name"`
	City      *string            `json:"city_name" bson:"city_name"`
	Pincode   *string            `json:"pincode" bson:"pincode"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"_id"`
	OrderCart     []ProductUser      `json:"order_list" bson:"order_list"`
	OrderedAt     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price         int                `json:"price" bson:"total_price"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"is_digital" bson:"is_digital"`
	COD     bool `json:"is_cod" bson:"is_cod"`
}
