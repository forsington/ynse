package importer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/forsington/ynse/bank"
	"github.com/forsington/ynse/budget"
)

// Importer is an importer of bank transaction statements
type Importer interface {
	Import(filename, path, bank string) ([]*budget.Transaction, error)
}

type importerImpl struct {
	parsers bank.Parsers
}

var (
	// ErrNoFileOrDir is returned when no filename or dir is specified
	ErrNoFileOrDir = errors.New("filename or dir must be specified")
)

// New returns a new Importer
func New(parsers bank.Parsers) Importer {
	return &importerImpl{
		parsers: parsers,
	}
}

// Import reads a file or dir and returns the transactions found in any files
func (i *importerImpl) Import(filename, dir, bank string) ([]*budget.Transaction, error) {
	// run file through parser
	parser, err := i.parsers.Find(bank)
	if err != nil {
		return nil, err
	}

	files := []string{}
	if filename != "" {
		files = append(files, filename)
	} else if dir != "" {
		filenames, err := readDir(dir, parser)
		if err != nil {
			return nil, err
		}
		files = append(files, filenames...)
	} else {
		return nil, ErrNoFileOrDir
	}

	transactions, err := readFiles(files, parser)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func readFiles(files []string, parser bank.Parser) ([]*budget.Transaction, error) {
	var transactions []*budget.Transaction

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		fmt.Printf("parsing %s file: %s", parser.Bank(), f.Name())
		b, err := ioutil.ReadFile(f.Name())
		if err != nil {
			return nil, err
		}

		newTransactions, err := parser.ReadFile(b)
		if err != nil {
			return nil, err
		}

		// there might be transactions that existing in multiple files, remove these duplicates
		newTransactions = removeDuplicateTransactions(transactions, newTransactions)
		transactions = append(transactions, newTransactions...)
	}
	return transactions, nil
}

func readDir(dir string, parser bank.Parser) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filenames := []string{}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		filenames = append(filenames, dir+file.Name())
	}
	return filenames, nil
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
