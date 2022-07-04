package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv" 
)

func goDotEnvVariable(key string) string { 
 
	// load .env file 
	err := godotenv.Load(".env") 
	  
	if err != nil { 
	  fmt.Print(err) 
	}
	return os.Getenv(key)
}

func Connect() {
	host := goDotEnvVariable("host")
	port := goDotEnvVariable("port")
	user := goDotEnvVariable("user")
	password := goDotEnvVariable("password")
	dbname := goDotEnvVariable("dbname")
	sslname := goDotEnvVariable("sslname")
	pg_con := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("mssql", pg_con)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
