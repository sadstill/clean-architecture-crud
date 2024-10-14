package repository

import (
	"context"
	"rest-api-crud/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (string, error)
	FindById(ctx context.Context, id string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.User) error
	DeleteById(ctx context.Context, id string) error
}

type AuthorRepository interface {
}
