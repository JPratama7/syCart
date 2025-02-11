package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	PENDING Status = "PENDING"
	SUCCESS Status = "SUCCESS"
)

type Token struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	UserId       primitive.ObjectID `json:"user_id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Salt         string             `json:"-" bson:"salt"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	CreatedAt    primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt    primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type Product struct {
	ProductId     primitive.ObjectID `json:"product_id" bson:"_id,omitempty"`
	CategoryId    primitive.ObjectID `json:"category_id" bson:"category_id"`
	Name          string             `json:"name" bson:"name"`
	Description   string             `json:"description" bson:"description"`
	Price         float64            `json:"price" bson:"price"`
	StockQuantity int                `json:"stock_quantity" bson:"stock_quantity"`
	ImageUrl      string             `json:"image_url" bson:"image_url"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type Category struct {
	CategoryId  primitive.ObjectID `json:"category_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

type CartItem struct {
	CartItemId primitive.ObjectID `json:"cart_item_id" bson:"_id,omitempty"`
	UserId     primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductId  primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	AddedAt    primitive.DateTime `json:"added_at" bson:"added_at"`
}

type CartItemWithProduct struct {
	CartItemId primitive.ObjectID `json:"cart_item_id" bson:"_id,omitempty"`
	UserId     primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductId  primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	AddedAt    primitive.DateTime `json:"added_at" bson:"added_at"`
	Product    Product            `bson:"product"`
}

type Order struct {
	OrderId     primitive.ObjectID `json:"order_id" bson:"_id,omitempty"`
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id"`
	TotalAmount float64            `json:"total_amount" bson:"total_amount"`
	Status      Status             `json:"status" bson:"status"`
	OrderItems  *[]OrderItem       `json:"order_items" bson:"order_items"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

type OrderItem struct {
	ProductId       primitive.ObjectID `json:"product_id" bson:"_id,omitempty"`
	Quantity        int                `json:"quantity" bson:"quantity"`
	PriceAtPurchase float64            `json:"price_at_purchase" bson:"price_at_purchase"`
}

type Payment struct {
	PaymentId     primitive.ObjectID `json:"payment_id" bson:"_id,omitempty"`
	OrderId       primitive.ObjectID `json:"order_id" bson:"order_id"`
	Amount        float64            `json:"amount" bson:"amount"`
	Status        Status             `json:"status" bson:"status"`
	TransactionId string             `json:"transaction_id" bson:"transaction_id"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`
}
