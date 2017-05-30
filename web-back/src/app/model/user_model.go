package model

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	UserName           string `json:username`
	Password           string `json:password`
	jwt.StandardClaims `gorm:"-"`
}

type TokenUser struct {
	User
}

type Users []User
