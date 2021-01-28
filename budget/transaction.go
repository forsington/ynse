package budget

import "time"

// Transaction is the Ynab Representation of a transaction
type Transaction struct {
	ID        string
	AccountID string
	Date      time.Time
	Amount    int64
	PayeeName string
	Memo      string
	Cleared   string
	Approved  bool
}
