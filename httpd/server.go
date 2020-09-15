package httpd

//InitializeHttpServer creates new server
func InitializeHttpServer() {

	router := RegisterHTTPHandlers()

	router.Run(":5000")
}
