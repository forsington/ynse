package budget

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServiceImpl_Get(t *testing.T) {

	// repo error
	repo := NewMockYNAB()
	someErr := errors.New("testing-error")
	repo.On("Budgets").Return(nil, someErr)
	srvc := New(repo)
	_, err := srvc.Get()
	assert.EqualError(t, err, someErr.Error())

	// empty set
	repo = NewMockYNAB()
	repo.On("Budgets").Return(nil, nil)
	srvc = New(repo)
	_, err = srvc.Get()
	assert.NoError(t, err)

	// Get Accounts error
	repo = NewMockYNAB()
	someBudgets := []*Budget{{ID: "test", Name: "test"}}
	repo.On("Budgets").Return(someBudgets, nil)
	repo.On("Accounts", "test").Return(nil, someErr)
	srvc = New(repo)
	_, err = srvc.Get()
	assert.EqualError(t, err, someErr.Error())

	// OK
	repo = NewMockYNAB()
	someBudgets = []*Budget{{ID: "test", Name: "test"}}
	repo.On("Budgets").Return(someBudgets, nil)
	someAccounts := []*Account{{ID: "test", Name: "test"}}
	repo.On("Accounts", "test").Return(someAccounts, nil)
	srvc = New(repo)
	_, err = srvc.Get()
	assert.NoError(t, err)
}

func TestServiceImpl_Push(t *testing.T) {
	transactions := []*Transaction{
		&Transaction{
			ID:        "",
			Date:      time.Now(),
			Amount:    0,
			PayeeName: "",
			Memo:      "",
			Cleared:   "",
			Approved:  false,
		},
	}

	// repo error
	repo := NewMockYNAB()
	someErr := errors.New("testing-error")
	repo.On("GetTransactions").Return(nil, someErr)
	srvc := New(repo)
	_, err := srvc.Push("", "", transactions, false)
	assert.EqualError(t, err, someErr.Error())

	// Allow Dupes
	ids := []string{"12345"}
	repo = NewMockYNAB()
	repo.On("GetTransactions").Return(nil, nil)
	repo.On("SendTransactions").Return(ids, nil)
	srvc = New(repo)
	_, err = srvc.Push("", "", transactions, false)
	assert.NoError(t, err)

	// OK
	ids = []string{"12345"}
	repo = NewMockYNAB()
	repo.On("GetTransactions").Return(transactions, nil)
	repo.On("SendTransactions").Return(ids, nil)
	srvc = New(repo)
	_, err = srvc.Push("", "", transactions, false)
	assert.NoError(t, err)
}