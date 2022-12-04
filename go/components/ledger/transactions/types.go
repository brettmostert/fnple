package transactions

import (
	"time"

	"github.com/brettmostert/fnple-go/go/components/ledger/transactions/entries"
)

type NewTransaction struct {
	CoalationId string
	Description string
	Entries     []entries.NewEntry
	TxOptions   TxOptions
}

type Transaction struct {
	Id           int
	CoalationId  string
	Description  string
	Status       string
	Entries      []entries.Entry
	CreatedTime  time.Time
	ModifiedTime time.Time
}

type TxOptions struct {
	AutoComplete bool
}
