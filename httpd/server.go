package httpd

import "os"

//InitializeHttpServer creates new server
func InitializeHttpServer() {

	router := RegisterHTTPHandlers()

	router.Run(":" + os.Getenv("PORT"))
}
