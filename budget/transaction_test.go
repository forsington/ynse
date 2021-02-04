package budget

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransaction_AmountPretty(t *testing.T) {
	tran := &Transaction{Amount: 0}
	assert.Equal(t,"0.00 ", tran.AmountPretty(""))

	tran = &Transaction{Amount: 100000}
	assert.Equal(t,"+100.00 ", tran.AmountPretty(""))

	tran = &Transaction{Amount: -10000}
	assert.Equal(t,"-10.00 ", tran.AmountPretty(""))

	tran = &Transaction{Amount: 100000}
	assert.Equal(t,"+100.00 SEK", tran.AmountPretty(CurrencySEK))

	tran = &Transaction{Amount: 100110}
	assert.Equal(t,"+100.11 ", tran.AmountPretty(""))

	tran = &Transaction{Amount: -543210}
	assert.Equal(t,"-543.21 ", tran.AmountPretty(""))

}
