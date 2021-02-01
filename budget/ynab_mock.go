package budget

import "github.com/stretchr/testify/mock"

type MockYNAB struct {
	mock.Mock
}

func NewMockYNAB() *MockYNAB {
	return &MockYNAB{}
}

func (m *MockYNAB) Budgets() ([]*Budget, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Budget), args.Error(1)
}

func (m *MockYNAB) Accounts(budgetID string) ([]*Account, error) {
	args := m.Called(budgetID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Account), args.Error(1)
}

func (m *MockYNAB) GetTransactions(budgetID string, accountID string) ([]*Transaction, error) {
	panic("implement me")
}

func (m *MockYNAB) SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]string, error) {
	panic("implement me")
}
