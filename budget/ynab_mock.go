package budget

import "github.com/stretchr/testify/mock"

// MockYNAB is a mock implementation
type MockYNAB struct {
	mock.Mock
}

// NewMockYNAB is a mock implementation
func NewMockYNAB() *MockYNAB {
	return &MockYNAB{}
}

// Budgets is a mock implementation
func (m *MockYNAB) Budgets() ([]*Budget, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Budget), args.Error(1)
}

// Accounts is a mock implementation
func (m *MockYNAB) Accounts(budgetID string) ([]*Account, error) {
	args := m.Called(budgetID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Account), args.Error(1)
}

// GetTransactions is a mock implementation
func (m *MockYNAB) GetTransactions(budgetID string, accountID string) ([]*Transaction, error) {
	panic("implement me")
}

// SendTransactions is a mock implementation
func (m *MockYNAB) SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]string, error) {
	panic("implement me")
}
