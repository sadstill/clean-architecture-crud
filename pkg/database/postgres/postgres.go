package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"rest-api-crud/internal/config"
	"rest-api-crud/pkg/utils"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func New(ctx context.Context, cfg config.PostgresConfig) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	err = utils.DoWithRetries(func() error {
		attemptCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(attemptCtx, dsn)
		if err != nil {
			return err
		}

		err = pool.Ping(attemptCtx)
		if err != nil {
			return err
		}

		return nil
	}, cfg.ConnRetryAttempts, cfg.ConnRetryDelay*time.Second)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to postgres. Error: %w", err)
	}

	return pool, nil
}
