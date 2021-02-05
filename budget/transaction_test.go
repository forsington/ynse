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
