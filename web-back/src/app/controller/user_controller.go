package controller

import (
	"app/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SignUp :Post /user/signup
func SignUp(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	// confirm not exist user_name
	var _user model.User
	json.Unmarshal(body, &_user)
	var _buf model.User
	db.First(&_buf, "user_name = ?", _user.UserName)
	if _buf.ID != 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var jsonBlob = []byte(`
			{ "code": "400", "message": "exists post user name" }
		`)
		w.Write(jsonBlob)
		return
	}

	password := []byte(_user.Password)
	cost := 10
	hash, _ := bcrypt.GenerateFromPassword(password, cost)

	err = bcrypt.CompareHashAndPassword(hash, password)

	if err == nil {
		println("ok")
	}

	_user.Password = string(hash)
	db.Create(&_user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var jsonBlob = []byte(`
		{ "code": "200", "message": "sign up success" }
	`)
	w.Write(jsonBlob)
}
