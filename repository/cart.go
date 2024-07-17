package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"syCart/helper"
	"syCart/model"
)

type cartImpl struct {
	coll *mongo.Collection
}

func NewCartImpl(db *mongo.Database, coll string) CartRepository {
	return &cartImpl{coll: db.Collection(coll)}
}

func (p *cartImpl) AddCart(ctx context.Context, userId *model.User, itemId primitive.ObjectID, quantity int) (res primitive.ObjectID, err error) {
	cur, err := p.coll.InsertOne(ctx, model.CartItem{
		UserId:    userId.UserId,
		ProductId: itemId,
		Quantity:  quantity,
		AddedAt:   helper.NewTimestamp(),
	})

	if err != nil {
		return
	}

	if id, ok := cur.InsertedID.(primitive.ObjectID); ok {
		res = id
		return
	}
	return
}

func (p *cartImpl) RemoveCart(ctx context.Context, userId *model.User, cartId primitive.ObjectID) (err error) {
	_, err = p.coll.DeleteOne(ctx, bson.M{
		"user_id":      userId.UserId,
		"cart_item_id": cartId,
	})
	return
}

func (p *cartImpl) GetCarts(ctx context.Context, userId *model.User) (res []model.CartItem, err error) {
	cur, err := p.coll.Find(ctx, bson.M{
		"user_id": userId.UserId,
	})

	if err != nil {
		return
	}
	err = cur.All(ctx, &res)
	return
}
