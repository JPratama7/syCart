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
	AddCart(ctx context.Context, userId *model.User, itemId primitive.ObjectID, quantity int) (primitive.ObjectID, error)
	RemoveCart(ctx context.Context, userId *model.User, itemId primitive.ObjectID) error
	GetCarts(ctx context.Context, userId *model.User) ([]model.CartItem, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, userId primitive.ObjectID) (model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
}

type OrderRepository interface {
	Checkout(ctx context.Context, userId *model.User) (model.Order, error)
	GetOrder(ctx context.Context, orderId primitive.ObjectID) (model.Order, error)
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, orderId primitive.ObjectID) (model.Payment, error)
	GetPayment(ctx context.Context, paymentId primitive.ObjectID) (model.Payment, error)
	FetchPayment(ctx context.Context, userId *model.User) ([]model.Payment, error)
	UpdatePayment(ctx context.Context, payment *model.Payment) error
}
