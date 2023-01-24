package accounts

import "time"

type NewAccount struct {
	CoalationId string
	Description string
}

type Account struct {
	Id           int
	CoalationId  string
	Description  string
	CreatedTime  time.Time
	ModifiedTime time.Time
}
