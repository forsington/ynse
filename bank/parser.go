package bank

import (
	"errors"

	"github.com/forsington/ynse/budget"
)

// Parser is a parser of bank transaction statements
type Parser interface {
	ReadFile(b []byte) ([]*budget.Transaction, error)
	Bank() string
}

var (
	// ErrParserNotFound is returned when no parser is found
	ErrParserNotFound = errors.New("parser not found")
)

const (
	// SHB is Svenska Handelsbanken
	SHB = "shb"
	// SEB is Skandinaviska Enskilda Banken
	SEB = "seb"
	// NORDEA is Nordea Bank
	NORDEA = "nordea"
	// SWEDBANK is Swedbank
	SWEDBANK = "swedbank"
)

// Parsers is a wrapper for a map of parsers
type Parsers map[string]Parser

// ImplementedParsers is the map of implemented parsers
var ImplementedParsers Parsers = map[string]Parser{SHB: &Handelsbanken{}}

// Find a bank parsers by bank name
func (p Parsers) Find(bank string) (Parser, error) {
	if _, ok := p[bank]; !ok {
		return nil, ErrParserNotFound
	}
	return p[bank], nil
}
