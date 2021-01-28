package importer

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/forsington/ynse/bank"
	"github.com/forsington/ynse/budget"
	"github.com/hashicorp/go-hclog"
)

type Importer interface {
	Import(filename, path, bank string) ([]*budget.Transaction, error)
}

type importerImpl struct {
	logger  hclog.Logger
	parsers bank.Parsers
}

var (
	ErrNoFileOrDir = errors.New("filename or dir must be specified")
)

func New(logger hclog.Logger, parsers bank.Parsers) Importer {
	return &importerImpl{
		logger:  logger,
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

	// clean up transactions

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
		f, err := os.Open(file.Name())
		if err != nil {
			return nil, err
		}
		newTransactions, err := parser.ReadFile(f)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, newTransactions...)
	}
	return transactions, nil
}
