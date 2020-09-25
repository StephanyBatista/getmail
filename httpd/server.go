package httpd

import (
	"getmail/infra/data"
	"os"
)

//InitializeHttpServer creates new server
func InitializeHttpServer() {

	connection := data.InitializeDB()
	repo := &data.RepositoryStruct{Connection: connection}
	router := RegisterHTTPHandlers(repo)

	router.Run(":" + os.Getenv("PORT"))
}
