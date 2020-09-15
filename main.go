package main

import (
	"getmail/httpd"
	"getmail/infra/data"
)

func init() {
	data.InitializeDB()
	httpd.InitializeHttpServer()
}

func main() {

}
