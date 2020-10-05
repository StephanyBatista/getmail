package handler_test

import (
	"fmt"
	"getmail/domain/subscribers"
	"getmail/util/fake"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetSubscriberMustReturn500WhenErrorOnDatabase(t *testing.T) {
	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("Find", mock.Anything).Return(nil, fmt.Errorf("An error"))

	code, _ := fake.NewJsonRequest(repo, "GET", "/subscriber", nil)

	assert.Equal(t, 500, code)
}

func Test_GetSubscriberMustReturnSubscribers(t *testing.T) {
	newSubscriber, _ := subscribers.New("teste@teste.com.br", "nome")
	subscriberList := [...]subscribers.Subscriber{newSubscriber}
	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("Find", mock.Anything).Return(subscriberList, nil)

	code, _ := fake.NewJsonRequest(repo, "GET", "/subscriber", nil)

	assert.Equal(t, 200, code)
}
