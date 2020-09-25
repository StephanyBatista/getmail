package handler_test

import (
	"getmail/domain/lists"
	"getmail/util/fake"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRepository struct {
	mock.Mock
	listToReturn *lists.List
}

func (r *MockRepository) Create(value interface{}) error {
	args := r.Called(value)
	return args.Error(0)
}

func (r *MockRepository) First(obj interface{}, query string, args ...interface{}) error {
	// We have to recombine all arguments into a single slice, then expand it
	// out to arguments.
	a := r.Called(append(append([]interface{}{}, obj, query), args...)...)

	// Set the obj arg (if any).
	if r.listToReturn != nil {
		*(obj.(*lists.List)) = *r.listToReturn
	}

	return a.Error(0)
}

func (r *MockRepository) Find(obj interface{}, conds ...interface{}) error {
	// See First.
	args := r.Called(append(append([]interface{}{}, obj), conds...)...)
	return args.Error(0)
}

func (r *MockRepository) Save(value interface{}) error {
	args := r.Called(value)
	return args.Error(0)
}

func Test_PostListMustReturn400WhenNameAlreadyExists(t *testing.T) {
	listToReturn, err := lists.New("list 1")
	require.NoError(t, err)

	repo := &MockRepository{
		listToReturn: listToReturn,
	}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, "name = ?", "list 1").Return(nil)

	var json = []byte(`{"name": "list 1"}`)
	code, _ := fake.NewJsonRequest(repo, "POST", "/list", json)

	assert.Equal(t, 400, code)
}

func Test_PostListMustReturn400WhenNotSendNome(t *testing.T) {
	code, body := fake.NewJsonRequest(nil, "POST", "/list", nil)
	assert.Equal(t, 400, code)
	assert.Contains(t, body, "The Name is required")
}

func Test_PostListMustReturn201WhenSaveANewSubscriber(t *testing.T) {
	repo := &MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, "name = ?", "list 1").Return(nil)
	repo.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		// This is where you verify that it's creating the correct record.
		list := args[0].(*lists.List)
		assert.Equal(t, "list 1", list.Name)
	}).Return(nil)

	var json = []byte(`{"name": "list 1"}`)
	code, _ := fake.NewJsonRequest(repo, "POST", "/list", json)

	assert.Equal(t, 201, code)
}
