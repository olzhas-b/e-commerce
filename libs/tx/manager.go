package tx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"route256/libs/postgresdb"
)

type Manager struct {
	*postgresdb.Postgres
}

func New(pg *postgresdb.Postgres) *Manager {
	return &Manager{Postgres: pg}
}

var txKey = struct{}{}

type DBProvider interface {
	GetDB(ctx context.Context) Querier
}

type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}

type TransactionManager interface {
	RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error
	RunSerializable(ctx context.Context, fn func(ctxTx context.Context) error) error
}

func (m *Manager) RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error {
	tx, err := m.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	})
	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	ctxTx := context.WithValue(ctx, txKey, tx)

	if err = fn(ctxTx); err != nil {
		return fmt.Errorf("exec body: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

func (m *Manager) RunSerializable(ctx context.Context, fn func(ctxTx context.Context) error) error {
	tx, err := m.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return fmt.Errorf("start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	ctxTx := context.WithValue(ctx, txKey, tx)

	if err = fn(ctxTx); err != nil {
		return fmt.Errorf("exec body: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

func (m *Manager) GetDB(ctx context.Context) Querier {
	tx, ok := ctx.Value(txKey).(Querier)
	if ok {
		return tx
	}

	return m.Pool
}
