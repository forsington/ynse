package bank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImporter_Import(t *testing.T) {
	i := New(ImplementedParsers)

	// invalid bank
	_, err := i.Import("test", "test", "")
	assert.EqualError(t, err, ErrParserNotFound.Error())

	_, err = i.Import("test", "test", SEB)
	assert.EqualError(t, err, ErrParserNotFound.Error())

	// no filename or path
	_, err = i.Import("", "", SHB)
	assert.EqualError(t, err, ErrNoFileOrDir.Error())

	// incorrect filename
	_, err = i.Import("incorrect-file-path.csv", "", SHB)
	assert.Error(t, err)

	// incorrect dir
	_, err = i.Import("", "incorrect-dir-path", SHB)
	assert.Error(t, err)
}