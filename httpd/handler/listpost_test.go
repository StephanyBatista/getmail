package handler_test

import (
	"getmail/domain/lists"
	"getmail/infra/data"
	"getmail/util/fake"
	"testing"

	"github.com/stretchr/testify/mock"
	"gopkg.in/go-playground/assert.v1"
)

type MockRepository struct {
	mock.Mock
}

func (r *MockRepository) Create(value interface{}) error {
	args := r.Called(value)
	return args.Error(0)
}

func (r *MockRepository) First(obj interface{}, query string, args ...interface{}) error {
	a := r.Called(obj, query, args)
	return a.Error(0)
}

func (r *MockRepository) Find(obj interface{}, conds ...interface{}) error {
	args := r.Called(obj, conds)
	return args.Error(0)
}

func (r *MockRepository) Save(value interface{}) error {
	args := r.Called(value)
	return args.Error(0)
}

func Test_PostListMustReturn400WhenNameAlreadyExists(t *testing.T) {

	const name string = "List 1"
	var json = []byte(`{"name": "` + name + `"}`)
	mock := &MockRepository{}
	listToReturn, _ := lists.New(name)
	data.Repository = mock

	mock.On("First", listToReturn, "name = ?", name).Return(nil)

	code, _ := fake.NewJsonRequest("POST", "/list", json)
	assert.Equal(t, 400, code)
}

// func Test_PostListMustReturn400WhenNotSendNome(t *testing.T) {

// 	code, body := fake.NewJsonRequest("POST", "/list", nil)
// 	assert.Equal(t, 400, code)
// 	if !strings.Contains(body, "The Name is required") {
// 		t.Fail()
// 	}
// }

// func Test_PostListMustReturn201WhenSaveANewSubscriber(t *testing.T) {

// 	var json = []byte(`{"name": "list 1"}`)

// 	code, _ := fake.NewJsonRequest("POST", "/list", json)

// 	assert.Equal(t, 201, code)
// }
