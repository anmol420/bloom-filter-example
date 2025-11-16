package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserStorage struct {
	col *mongo.Collection
}

func (u *UserStorage) Create(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	res, err := u.col.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	user.ID = res.InsertedID.(bson.ObjectID)
	return nil
}

func (u *UserStorage) FindUser(ctx context.Context, username string, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
			{"email": email},
		},
	}
	var user User
	err := u.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *UserStorage) SearchByUsername(ctx context.Context, username string) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	filter := bson.M{
		"$or": []bson.M{
			{"username": username},
		},
	}
	var user User
	err := u.col.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}