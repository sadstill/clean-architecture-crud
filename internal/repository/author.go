package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"rest-api-crud/internal/converter"
	"rest-api-crud/internal/model"
	"rest-api-crud/pkg/database/postgres"
)

var _ AuthorRepository = (*authorRepo)(nil)

type authorRepo struct {
	client postgres.Client
}

func NewAuthorRepo(client postgres.Client) AuthorRepository {
	return &authorRepo{
		client: client,
	}
}

func (a *authorRepo) Create(ctx context.Context, author model.Author) (model.Author, error) {
	storageAuthor := converter.ToStorageAuthor(author)

	q := `INSERT INTO author (name) VALUES ($1) RETURNING id`

	if err := a.client.QueryRow(ctx, q, author.Name).Scan(&storageAuthor.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return model.Author{}, fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQL State: %s. %w",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState(), err)
		}
		return model.Author{}, fmt.Errorf("error creating author: %w", err)
	}

	return converter.ToModelAuthor(storageAuthor), nil
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
