package repository

import (
	"context"
	"rest-api-crud/internal/repository/users"
)

type UserRepository interface {
	Create(ctx context.Context, user users.UserMongo) (string, error)
	FindById(ctx context.Context, id string) (users.UserMongo, error)
	FindAll(ctx context.Context) ([]users.UserMongo, error)
	Update(ctx context.Context, user users.UserMongo) error
	DeleteById(ctx context.Context, id string) error
}
