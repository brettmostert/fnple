package transactions

import (
	"context"
	"log"
	"time"

	"github.com/brettmostert/fnple-go/go/components/ledger/internal/common"
	"github.com/brettmostert/fnple-go/go/components/ledger/transactions/entries"
	"github.com/jackc/pgx/v4"
)

// TODO: refactor entries and 0 logic
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
	whoops := false
	// for each entry insert it
	newEntries := []entries.Entry{}
	for _, entry := range tx.Entries {
		entry.TransactionId = newTx.Id
		newEntries = append(newEntries, entries.CreateEntry(ctx, trx, entry))
		// TODO: move to create entry
		if entry.Type == entries.Debit {
			var currentBalance int
			trx.QueryRow(context.Background(), `SELECT balance from Account
		WHERE Id = $1`,
				entry.AccountId).Scan(&currentBalance)
			log.Printf("%v balance (%v)", currentBalance, entry.AccountId)
			if currentBalance-entry.Amount < 0 {
				log.Printf("(%v)", currentBalance-entry.Amount)
				whoops = true
				break
			}
			// TODO: move to create entry
			trx.Exec(context.Background(), `UPDATE Account
		SET balance=balance-$1 WHERE Id = $2`,
				entry.Amount, entry.AccountId)

		} else {
			trx.Exec(context.Background(), `UPDATE Account
		SET balance=balance+$1 WHERE Id = $2`,
				entry.Amount, entry.AccountId)
		}

	}
	if whoops {
		log.Println("NOT ENOUGH MONEY")
		trx.Rollback(context.Background())
		return Transaction{}
	}
	newTx.Entries = newEntries
	trx.Commit(context.Background())
	return newTx
}
