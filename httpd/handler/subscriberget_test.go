package handler_test

import (
	"fmt"
	"getmail/domain/subscribers"
	"getmail/infra/data"
	"getmail/util/fake"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func init() {

	data.Repository = &fake.MockRepository{}
}

func Test_MustReturn500WhenAnErrorOccurs(t *testing.T) {

	data.Repository = &fake.MockRepository{ReturnError: fmt.Errorf("Error")}

	code, _ := fake.NewJsonRequest("GET", "/subscriber", nil)

	assert.Equal(t, 500, code)
}

func Test_MustReturnSubscribers(t *testing.T) {

	list := []subscribers.Subscriber{
		subscribers.Subscriber{Name: "Jose"},
	}
	data.Repository = &fake.MockRepository{ReturnObj: &list}

	code, _ := fake.NewJsonRequest("GET", "/subscriber", nil)

	assert.Equal(t, 200, code)
}
