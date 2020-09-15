package fake

import "reflect"

type MockRepository struct {
	ReturnError error
	ObjSent     interface{}
	ReturnObj   interface{}
}

func (m *MockRepository) Create(obj interface{}) error {
	if m.ReturnError != nil {
		return m.ReturnError
	}
	m.ObjSent = obj
	return nil
}

func (m *MockRepository) First(obj interface{}, query string, args ...interface{}) error {
	if m.ReturnError != nil {
		return m.ReturnError
	}

	v := reflect.ValueOf(m.ReturnObj).Elem()
	if v.CanSet() {
		v.Set(reflect.ValueOf(obj))
	}

	return nil
}

func (m *MockRepository) Find(obj interface{}, conds ...interface{}) error {
	if m.ReturnError != nil {
		return m.ReturnError
	}

	v := reflect.ValueOf(&m.ReturnObj).Elem()
	if v.CanSet() {
		v.Set(reflect.ValueOf(obj))
	}

	return nil
}
