package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Files struct {
	gorm.Model
	FileID   string `gorm:"primaryKey"`
	Filename string
	FileSize int64
}

var Db *gorm.DB

func Init() {
	var err error
	Db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	Db.AutoMigrate(&Files{})
}
