package data

import (
	"fmt"

	"gorm.io/gorm"
)

type IRepository interface {
	Create(value interface{}) error
	Save(value interface{}) error
	First(obj interface{}, query string, args ...interface{}) (interface{}, error)
	Find(obj interface{}, conds ...interface{}) (interface{}, error)
}

type RepositoryStruct struct {
	Connection *gorm.DB
}

//Create inserts a new row
func (r *RepositoryStruct) Create(value interface{}) error {
	result := r.Connection.Create(value)
	if result == nil || result.Error != nil {
		return fmt.Errorf("an error ocurred")
	}
	return nil
}

//Save updates a row
func (r *RepositoryStruct) Save(value interface{}) error {
	result := r.Connection.Save(value)
	if result == nil || result.Error != nil {
		return fmt.Errorf("an error ocurred")
	}
	return nil
}

//First select top 1 row
func (r *RepositoryStruct) First(obj interface{}, query string, args ...interface{}) (interface{}, error) {

	result := r.Connection.Where(query, args).First(obj)
	if result.Error != nil && result.Error.Error() == "record not found" {
		return nil, nil
	}
	return obj, result.Error
}

//Find select by interface
func (r *RepositoryStruct) Find(obj interface{}, conds ...interface{}) (interface{}, error) {
	var result *gorm.DB

	if conds == nil {
		result = r.Connection.Find(obj)
	} else {
		result = r.Connection.Find(obj, conds)
	}

	if result.Error != nil {
		return obj, result.Error
	}
	return obj, nil
}
