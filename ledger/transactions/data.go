package transactions

import (
	"context"
	"fmt"
	"os"

	"github.com/brettmostert/fnple/ledger/internal/common"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

// Get all habits from db
func getAll(ctx *common.AppContext) []*Transaction {
	var data []*Transaction

	err := pgxscan.Select(context.Background(), ctx.Db, &data, `SELECT id, coalation_id, description, status, created_time, modified_time FROM transaction`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Select failed: %v\n", err)
		os.Exit(1)
	}

	return data
}

func insertTransaction(ctx *common.AppContext, trx pgx.Tx, tx Transaction) Transaction {
	var data Transaction
	var id int

	err := trx.QueryRow(context.Background(), `INSERT INTO transaction
		(coalation_id, description, status, created_time, modified_time)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
		tx.CoalationId, tx.Description, tx.Status,
		tx.CreatedTime, tx.ModifiedTime).Scan(&id)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Insert failed: %v\n", err)
		os.Exit(1)
	}

	data = tx
	data.Id = id

	return data
}
