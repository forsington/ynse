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

// Repo is the budgets repository interface
type Repo interface {
	Budgets() ([]*Budget, error)
	Accounts(budgetID string) ([]*Account, error)
	GetTransactions(budgetID string, accountID string) ([]*Transaction, error)
	SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]string, error)
}
