package database

// import libraries
import (
	"example/internal/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// set the global database variable
var (
	DBConn *gorm.DB
)

// a function that connected to the sqlite database with an error handler & automigrate the models
func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DBConn.AutoMigrate(&model.Book{})
	DBConn.AutoMigrate(&model.Author{})
	DBConn.AutoMigrate(&model.Genre{})

	fmt.Println("Database Migrated")
}
