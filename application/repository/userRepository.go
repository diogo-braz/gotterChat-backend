package repository

import (
	"context"
	"errors"

	"github.com/joaovds/chat/application/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

func (ur *UserRepository) GetUserByNickName(nickName string) (*domain.User, error) {
	filter := bson.M{"nickname": nickName}
	var user domain.User
	result := ur.Collection.FindOne(context.TODO(), filter)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(user *domain.User) error {
	_, err := ur.Collection.InsertOne(context.TODO(), user)
	return err
}

func (ur *UserRepository) DeleteUser(nickname string) (*mongo.DeleteResult, error) {
	filter := bson.M{"nickname": nickname}
	result, err := ur.Collection.DeleteOne(context.TODO(), filter)
	return result, err
}

func (ur *UserRepository) UpdateUser(user *domain.User) (*mongo.UpdateResult, error) {
	filter := bson.M{"nickname": user.Nickname}

	update := bson.M{
		"$set": bson.M{
			"nickname":    user.Nickname,
			"password":    user.Password,
			"gender":      user.Gender,
			"phoneNumber": user.PhoneNumber,
			"interests":   user.Interests,
		},
	}

	result, err := ur.Collection.UpdateOne(context.TODO(), filter, update)

	if result.ModifiedCount == 0 {
		return result, errors.New("User not found or no changes applied")
	}

	return result, err
}
