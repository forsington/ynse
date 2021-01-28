package budget

import (
	yn "go.bmvs.io/ynab"
)

type Ynab struct {
	client yn.ClientServicer
}

func NewRepo(apiKey string) Repo {
	client := yn.NewClient(apiKey)

	return &Ynab{
		client: client,
	}
}

func (y *Ynab) Budgets() ([]*Budget, error) {
	result, err := y.client.Budget().GetBudgets()
	if err != nil {
		return nil, err
	}

	var budgets []*Budget
	for _, b := range result {
		budgets = append(budgets, &Budget{
			Name: b.Name,
			ID:   b.ID,
		})
	}

	return budgets, nil
}

func (y *Ynab) Accounts(budgetID string) ([]*Account, error) {
	result, err := y.client.Account().GetAccounts(budgetID)
	if err != nil {
		return nil, err
	}

	var accounts []*Account
	for _, a := range result {
		accounts = append(accounts, &Account{
			Name: a.Name,
			ID:   a.ID,
			Type: string(a.Type),
		})
	}
	return accounts, nil
}

func (y *Ynab) GetTransactions(budgetID string, accountID string) ([]*Transaction, error) {
	panic("implement me")
}

func (y *Ynab) SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]*Transaction, error) {
	panic("implement me")
}
