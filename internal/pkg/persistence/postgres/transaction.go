package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolTransactionManager struct {
	connection *pgxpool.Pool
}

func (p *PoolTransactionManager) Do(ctx context.Context, f func(ctx context.Context) error) error {
	// TODO: begin transaction
	// TODO: set transaction to context

	if err := f(ctx); err != nil {
		// TODO: rollback transaction

		return err
	}

	// TODO: commit transaction

	return nil
}

type PoolConnection struct {
	pool *pgxpool.Pool
}

func NewPoolConnection(pool *pgxpool.Pool) *PoolConnection {
	return &PoolConnection{pool: pool}
}

func (c *PoolConnection) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	// TODO: если транзакция в контексте, выполнять запрос с ее помощью

	return c.pool.Exec(ctx, sql, arguments...)
}

func (c *PoolConnection) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	// TODO: если транзакция в контексте, выполнять запрос с ее помощью

	return c.pool.Query(ctx, sql, args...)
}

func (c *PoolConnection) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	// TODO: если транзакция в контексте, выполнять запрос с ее помощью

	return c.pool.QueryRow(ctx, sql, args...)
}
