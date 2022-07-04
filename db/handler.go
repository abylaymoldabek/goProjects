package db

import (
	"fmt"
	"net/http"

	"github.com/abylaymoldabek/redisExample/models"
	"github.com/gorilla/mux"
)

func SetData(urlData string) {
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

func SelectData(w http.ResponseWriter, r *http.Request) {
	c := &models.Currency{}
	vars := mux.Vars(r)
	code := vars["code"]
	if err := db.QueryRow(
		"SELECT TITLE, CODE, VALUE_V FROM R_CURRENCY WHERE code=$1",
		code,
	).Scan(&c.Fullname, &c.Title, &c.Description); err != nil {
		return nil, err
	}

	return c, nil
}
