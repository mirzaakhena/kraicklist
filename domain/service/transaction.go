package service

import (
	"challenge.haraj.com.sa/kraicklist/domain/repository"
	"context"
)

// ExecuteTransaction is helper function that simplify the transaction execution handling
func ExecuteTransaction(ctx context.Context, trx repository.TransactionRepo, trxFunc func(dbCtx context.Context) error) (err error) {
	dbCtx, err := trx.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = trx.RollbackTransaction(dbCtx)
			panic(p)

		} else if err != nil {
			err = trx.RollbackTransaction(dbCtx)

		} else {
			err = trx.CommitTransaction(dbCtx)

		}
	}()

	err = trxFunc(dbCtx)
	return err
}

