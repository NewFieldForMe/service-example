package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct {
	gorm.Model
	Title string    `json:"title"`
	Body  string    `json:"body"`
	Due   time.Time `json:"due"`
}

type Articles []Article
