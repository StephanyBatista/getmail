package util_test

import (
	"getmail/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MustGenerateNewIdWith8Char(t *testing.T) {

	const numberOfChar = 8

	newID := util.NewID()

	assert.Equal(t, numberOfChar, len(newID))
}

func Test_MustNotGenerateTheSameID(t *testing.T) {

	newID1 := util.NewID()
	newID2 := util.NewID()

	assert.NotEqual(t, newID1, newID2)
}
