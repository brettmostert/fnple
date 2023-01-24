package accounts

import (
	"context"
	"fmt"
	"os"

	"github.com/brettmostert/fnple/ledger/internal/common"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

// Get all habits from db.
func getAll(ctx *common.AppContext) []*Account {
	var data []*Account

	err := pgxscan.Select(context.Background(), ctx.Db, &data, `SELECT id, coalation_id, description, created_time, modified_time FROM account`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Select failed: %v\n", err)
		os.Exit(1)
	}

	return data
}

func insertAccount(ctx *common.AppContext, trx pgx.Tx, acc Account) Account {
	var data Account

	var id int

	err := trx.QueryRow(context.Background(), `INSERT INTO account
		(coalation_id, description, created_time, modified_time)
		VALUES ($1, $2, $3, $4) RETURNING id;`,
		acc.CoalationId, acc.Description,
		acc.CreatedTime, acc.ModifiedTime).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Insert failed: %v\n", err)
		os.Exit(1)
	}

	data = acc
	data.Id = id

	return data
}
