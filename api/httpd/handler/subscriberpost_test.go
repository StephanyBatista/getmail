package handler_test

import (
	"getmail/domain/lists"
	"getmail/domain/subscribers"
	"getmail/util/fake"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PostSubscriberMustReturn409WhenEmailAlreadyExists(t *testing.T) {
	const email = "teste@teste.com.br"

	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, "email = ?", email).Return(&subscribers.Subscriber{}, nil)

	var json = []byte(`{"email": "` + email + `"}`)
	code, _ := fake.NewJsonRequest(repo, "POST", "/subscriber", json)

	assert.Equal(t, 409, code)
}

func Test_PostSubscriberMustReturn400WhenSubscriberListNotExist(t *testing.T) {
	const email = "teste@teste.com.br"
	const listID = "XPTO"
	const name = "José Eduardo"
	const jsonString = `{"email": "` + email + `", "name": "` + name + `", "listid": "` + listID + `"}`

	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, "email = ?", email).Return(nil, nil)
	repo.On("First", mock.Anything, "ID = ?", listID).Return(nil, nil)

	var json = []byte(jsonString)
	code, body := fake.NewJsonRequest(repo, "POST", "/subscriber", json)

	assert.Equal(t, 400, code)
	assert.Contains(t, body, "List not exist")
}

func Test_PostSubscriberMustReturn201WhenCreateSubscriber(t *testing.T) {
	subscriberList, _ := lists.New("List XPTO")
	const email = "teste@teste.com.br"
	const name = "José Eduardo"
	var jsonString = `{"email": "` + email + `", "name": "` + name + `", "listid": "` + subscriberList.Base.ID + `"}`

	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, "email = ?", email).Return(nil, nil)
	repo.On("First", mock.Anything, "ID = ?", subscriberList.Base.ID).Return(&subscriberList, nil)
	repo.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		model := args[0].(subscribers.Subscriber)
		assert.Equal(t, subscriberList.Base.ID, model.ListID)
	}).Return(nil)

	var json = []byte(jsonString)
	code, _ := fake.NewJsonRequest(repo, "POST", "/subscriber", json)

	assert.Equal(t, 201, code)
}
