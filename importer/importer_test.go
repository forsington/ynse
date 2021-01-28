package importer

import (
	"testing"

	"github.com/forsington/ynse/bank"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestImporter_Import(t *testing.T) {
	i := New(hclog.NewNullLogger(), bank.ImplementedParsers)

	// invalid bank
	_, err := i.Import("test", "test", "")
	assert.EqualError(t, err, bank.ErrParserNotFound.Error())

	_, err = i.Import("test", "test", bank.SEB)
	assert.EqualError(t, err, bank.ErrParserNotFound.Error())

	// no filename or path
	_, err = i.Import("", "", bank.SHB)
	assert.EqualError(t, err, ErrNoFileOrDir.Error())

	// incorrect filename
	_, err = i.Import("incorrect-file-path.csv", "", bank.SHB)
	assert.Error(t, err)

	// incorrect dir
	_, err = i.Import("", "incorrect-dir-path", bank.SHB)
	assert.Error(t, err)
}