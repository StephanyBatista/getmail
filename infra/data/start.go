package data

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/domain/subscribers"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//Connection return a connection open from database
var connection *gorm.DB

func openDB() {
	connectionString := "sqlserver://sa:p4ssw0rd*@localhost:1433?database=getmail"
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	connection = db
}

//InitializeDB initializes connection and execute migrate from database
func InitializeDB() {
	openDB()
	connection.AutoMigrate(&subscribers.Subscriber{}, &lists.List{})
	fmt.Println("Connection SQL: ", connection)

	Repository = &RepositoryStruct{connection: connection}
}
