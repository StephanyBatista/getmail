package domain_test

import (
	"getmail/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MustNewBaseGeneratesNewIDWith8Char(t *testing.T) {

	const numberOfChar = 8

	base := domain.NewBase()

	assert.Equal(t, numberOfChar, len(base.ID))
}

func Test_MustNewBaseTellsWhenWasCreated(t *testing.T) {

	base := domain.NewBase()

	assert.NotNil(t, base.CreatedAt)
}
