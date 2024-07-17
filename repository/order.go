package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"syCart/model"
)

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database, collectionName string) OrderRepository {
	return &orderRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *orderRepository) Checkout(ctx context.Context, order *model.Order) (res primitive.ObjectID, err error) {

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return
	}

	res, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err = errors.New("inserted id is not primitive.ObjectID")
		return
	}

	return
}

func (r *orderRepository) GetOrder(ctx context.Context, orderId primitive.ObjectID) (model.Order, error) {
	var order model.Order
	filter := bson.M{"_id": orderId}
	err := r.collection.FindOne(ctx, filter).Decode(&order)
	return order, err
}
