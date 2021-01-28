package budget

import (
	"errors"
)

// Service is the Budget service
type Service interface {
	Get() ([]*Budget, error)
	Push(budgetId, accountID string) ([]*Transaction, error)
}

type serviceImpl struct {
	repo Repo
}

// New returns a new Budget service
func New(apiKey string) Service {
	return &serviceImpl{
		repo: NewRepo(apiKey),
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

func (s *serviceImpl) Push(budgetID, accountID string) ([]*Transaction, error) {
	return nil, errors.New("implement me")
}
