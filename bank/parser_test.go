package bank

import (
	"os"

	"github.com/forsington/ynse/budget"
	"github.com/stretchr/testify/mock"
)

type MockParser struct {
	mock.Mock
}

func (p *MockParser) ReadFile(f *os.File) ([]*budget.Transaction, error) {
	args := p.Called(f)
	return args.Get(0).([]*budget.Transaction), args.Error(1)
}
