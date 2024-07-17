package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"syCart/model"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, productId primitive.ObjectID) (model.Product, error)
	FetchProduct(ctx context.Context) ([]model.Product, error)
	GetProductByCategoryName(ctx context.Context, categoryName string) ([]model.Product, error)
}

type CartRepository interface {
	AddCart(ctx context.Context, userId *model.User, productId primitive.ObjectID, quantity int) (primitive.ObjectID, error)
	UpdateCart(ctx context.Context, userId *model.User, cart *model.CartItem) error
	GetCartItem(ctx context.Context, userId *model.User, productId primitive.ObjectID) (model.CartItem, error)
	CartsWithProduct(ctx context.Context, userId *model.User) ([]model.CartItemWithProduct, error)
	RemoveCart(ctx context.Context, userId *model.User, itemId primitive.ObjectID) error
	GetCarts(ctx context.Context, userId *model.User) ([]model.CartItem, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, userId primitive.ObjectID) (model.User, error)
	GetUserByEmailPass(ctx context.Context, username, password string) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	UsernameOrEmailExist(ctx context.Context, username, email string) (model.User, error)
}

type OrderRepository interface {
	Checkout(ctx context.Context, user *model.User, items *[]model.OrderItem) (primitive.ObjectID, error)
	GetOrder(ctx context.Context, orderId primitive.ObjectID) (model.Order, error)
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, order *model.Order) (primitive.ObjectID, error)
	GetPayment(ctx context.Context, paymentId primitive.ObjectID) (model.Payment, error)
	FetchPayment(ctx context.Context, userId *model.User) ([]model.Payment, error)
	UpdatePayment(ctx context.Context, payment *model.Payment) error
}
