package domain_test

import (
	"getmail/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MustNewBaseGeneratesNewUUID(t *testing.T) {

	base := domain.NewBase()

	assert.NotNil(t, base.ID)
}

func Test_MustNewBaseInformWhenWasCreated(t *testing.T) {

	base := domain.NewBase()

	assert.NotNil(t, base.CreatedAt)
}
