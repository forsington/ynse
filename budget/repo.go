package budget

type Repo interface {
	Budgets() ([]*Budget, error)
	Accounts(budgetID string) ([]*Account, error)
	GetTransactions(budgetID string, accountID string) ([]*Transaction, error)
	SendTransactions(budgetID string, accountID string, transactions []*Transaction) ([]string, error)
}
