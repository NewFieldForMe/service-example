package controller

import (
	"app/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"app/helper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ArticleIndex :GET /articles
func ArticleIndex(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var _articles model.Articles
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

	var _article model.Article
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

	// json
	var _buf model.ArticlePostJSON
	json.Unmarshal(body, &_buf)
	var _article model.Article
	_article.Title = _buf.Title
	_article.Body = _buf.Body
	db.Create(&_article)

	// Create temp file
	var _tmpFileName = fmt.Sprint(_article.ID) + "_encode_and_decord.jpg"
	filestrings := strings.Split(_buf.FileData, ",")
	data, err := base64.StdEncoding.DecodeString(filestrings[1]) //[]byte
	if err != nil {
		log.Fatal(err)
	}
	tmp, _ := os.Create(_tmpFileName)
	tmp.Write(data)
	tmp.Close()
	file, _ := os.Open(_tmpFileName)

	// Save AWS S3
	var key = "/article_images/" + fmt.Sprint(_article.ID) + ".jpg"
	cre := credentials.NewStaticCredentials(
		helper.AWS_ACCESS_KEY,
		helper.AWS_SECRET_KEY,
		"")

	cli := s3.New(session.New(), &aws.Config{
		Credentials: cre,
		Region:      aws.String(helper.AWS_S3_REGION),
	})

	cli.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(helper.AWS_S3_BUCKET),
		Key:    aws.String(key),
	})

	s3Uploader := s3manager.NewUploaderWithClient(cli)
	input := &s3manager.UploadInput{
		Bucket: aws.String(helper.AWS_S3_BUCKET),
		Key:    aws.String(key),
		ACL:    aws.String("public-read"),
		Body:   file,
	}
	result, err := s3Uploader.Upload(input)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	if err := os.Remove(_tmpFileName); err != nil {
		fmt.Println(err)
	}

	_article.ImageUrl = result.Location
	db.Save(&_article)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(_article); err != nil {
		panic(err)
	}
}
