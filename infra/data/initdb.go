package data

import (
	"fmt"
	"getmail/domain/lists"
	"getmail/domain/subscribers"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func getDB() *gorm.DB {
	db, err := gorm.Open(sqlserver.Open(os.Getenv("DB_CONNECTION")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

//InitializeDB initializes connection and execute migrate from database
func InitializeDB() {
	connection := getDB()
	connection.AutoMigrate(&subscribers.Subscriber{}, &lists.List{})
	fmt.Println("Connection SQL: ", connection)

	Repository = &RepositoryStruct{connection: connection}
}
