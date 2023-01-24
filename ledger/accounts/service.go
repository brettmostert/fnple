package accounts

import (
	"context"
	"log"
	"time"

	"github.com/brettmostert/fnple/ledger/internal/common"
	"github.com/brettmostert/fnple/ledger/transactions/entries"

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

func UpdateBalance(ctx *common.AppContext, trx pgx.Tx, entry entries.NewEntry) (int, int) {
	// Ensure there is enough money to fulfill the transaction
	var currentBalance int
	var newBalance int
	trx.QueryRow(context.Background(), `SELECT balance from Account WHERE Id = $1`,
		entry.AccountId).Scan(&currentBalance)
	if entry.Type == entries.Debit {
		if currentBalance-entry.Amount < 0 {
			log.Printf("Not enough money, %v, %v - %v = %v", entry.AccountId, currentBalance, entry.Amount, currentBalance-entry.Amount)
			//TODO: this is not good at all fix it
			trx.Rollback(context.Background())
			panic("Not enough money")
		}

		trx.Exec(context.Background(), `UPDATE Account SET balance=balance-$1 WHERE Id = $2`,
			entry.Amount, entry.AccountId)

	} else {
		// Credit account
		trx.Exec(context.Background(), `UPDATE Account SET balance=balance+$1 WHERE Id = $2`,
			entry.Amount, entry.AccountId)
	}

	trx.QueryRow(context.Background(), `SELECT balance from Account WHERE Id = $1`,
		entry.AccountId).Scan(&newBalance)

	return currentBalance, newBalance
}
