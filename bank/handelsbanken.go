package bank

import (
	"os"

	"github.com/forsington/ynse/budget"
)

type Handelsbanken struct {
}

func (h *Handelsbanken) ReadFile(f *os.File) ([]*budget.Transaction, error) {
	panic("implement me")
}
