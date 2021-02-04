package budget

import (
	"fmt"
	"time"
)

const (
	// AmountFactor is the amount by which the actual currency
	AmountFactor = 1000
	// CurrencySEK is the print string for SEK
	CurrencySEK = "SEK"
)

// Transaction is the Ynab Representation of a transaction
type Transaction struct {
	ID        string
	Date      time.Time
	Amount    int64
	PayeeName string
	Memo      string
	Cleared   string
	Approved  bool
}

// AmountPretty returns a formatted currency string for the amount
func (t *Transaction) AmountPretty(currency string) string {
	actualAmount := float64(t.Amount) / AmountFactor
	if t.Amount > 0 {
		return fmt.Sprintf("+%.2f %s", actualAmount, currency)
	}
	return fmt.Sprintf("%.2f %s", actualAmount, currency)
}