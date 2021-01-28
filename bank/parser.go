package bank

import (
	"errors"
	"os"

	"github.com/forsington/ynse/budget"
)

type Parser interface {
	ReadFile(f *os.File) ([]*budget.Transaction, error)
}

var (
	ErrParserNotFound = errors.New("parser not found")
)
const (
	SHB      = "shb"
	SEB      = "seb"
	NORDEA   = "nordea"
	SWEDBANK = "swedbank"
)

type Parsers map[string]Parser

var ImplementedParsers Parsers = map[string]Parser{SHB: &Handelsbanken{}}

func (p Parsers) Find(bank string) (Parser, error) {
	if _, ok := p[bank]; !ok {
		return nil, ErrParserNotFound
	}
	return p[bank], nil
}
