package main

import (
	"app/model"
	"flag"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var migrate = flag.Bool("migrate", false, "execute db migrate")
	flag.Parse()

	if err := os.Mkdir("./db", 0777); err != nil {
	}
	db, err := gorm.Open("sqlite3", "./db/gorm.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	if *migrate == true {
		dbMigrate(db)
		return
	}
	fmt.Print("To show how to use 'go run db_migrate.go --help'")
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Article{})
}
