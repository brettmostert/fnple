package entries

import (
	"context"
	"fmt"
	"os"

	"github.com/brettmostert/fnple-go/go/components/ledger/internal/common"
	"github.com/jackc/pgx/v4"
)

// // Get all habits from db
// func getAll(ctx *common.AppContext) []*Entry {
// 	var data []*Transaction

// 	err := pgxscan.Select(context.Background(), ctx.Db, &data, `SELECT id, coalation_id, description, status, created_time, modified_time FROM transaction`)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Select failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	return data
// }

func insertEntry(ctx *common.AppContext, trx pgx.Tx, en Entry) Entry {
	var data Entry
	var id int

	err := trx.QueryRow(context.Background(), `INSERT INTO entry
		(transaction_id, description, type, amount, account_id )
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
		en.TransactionId, en.Description, en.Type, en.Amount, en.AccountId).Scan(&id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Insert failed: %v\n", err)
		os.Exit(1)
	}

	data = en
	data.Id = id

	return data
}
