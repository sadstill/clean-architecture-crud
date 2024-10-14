package converter

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"rest-api-crud/internal/model"
	"rest-api-crud/internal/storage"
)

func ToUser(userMongo storage.UserMongo) model.User {
	return model.User{
		ID:       userMongo.ID.Hex(),
		Username: userMongo.Username,
		Email:    userMongo.Email,
	}
}

func ToUserMongo(user model.User) (storage.UserMongo, error) {
	oid, err := bson.ObjectIDFromHex(user.ID)
	if err != nil {
		return storage.UserMongo{}, fmt.Errorf("failed to convert user ID to ObjectID: %v", err)
	}

	return storage.UserMongo{
		ID:       oid,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func ToUserSlice(usersMongo []storage.UserMongo) []model.User {
	var users []model.User

	for _, userMongo := range usersMongo {
		users = append(users, ToUser(userMongo))
	}

	return users
}
