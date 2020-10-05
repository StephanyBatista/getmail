package lists_test

import (
	"getmail/domain/lists"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func Test_MustNewGeneratesNewList(t *testing.T) {

	name := gofakeit.Name()

	list, err := lists.New(name)

	assert.NotNil(t, list)
	assert.Nil(t, err)
}

func Test_MustValidateNameWhenCreateNewList(t *testing.T) {

	const errorExpected = "The Name is required"
	const invalidField = ""

	_, err := lists.New(invalidField)

	assert.Equal(t, errorExpected, err.Error())
}
