package store

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserStorage struct {
	col *mongo.Collection
}

func (u *UserStorage) Create(ctx context.Context, user *User) error {
	return nil
}