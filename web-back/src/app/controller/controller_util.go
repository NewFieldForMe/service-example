package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func apiInit(out interface{}, r *http.Request) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	if out != nil {
		var body []byte
		body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		json.Unmarshal(body, &out)
	}
	return db, err
}

func setMessageResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	jsondata := messageJSON{}
	jsondata.Code = status
	jsondata.Message = message
	setJSONResponse(w, jsondata)
}

func setJSONResponse(w http.ResponseWriter, i interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(i); err != nil {
		return err
	}
	return nil
}
