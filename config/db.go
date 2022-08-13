package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DBConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("akubisa2.db"), &gorm.Config{})
	return db, err
}
