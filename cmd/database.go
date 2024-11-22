package cmd

import (
	"composer_vue/backend/app/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/data/database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&model.Composer{},
		&model.PublicType{},
		&model.PublicData{},
		&model.Container{},
		&model.Network{},
	)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
