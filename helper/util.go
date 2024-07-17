package helper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewTimestamp() primitive.Timestamp {
	return primitive.Timestamp{T: uint32(time.Now().Unix())}
}
