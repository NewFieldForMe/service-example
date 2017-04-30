package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Articles []Article
