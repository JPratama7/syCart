package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	UserId       primitive.ObjectID `json:"user_id" bson:"user_id"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type Product struct {
	ProductId     primitive.ObjectID `json:"product_id" bson:"product_id"`
	CategoryId    int                `json:"category_id" bson:"category_id"`
	Name          string             `json:"name" bson:"name"`
	Description   string             `json:"description" bson:"description"`
	Price         float64            `json:"price" bson:"price"`
	StockQuantity int                `json:"stock_quantity" bson:"stock_quantity"`
	ImageUrl      string             `json:"image_url" bson:"image_url"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}

type Category struct {
	CategoryId  primitive.ObjectID `json:"category_id" bson:"category_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
}

type CartItem struct {
	CartItemId primitive.ObjectID `json:"cart_item_id" bson:"cart_item_id"`
	UserId     primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductId  primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	AddedAt    time.Time          `json:"added_at" bson:"added_at"`
}

type Order struct {
	OrderId     primitive.ObjectID `json:"order_id" bson:"order_id"`
	UserId      primitive.ObjectID `json:"user_id" bson:"user_id"`
	TotalAmount float64            `json:"total_amount" bson:"total_amount"`
	Status      string             `json:"status" bson:"status"`
	OrderItems  *[]OrderItem       `json:"order_items" bson:"order_items"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type OrderItem struct {
	ProductId       primitive.ObjectID `json:"product_id" bson:"product_id"`
	Quantity        int                `json:"quantity" bson:"quantity"`
	PriceAtPurchase float64            `json:"price_at_purchase" bson:"price_at_purchase"`
}

type Payment struct {
	PaymentId     primitive.ObjectID `json:"payment_id" bson:"payment_id"`
	OrderId       primitive.ObjectID `json:"order_id" bson:"order_id"`
	Amount        float64            `json:"amount" bson:"amount"`
	Status        string             `json:"status" bson:"status"`
	TransactionId string             `json:"transaction_id" bson:"transaction_id"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
}
