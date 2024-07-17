package repository

import (
	"context"
	"errors"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"syCart/model"
)

type paymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(db *mongo.Database, collectionName string) PaymentRepository {
	return &paymentRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, orderId primitive.ObjectID, order *model.Order) (res primitive.ObjectID, err error) {
	payment := model.Payment{
		OrderId:       orderId,
		Amount:        order.TotalAmount,
		Status:        "PENDING",
		TransactionId: "TRN-" + xid.New().String(),
	}
	result, err := r.collection.InsertOne(ctx, payment)
	if err != nil {
		return
	}
	res, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		err = errors.New("payment id is not primitive.ObjectID")
		return
	}

	return
}

func (r *paymentRepository) GetPayment(ctx context.Context, paymentId primitive.ObjectID) (res model.Payment, err error) {
	filter := bson.M{"_id": paymentId}
	err = r.collection.FindOne(ctx, filter).Decode(&res)
	return
}

func (r *paymentRepository) FetchPayment(ctx context.Context, user *model.User) (res []model.Payment, err error) {
	filter := bson.M{"user_id": user.UserId}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &res)

	return
}
func (r *paymentRepository) FetchPaymentByUser(ctx context.Context, user *model.User) (res []model.Payment, err error) {
	filter := mongo.Pipeline{
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "orders"},
					{"localField", "order_id"},
					{"foreignField", "_id"},
					{"as", "order"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$order"},
					{"preserveNullAndEmptyArrays", false},
				},
			},
		},
		bson.D{{"$match", bson.D{{"order.user_id", user.UserId}}}},
	}
	cursor, err := r.collection.Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &res)
	return
}

func (r *paymentRepository) UpdatePayment(ctx context.Context, payment *model.Payment) error {
	filter := bson.M{"_id": payment.PaymentId}
	update := bson.M{"$set": payment}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
