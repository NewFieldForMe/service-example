package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Index :GET
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// ArticleIndex :GET /articles
func ArticleIndex(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var _articles Articles
	db.Find(&_articles)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(_articles); err != nil {
		panic(err)
	}
}

// ArticleShow :GET /articles/{articleId}
func ArticleShow(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	_articleID := vars["articleId"]

	var _article Article
	db.First(&_article, _articleID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(_article); err != nil {
		panic(err)
	}
}

// ArticleCreate :POST /articles
func ArticleCreate(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	// var _art_w = Article{Title: "test", Body: "hoge"}
	// db.AutoMigrate(&Article{})
	// db.Create(&_art_w)

	// db.First(&_art, "Title = ?", "test")

	var _article Article
	// json
	if err := json.Unmarshal(body, &_article); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.Create(&_article)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(_article); err != nil {
		panic(err)
	}
}
