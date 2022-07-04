package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/abylaymoldabek/redisExample/handler"
	"github.com/gorilla/mux"
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

func main() {
	port := goDotEnvVariable("port")
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", handler.SaveData)
	router.HandleFunc("/currency/{id}/{*code}", handler.SelectData)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+port, nil)

}
