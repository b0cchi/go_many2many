package main

import (
	"fmt"
	"many2many/model"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

const DATABASE_NAME = "sample.db"

func main() {
	isInitRequired := true
	if _, err := os.Stat(filepath.Base(fmt.Sprintf("./%s", DATABASE_NAME))); err == nil {
		isInitRequired = false
	}
	db, err := gorm.Open(sqlite.Open(DATABASE_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if isInitRequired {
		initDb(db)
	}
}

func initDb(db *gorm.DB) {
	fmt.Println("-- exec init database --")
	db.AutoMigrate(&model.User{}, &model.Dept{})
}
