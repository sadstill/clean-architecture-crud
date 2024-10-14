package repository

import (
	"context"
	"rest-api-crud/internal/model"
	"rest-api-crud/pkg/database/postgres"
	"rest-api-crud/pkg/logging"
)

var _ AuthorRepository = (*authorRepo)(nil)

type authorRepo struct {
	client postgres.Client
	logger *logging.Logger
}

func NewAuthorRepo(client postgres.Client, logger *logging.Logger) AuthorRepository {

}

func (a *authorRepo) Create(ctx context.Context, author model.Author) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepo) FindById(ctx context.Context, id string) (model.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepo) FindAll(ctx context.Context) ([]model.Author, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepo) Update(ctx context.Context, author model.Author) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepo) DeleteById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
