package repository

import (
	"context"
	"rest-api-crud/internal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.User) (string, error)
	FindById(ctx context.Context, id string) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
	Update(ctx context.Context, user domain.User) error
	DeleteById(ctx context.Context, id string) error
}
