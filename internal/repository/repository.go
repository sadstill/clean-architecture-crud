package repository

import (
	"context"
	"rest-api-crud/internal/repository/user"
)

type UserRepository interface {
	Create(ctx context.Context, user user.UserMongo) (string, error)
	FindById(ctx context.Context, id string) (user.UserMongo, error)
	FindAll(ctx context.Context) ([]user.UserMongo, error)
	Update(ctx context.Context, user user.UserMongo) error
	DeleteById(ctx context.Context, id string) error
}
