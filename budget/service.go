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

// Service is the Budget service
type Service interface {
	Get() ([]*Budget, error)
	Push(budgetID, accountID string, transactions []*Transaction, allowDuplicates bool) ([]string, error)
}

type serviceImpl struct {
	repo Repo
}

// New returns a new Budget service
func New(repo Repo) Service {
	return &serviceImpl{
		repo: repo,
	}
}

func (s *serviceImpl) Get() ([]*Budget, error) {
	budgets, err := s.repo.Budgets()
	if err != nil {
		return nil, err
	}

	for _, budget := range budgets {
		budget.Accounts, err = s.repo.Accounts(budget.ID)
		if err != nil {
			return nil, err
		}
	}
	return budgets, nil
}

func (s *serviceImpl) Push(budgetID, accountID string, transactions []*Transaction, allowDuplicates bool) ([]string, error) {
	if !allowDuplicates {
		var err error
		transactions, err = s.removeDuplicateTransactions(budgetID, accountID, transactions)
		if err != nil {
			return nil, err
		}
	}
	return s.repo.SendTransactions(budgetID, accountID, transactions)
}

func (s *serviceImpl) removeDuplicateTransactions(budgetID, accountID string, importing []*Transaction) ([]*Transaction, error) {
	existing, err := s.repo.GetTransactions(budgetID, accountID)
	if err != nil {
		return nil, err
	}

	var trans []*Transaction
	for _, newTransaction := range importing {
		alreadyImported := false
		for _, oldTransaction := range existing {
			// fuzzy matching for existing transactions to avoid duplicates
			if oldTransaction.Date == newTransaction.Date &&
				oldTransaction.Amount == newTransaction.Amount &&
				oldTransaction.PayeeName == newTransaction.PayeeName {
				alreadyImported = true
			}
		}
		if !alreadyImported {
			trans = append(trans, newTransaction)
		}
	}
	return trans, nil
}
