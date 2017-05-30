package controller

import (
	"app/model"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"app/helper"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SignUp :Post /user/signup
func SignUp(w http.ResponseWriter, r *http.Request) {
	var _buf, _user model.User
	db, err := apiInit(&_user, r)
	defer db.Close()
	if err != nil {
		returnMessage(w, http.StatusServiceUnavailable, "DB Server down")
		return
	}
	// confirm not exist user_name
	db.First(&_buf, "user_name = ?", _user.UserName)
	if _buf.ID != 0 {
		returnMessage(w, http.StatusBadRequest, "exists post user name")
		return
	}

	password := []byte(_user.Password)
	cost := 10
	hash, _ := bcrypt.GenerateFromPassword(password, cost)
	_user.Password = string(hash)
	db.Create(&_user)

	returnMessage(w, http.StatusOK, "sign up success")
}

// SignUp :Post /user/login
func Login(w http.ResponseWriter, r *http.Request) {
	var _buf, _user model.User
	db, err := apiInit(&_buf, r)
	defer db.Close()
	if err != nil {
		returnMessage(w, http.StatusServiceUnavailable, "DB Server down")
		return
	}
	password := []byte(_buf.Password)
	db.First(&_user, "user_name = ?", _buf.UserName)

	if _user.ID == 0 {
		returnMessage(w, http.StatusBadRequest, "login fault")
		return
	}
	var hash = []byte(_user.Password)
	err = bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		returnMessage(w, http.StatusBadRequest, "login fault")
		return
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &_user)
	var _tokenstring string
	_tokenstring, err = token.SignedString([]byte(helper.SECRET))
	if err != nil {
		returnMessage(w, http.StatusBadRequest, "login fault")
		return
	}
	returnMessage(w, http.StatusOK, "login success")
	var jsonBlob = []byte(`
		{ "token": ` + _tokenstring + `" }
	`)
	w.Write(jsonBlob)
}
