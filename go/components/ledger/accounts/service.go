package accounts

import (
	"context"
	"time"

	"github.com/brettmostert/fnple-go/go/components/ledger/internal/common"
	"github.com/jackc/pgx/v4"
)

func CreateAccount(ctx *common.AppContext, acc NewAccount) Account {
	now := time.Now()

	newAcc := Account{
		CoalationId:  acc.CoalationId,
		Description:  acc.Description,
		CreatedTime:  now,
		ModifiedTime: now,
	}
	trx, _ := ctx.Db.BeginTx(context.Background(), pgx.TxOptions{})
	defer trx.Rollback(context.Background())
	newAcc = insertAccount(ctx, trx, newAcc)
	trx.Commit(context.Background())
	return newAcc
}
