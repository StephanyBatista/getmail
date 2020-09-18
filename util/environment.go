package util

import "os"

//InitializeEnvironmentVariables inits the environment variables
func InitializeEnvironmentVariables() {
	if len(os.Getenv("DB_CONNECTION")) == 0 {
		os.Setenv("DB_CONNECTION", "sqlserver://sa:p4ssw0rd*@localhost:1433?database=getmail")
	}
	if len(os.Getenv("PORT")) == 0 {
		os.Setenv("PORT", "5000")
	}
}
