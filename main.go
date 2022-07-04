package main

import (

	"github.com/abylaymoldabek/redisExample/models"
	"github.com/abylaymoldabek/redisExample/handler"
	"github.com/gorilla/mux"
	
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
    router.HandleFunc("/products/{id}", SaveData)
	router.HandleFunc("/currency/{id}/{*code}", SelectData)
    fmt.Println("Server is listening...")
    http.ListenAndServe(":"+port, nil)

}
