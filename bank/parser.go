package bank

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
