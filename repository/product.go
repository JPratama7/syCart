package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"syCart/model"
)

type productImpl struct {
	coll *mongo.Collection
}

func NewProductImpl(db *mongo.Database, coll string) ProductRepository {
	return &productImpl{coll: db.Collection(coll)}
}

func (p *productImpl) GetProduct(ctx context.Context, productId primitive.ObjectID) (res model.Product, err error) {
	err = p.
		coll.
		FindOne(ctx, bson.M{"_id": productId}).
		Decode(&res)

	return
}

func (p *productImpl) FetchProduct(ctx context.Context) (res []model.Product, err error) {
	curr, err := p.coll.Find(ctx, bson.M{})

	if err != nil {
		return
	}

	err = curr.All(ctx, &res)
	return
}

func (p *productImpl) GetProductByCategoryName(ctx context.Context, categoryName string) (res []model.Product, err error) {
	pipeline := mongo.Pipeline{
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "categories"},
					{"localField", "category_id"},
					{"foreignField", "category_id"},
					{"as", "category"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$category"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{
				"$match", bson.D{
					{
						"category.name", bson.D{
							{
								"$regex", categoryName,
							},
						},
					},
				},
			},
		},
	}

	curr, err := p.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	err = curr.All(ctx, &res)

	return
}
