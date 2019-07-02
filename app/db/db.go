package db

import (
	"github.com/arthurc0102/dcard-popular-post-notify/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Connection *gorm.DB

// Setup database connection
func Setup(dbConfig string) error {
	db, err := gorm.Open("postgres", dbConfig)

	if err != nil {
		return err
	}

	Connection = db
	return nil
}

// Migrate database
func Migrate() {
	Connection.AutoMigrate(&models.Post{})
}
