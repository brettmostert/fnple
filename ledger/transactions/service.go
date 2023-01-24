package transactions

import (
	"context"
	"log"
	"time"

	"github.com/brettmostert/fnple/ledger/accounts"
	"github.com/brettmostert/fnple/ledger/internal/common"
	"github.com/brettmostert/fnple/ledger/transactions/entries"
	"github.com/jackc/pgx/v4"
)

// TODO: refactor entries and 0 logic.
func CreateTransaction(ctx *common.AppContext, tx NewTransaction) Transaction {
	now := time.Now()
	status := "reserved"

	if tx.TxOptions.AutoComplete {
		status = "complete"
	}

	newTx := Transaction{
		CoalationId:  tx.CoalationId,
		Description:  tx.Description,
		Status:       status,
		CreatedTime:  now,
		ModifiedTime: now,
	}

	trx, _ := ctx.Db.BeginTx(context.Background(), pgx.TxOptions{})
	defer trx.Rollback(context.Background())

	newTx = insertTransaction(ctx, trx, newTx)
	// whoops := false
	// for each entry insert it
	newEntries := []entries.Entry{}

	for _, entry := range tx.Entries {
		entry.TransactionId = newTx.Id
		newEntries = append(newEntries, entries.CreateEntry(ctx, trx, entry))
		currentBalance, newBalance := accounts.UpdateBalance(ctx, trx, entry)
		log.Printf("Current Balance: %v, New Balance: %v", currentBalance, newBalance)
	}
	// if whoops {
	// 	log.Println("NOT ENOUGH MONEY")
	// 	trx.Rollback(context.Background())
	// 	return Transaction{}
	// }
	newTx.Entries = newEntries
	trx.Commit(context.Background())

	return newTx
}
