package importer

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/forsington/ynse/bank"
	"github.com/forsington/ynse/budget"
)

type Importer interface {
	Import(filename, path, bank string) ([]*budget.Transaction, error)
}

type importerImpl struct {
	parsers bank.Parsers
}

var (
	ErrNoFileOrDir = errors.New("filename or dir must be specified")
)

func New(parsers bank.Parsers) Importer {
	return &importerImpl{
		parsers: parsers,
	}
}

func (i *importerImpl) Import(filename, dir, bank string) ([]*budget.Transaction, error) {
	// run file through parser
	parser, err := i.parsers.Find(bank)
	if err != nil {
		return nil, err
	}

	var transactions []*budget.Transaction
	if filename != "" {
		transactions, err = readFile(filename, parser)
		if err != nil {
			return nil, err
		}
	} else if dir != "" {
		transactions, err = readDir(dir, parser)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, ErrNoFileOrDir
	}

	return transactions, nil
}

func readFile(filename string, parser bank.Parser) ([]*budget.Transaction, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return parser.ReadFile(f)
}

func readDir(dir string, parser bank.Parser) ([]*budget.Transaction, error) {
	var transactions []*budget.Transaction

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		f, err := os.Open(dir + file.Name())
		if err != nil {
			return nil, err
		}
		newTransactions, err := parser.ReadFile(f)
		if err != nil {
			return nil, err
		}

		// there might be transactions that existing in multiple files, remove these duplicates
		newTransactions = removeDuplicateTransactions(transactions, newTransactions)
		transactions = append(transactions, newTransactions...)
	}
	return transactions, nil
}

func removeDuplicateTransactions(existing, incoming []*budget.Transaction) []*budget.Transaction {
	var cleanTransactions []*budget.Transaction
	for _, incomingTransaction := range incoming {
		duplicate := false

		for _, existingTransaction := range existing {
			// fuzzy matching for existing transactions to avoid duplicates
			if incomingTransaction.Date == existingTransaction.Date &&
				incomingTransaction.Amount == existingTransaction.Amount &&
				incomingTransaction.PayeeName == existingTransaction.PayeeName {
				duplicate = true
			}
		}
		if !duplicate {
			cleanTransactions = append(cleanTransactions, incomingTransaction)
		}
	}

	return cleanTransactions
}
