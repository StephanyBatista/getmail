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

//InitializeDB initializes Connection and execute migrate from database
func InitializeDB() *gorm.DB {
	connection := getDB()
	connection.AutoMigrate(&subscribers.Subscriber{}, &lists.List{})
	fmt.Println("Connection SQL: ", connection)

	return connection
}
