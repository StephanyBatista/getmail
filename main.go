package main

import (
	"getmail/httpd"
	"getmail/infra/data"
	"os"
)

func init() {
	if len(os.Getenv("DB_CONNECTION")) == 0 {
		os.Setenv("DB_CONNECTION", "sqlserver://sa:p4ssw0rd*@localhost:1433?database=getmail")
	}
	if len(os.Getenv("PORT")) == 0 {
		os.Setenv("PORT", "5000")
	}
}

func main() {
	data.InitializeDB()
	httpd.InitializeHttpServer()
}
