package budget

import (
	yn "go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
	"go.bmvs.io/ynab/api/transaction"
)

// Ynab is the YNAB Repo implementation
type Ynab struct {
	client yn.ClientServicer
}

// NewRepo returns a new YNAB Repo
func NewRepo(accessToken string) Repo {
	client := yn.NewClient(accessToken)

	return &Ynab{
		client: client,
	}
}

// Budgets returns a slice of budgets
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

// Accounts returns a slice of Accounts for a given budget
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

// GetTransactions returns all transactions from a given account
func (y *Ynab) GetTransactions(budgetID string, accountID string) ([]*Transaction, error) {
	var transactions []*Transaction

	filter := transaction.Filter{}

	res, err := y.client.Transaction().GetTransactionsByAccount(budgetID, accountID, &filter)

	if err != nil {
		return nil, err
	}

	for _, tran := range res {
		t := &Transaction{
			Date:      tran.Date.Time,
			Amount:    tran.Amount,
			PayeeName: *tran.PayeeName,
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// SendTransactions sends a slice of transactions to YNAB
func (y *Ynab) SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]string, error) {
	if len(transactions) == 0 {
		return []string{}, nil
	}

	var trans []transaction.PayloadTransaction

	for _, tran := range transactions {
		date := api.Date{
			Time: tran.Date,
		}
		payee := tran.PayeeName
		memo := tran.Memo
		trans = append(trans, transaction.PayloadTransaction{
			AccountID: accountID,
			Date:      date,
			Amount:    tran.Amount,
			PayeeName: &payee,
			Memo:      &memo,
			Cleared:   transaction.ClearingStatusUncleared,
			Approved:  tran.Approved,
		})
	}
	res, err := y.client.Transaction().CreateTransactions(budgetID, trans)

	if err != nil {
		return nil, err
	}

	return res.TransactionIDs, nil
}
