package converter

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"rest-api-crud/internal/model"
	"rest-api-crud/internal/storage"
)

func ToModelUser(userMongo storage.User) model.User {
	return model.User{
		ID:       userMongo.ID.Hex(),
		Username: userMongo.Username,
		Email:    userMongo.Email,
	}
}

func ToStorageUser(user model.User) (storage.User, error) {
	oid, err := bson.ObjectIDFromHex(user.ID)
	if err != nil {
		return storage.User{}, fmt.Errorf("failed to convert user ID to ObjectID: %v", err)
	}

	return storage.User{
		ID:       oid,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func ToModelUserSlice(usersMongo []storage.User) []model.User {
	var users []model.User

	for _, userMongo := range usersMongo {
		users = append(users, ToModelUser(userMongo))
	}

	return users
}
