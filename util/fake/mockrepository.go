package fake

import (
	"getmail/domain/lists"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
	ListToReturn *lists.List
}

func (r *MockRepository) Create(value interface{}) error {
	args := r.Called(value)
	return args.Error(0)
}

func (r *MockRepository) Save(value interface{}) error {
	args := r.Called(value)
	return args.Error(0)
}

func (r *MockRepository) First(obj interface{}, query string, args ...interface{}) (interface{}, error) {
	// We have to recombine all arguments into a single slice, then expand it
	// out to arguments.
	a := r.Called(append(append([]interface{}{}, obj, query), args...)...)

	return a.Get(0), a.Error(1)
}

func (r *MockRepository) Find(obj interface{}, conds ...interface{}) (interface{}, error) {
	// See First.
	args := r.Called(append(append([]interface{}{}, obj), conds...)...)
	return args.Get(0), args.Error(1)
}
