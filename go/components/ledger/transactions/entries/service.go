package entries

import (
	"github.com/brettmostert/fnple-go/go/components/ledger/internal/common"
	"github.com/jackc/pgx/v4"
)

func CreateEntry(ctx *common.AppContext, trx pgx.Tx, en NewEntry) Entry {
	newEn := Entry{
		TransactionId: en.TransactionId,
		AccountId:     en.AccountId,
		Type:          en.Type,
		Amount:        en.Amount,
		Description:   en.Description,
	}

	// trx, _ := ctx.Db.BeginTx(context.Background(), pgx.TxOptions{})
	// defer trx.Rollback(context.Background())
	newEn = insertEntry(ctx, trx, newEn)
	// trx.Commit(context.Background())
	return newEn
}
