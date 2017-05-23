package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	UserName string `json:user_name`
	Password string `json:password`
}

type Users []User
