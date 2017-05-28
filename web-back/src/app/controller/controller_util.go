package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func apiInit(out interface{}, r *http.Request) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	json.Unmarshal(body, &out)
	return db, nil
}

func returnMessage(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var jsonBlob = []byte(`
		{ "code": ` + strconv.Itoa(status) + `", "message": "` + message + `" }
	`)
	w.Write(jsonBlob)
}
