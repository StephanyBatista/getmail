package handler_test

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/domain/subscribers"
	"getmail/infra/data"
	"getmail/util/fake"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func init() {

	data.Repository = &fake.MockRepository{}
}

func Test_MustReturn400WhenNotSendEmail(t *testing.T) {

	var json = []byte(`{"name": "teste"}`)

	code, body := fake.NewJsonRequest("POST", "/subscriber", json)
	assert.Equal(t, 400, code)
	if !strings.Contains(body, "The Email is required") {
		t.Fail()
	}
}

func Test_MustReturn400WhenNotSendAValidEmail(t *testing.T) {

	var json = []byte(`{"email": "invalid email"}`)

	code, body := fake.NewJsonRequest("POST", "/subscriber", json)

	assert.Equal(t, 400, code)
	if !strings.Contains(body, "The Email should be a valid email") {
		t.Fail()
	}
}

func Test_MustReturn400WhenNotSendName(t *testing.T) {

	var json = []byte(`{"email": "teste@teste.com.br"}`)

	code, body := fake.NewJsonRequest("POST", "/subscriber", json)

	assert.Equal(t, 400, code)
	if !strings.Contains(body, "The Name is required") {
		t.Fail()
	}
}

func Test_MustReturn201WhenSaveANewSubscriber(t *testing.T) {

	var json = []byte(`{"name": "teste", "email": "teste@teste.com.br"}`)

	code, _ := fake.NewJsonRequest("POST", "/subscriber", json)

	assert.Equal(t, 201, code)
}

func Test_MustSaveANewSubscriberOnListWhenSent(t *testing.T) {

	list, _ := lists.New("Test")
	mockRepository := &fake.MockRepository{ReturnObj: &list}
	data.Repository = mockRepository
	var json = []byte(`{"name": "teste", "email": "teste@teste.com.br", "listId": "` + list.Base.ID.String() + `"}`)

	fake.NewJsonRequest("POST", "/subscriber", json)

	assert.Equal(t, list.Base.ID, mockRepository.ObjSent.(*subscribers.Subscriber).ListID)
}

func Test_MustReturn400WhenListSentNotFound(t *testing.T) {

	mockRepository := &fake.MockRepository{ReturnError: fmt.Errorf("Teste")}
	data.Repository = mockRepository
	var json = []byte(`{"name": "teste", "email": "teste@teste.com.br", "listId": "3rer3-34eer-34e"}`)

	code, body := fake.NewJsonRequest("POST", "/subscriber", json)

	assert.Equal(t, 400, code)
	if !strings.Contains(body, "List not found") {
		t.Fail()
	}
}
