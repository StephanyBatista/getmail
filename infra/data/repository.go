package data

import (
	"fmt"

	"gorm.io/gorm"
)

type IRepository interface {
	Create(value interface{}) error
	Save(value interface{}) error
	First(obj interface{}, query string, args ...interface{}) error
	Find(obj interface{}, conds ...interface{}) error
}

type RepositoryStruct struct {
	connection *gorm.DB
}

var Repository IRepository

//Create inserts a new row
func (r *RepositoryStruct) Create(value interface{}) error {
	result := r.connection.Create(value)
	if result == nil || result.Error != nil {
		return fmt.Errorf("an error ocurred")
	}
	return nil
}

//Save updates a row
func (r *RepositoryStruct) Save(value interface{}) error {
	result := r.connection.Save(value)
	if result == nil || result.Error != nil {
		return fmt.Errorf("an error ocurred")
	}
	return nil
}

//First select top 1 row
func (r *RepositoryStruct) First(obj interface{}, query string, args ...interface{}) error {

	result := r.connection.Where(query, args).First(obj)
	return result.Error
}

//Find select by interface
func (r *RepositoryStruct) Find(obj interface{}, conds ...interface{}) error {
	var result *gorm.DB

	if conds == nil {
		result = r.connection.Find(obj)
	} else {
		result = r.connection.Find(obj, conds)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}
