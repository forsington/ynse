package budget

/*
Copyright © 2021 HAMPUS FORS <h@f0.rs>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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