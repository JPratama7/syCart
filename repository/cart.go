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

func NewCartRepository(db *mongo.Database, coll string) CartRepository {
	return &cartImpl{coll: db.Collection(coll)}
}

func (p *cartImpl) AddCart(ctx context.Context, userId *model.User, productId primitive.ObjectID, quantity int) (res primitive.ObjectID, err error) {
	cur, err := p.coll.InsertOne(ctx, model.CartItem{
		UserId:    userId.UserId,
		ProductId: productId,
		Quantity:  quantity,
		AddedAt:   helper.NewDatetime(),
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

func (p *cartImpl) UpdateCart(ctx context.Context, userId *model.User, cart *model.CartItem) (err error) {
	_, err = p.coll.UpdateOne(ctx, bson.M{
		"user_id": userId.UserId,
		"_id":     cart.CartItemId,
	}, bson.M{
		"$set": bson.M{
			"quantity": cart.Quantity,
			"added_at": helper.NewDatetime(),
		},
	})
	return
}

func (p *cartImpl) GetCartItem(ctx context.Context, userId *model.User, productId primitive.ObjectID) (res model.CartItem, err error) {
	err = p.coll.FindOne(ctx, bson.M{
		"user_id":    userId.UserId,
		"product_id": productId,
	}).Decode(&res)

	return
}

func (p *cartImpl) RemoveCart(ctx context.Context, userId *model.User, cartId primitive.ObjectID) (err error) {
	_, err = p.coll.DeleteOne(ctx, bson.M{
		"user_id": userId.UserId,
		"_id":     cartId,
	})
	return
}

func (p *cartImpl) RemoveAll(ctx context.Context, userId *model.User) (err error) {
	_, err = p.coll.DeleteMany(ctx, bson.M{
		"user_id": userId.UserId,
	})
	return
}

func (p *cartImpl) CartsWithProduct(ctx context.Context, userId *model.User) (res []model.CartItemWithProduct, err error) {
	filter := mongo.Pipeline{bson.D{
		{"$lookup",
			bson.D{
				{"from", "products"},
				{"localField", "product_id"},
				{"foreignField", "_id"},
				{"as", "product"},
			},
		},
	},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$product"},
					{"preserveNullAndEmptyArrays", false},
				},
			},
		},
		bson.D{{"$match", bson.D{{"user_id", userId.UserId}}}}}

	curr, err := p.coll.Aggregate(ctx, filter)
	if err != nil {
		return
	}

	err = curr.All(ctx, &res)
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
