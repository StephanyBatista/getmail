package handler_test

import (
	"getmail/infra/data"
	"getmail/util/fake"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func init() {

	data.Repository = &fake.MockRepository{}
}

func Test_PostListMustReturn400WhenNotSendNome(t *testing.T) {

	code, body := fake.NewJsonRequest("POST", "/list", nil)
	assert.Equal(t, 400, code)
	if !strings.Contains(body, "The Name is required") {
		t.Fail()
	}
}

func Test_PostListMustReturn201WhenSaveANewSubscriber(t *testing.T) {

	var json = []byte(`{"name": "list 1"}`)

	code, _ := fake.NewJsonRequest("POST", "/list", json)

	assert.Equal(t, 201, code)
}
