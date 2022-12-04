package entries

type NewEntry struct {
	TransactionId int
	AccountId     int
	Description   string
	Amount        int
	Type          EntryType
}

type Entry struct {
	Id            int
	TransactionId int
	AccountId     int
	Description   string
	Amount        int
	Type          EntryType
}

type EntryType string

const (
	Debit  EntryType = "dr"
	Credit EntryType = "cr"
)
