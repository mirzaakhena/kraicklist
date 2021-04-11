package repository

import "context"

// TransactionRepo used for common transaction handling
// all the context must use the same database session.
type TransactionRepo interface {

	BeginTransaction(ctx context.Context) (context.Context, error)

	CommitTransaction(ctx context.Context) error

	RollbackTransaction(ctx context.Context) error
}