package handler_test

import (
	"getmail/domain/lists"
	"getmail/util/fake"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_PostListMustReturn400WhenNameAlreadyExists(t *testing.T) {
	listToReturn, err := lists.New("list 1")
	require.NoError(t, err)

	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, "name = ?", "list 1").Return(listToReturn, nil)

	var json = []byte(`{"name": "list 1"}`)
	code, _ := fake.NewJsonRequest(repo, "POST", "/list", json)

	assert.Equal(t, 400, code)
}

func Test_PostListMustReturn400WhenNomeDidNotSend(t *testing.T) {
	code, body := fake.NewJsonRequest(nil, "POST", "/list", nil)
	assert.Equal(t, 400, code)
	assert.Contains(t, body, "The Name is required")
}

func Test_PostListMustReturn201WhenCreateList(t *testing.T) {
	var listNameExpected = "list 1"
	var json = []byte(`{"name": "` + listNameExpected + `"}`)
	repo := &fake.MockRepository{}
	defer repo.AssertExpectations(t)
	repo.On("First", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	repo.On("Create", mock.Anything).Run(func(args mock.Arguments) {
		list := args[0].(lists.List)
		assert.Equal(t, listNameExpected, list.Name)
	}).Return(nil)

	code, _ := fake.NewJsonRequest(repo, "POST", "/list", json)

	assert.Equal(t, 201, code)
}
