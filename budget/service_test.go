package budget

import (
	"errors"
	"testing"

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
	someBudgets := []*Budget{&Budget{ID: "test", Name: "test"}}
	repo.On("Budgets").Return(someBudgets, nil)
	repo.On("Accounts", "test").Return(nil, someErr)
	srvc = New(repo)
	_, err = srvc.Get()
	assert.EqualError(t, err, someErr.Error())

	// OK
	repo = NewMockYNAB()
	someBudgets = []*Budget{&Budget{ID: "test", Name: "test"}}
	repo.On("Budgets").Return(someBudgets, nil)
	someAccounts := []*Account{&Account{ID: "test", Name: "test"}}
	repo.On("Accounts", "test").Return(someAccounts, nil)
	srvc = New(repo)
	_, err = srvc.Get()
	assert.NoError(t, err)
}
