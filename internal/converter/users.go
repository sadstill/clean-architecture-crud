package converter

import (
	"rest-api-crud/internal/model"
	"rest-api-crud/internal/repository/users"
	"strconv"
)

func ToCreateUserResponseFromUserPG(user *users.UserPG) *model.CreateUserResponse {
	return &model.CreateUserResponse{
		ID: strconv.FormatInt(user.ID, 10),
	}
}

func ToCreateUserResponseFromUserMongo(user *users.UserMongo) *model.CreateUserResponse {
	return &model.CreateUserResponse{
		ID: user.ID.Hex(),
	}
}
