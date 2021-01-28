package budget

// Account is a bank account
type Account struct {
	Name         string
	ID           string
	Type         string
	Transactions []*Transaction
}
