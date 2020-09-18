package main

import (
	"getmail/httpd"
	"getmail/infra/data"
	"getmail/util"
)

func init() {
	util.InitializeEnvironmentVariables()
}

func main() {
	data.InitializeDB()
	httpd.InitializeHttpServer()
}
