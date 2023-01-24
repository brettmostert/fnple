package entries

// TODO: There should be a credit and debit account id
// double check double entry accounting models
// at the very least the interface should not match the storage
// i.e. JSON could be crId, dbId, and stored as separate entries if needed
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
