package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrNotFound          = mongo.ErrNoDocuments
	QueryTimeoutDuration = 5 * time.Second
)

type Storage struct {
	Users interface {
		Create(context.Context, *User) error
	}
	client *mongo.Client
}

func NewMongoStorage(client *mongo.Client, dbName string) Storage {
	db := client.Database(dbName)
	return Storage{
		Users: &UserStorage{ col: db.Collection("users") },
		client: client,
	}
}