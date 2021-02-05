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