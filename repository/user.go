package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"syCart/model"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) UserRepository {
	return &userRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *userRepository) GetUser(ctx context.Context, userId primitive.ObjectID) (model.User, error) {
	var user model.User
	filter := bson.M{"_id": userId}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func (r *userRepository) GetUserByEmailPass(ctx context.Context, email, password string) (model.User, error) {
	var user model.User
	filter := bson.M{"email": email, "passwordHash": password}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	filter := bson.M{"_id": user.UserId}
	update := bson.M{"$set": user}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
