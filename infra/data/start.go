package data

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/domain/subscribers"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//Connection return a connection open from database
var connection *gorm.DB

func openDB() {
	db, err := gorm.Open(sqlserver.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
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
