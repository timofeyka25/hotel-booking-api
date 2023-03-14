package db

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

type TransactionRepository struct {
	db *bun.DB
}

func NewTransactionRepository(db *bun.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) NewSelect(ctx context.Context) *bun.SelectQuery {
	tx := extractTransaction(ctx)
	if tx != nil {
		return tx.NewSelect()
	}

	return r.db.NewSelect()
}

func (r *TransactionRepository) NewInsert(ctx context.Context) *bun.InsertQuery {
	tx := extractTransaction(ctx)
	if tx != nil {
		return tx.NewInsert()
	}

	return r.db.NewInsert()
}

func (r *TransactionRepository) NewUpdate(ctx context.Context) *bun.UpdateQuery {
	tx := extractTransaction(ctx)
	if tx != nil {
		return tx.NewUpdate()
	}

	return r.db.NewUpdate()
}

func (r *TransactionRepository) NewDelete(ctx context.Context) *bun.DeleteQuery {
	tx := extractTransaction(ctx)
	if tx != nil {
		return tx.NewDelete()
	}

	return r.db.NewDelete()
}

func (r *TransactionRepository) Query(ctx context.Context, query string) (*sql.Rows, error) {
	tx := extractTransaction(ctx)
	if tx != nil {
		return tx.Query(query)
	}

	return r.db.Query(query)
}

func extractTransaction(ctx context.Context) *bun.Tx {
	if tx, ok := ctx.Value(transactionKey{}).(*bun.Tx); ok {
		return tx
	}

	return nil
}

type transactionKey struct{}
