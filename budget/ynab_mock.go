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
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Transaction), args.Error(1)
}

// SendTransactions is a mock implementation
func (m *MockYNAB) SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]string, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}
