package handler

package main

import (
	"fmt"
	"github.com/abylaymoldabek/redisExample/models"
	"github.com/abylaymoldabek/redisExample/db"

	"github.com/gorilla/mux"
	"time"
)

func SetData(urlData) {
	respData, err := RequestData(urlData)
	if err != nil {
		fmt.Println(err)
	}
	sqlInsert := `
	INSERT INTO R_CURRENCY (TITLE, CODE, VALUE_V, A_DATE)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	for _, v := range respData.Item {
		_, err = db.Exec(sqlInsert, v.Fullname, v.Title, v.Description, respData.Date)
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}
}


func SaveData(w http.ResponseWriter, r *http.Request) {
	url := goDotEnvVariable("urlData")
	vars := mux.Vars(r)
	id := vars["id"]
	request := fmt.Sprintf("%sfdate=%v", url, id)
	go SetData(request)
	time.Sleep(5 * time.Second)
	
}