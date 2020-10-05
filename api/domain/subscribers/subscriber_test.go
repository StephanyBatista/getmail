package subscribers_test

import (
	"getmail/domain/subscribers"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/brianvoe/gofakeit"
)

func Test_MustNewGeneratesNewSubscriber(t *testing.T) {

	email := gofakeit.Email()
	name := gofakeit.Name()

	subscriber, err := subscribers.New(email, name)

	assert.NotNil(t, subscriber)
	assert.Nil(t, err)
}

func Test_MustValidateEmailWhenCreateNewSubscriber(t *testing.T) {

	const errorExpected = "The Email is required"
	const invalidField = ""

	_, err := subscribers.New(invalidField, gofakeit.Name())

	assert.Equal(t, errorExpected, err.Error())
}

func Test_MustValidateNameWhenCreateNewSubscriber(t *testing.T) {

	const errorExpected = "The Name is required"
	const invalidField = ""

	_, err := subscribers.New(gofakeit.Email(), invalidField)

	assert.Equal(t, errorExpected, err.Error())
}

func Test_MustPutSubscriberOnList(t *testing.T) {

	const listId = "xptddew"
	subscriber, _ := subscribers.New(gofakeit.Email(), gofakeit.Name())

	subscriber.PutOnList(listId)

	assert.Equal(t, listId, subscriber.ListID)
}

func Test_MustValidateListWhenPutSubscriberOnList(t *testing.T) {

	const errorExpected = "The list must be valid to put subscriber in a list"
	const invalidListId = ""
	subscriber, _ := subscribers.New(gofakeit.Email(), gofakeit.Name())

	err := subscriber.PutOnList(invalidListId)

	assert.Equal(t, errorExpected, err.Error())
}
