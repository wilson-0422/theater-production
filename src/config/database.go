package config

import (
	"theater-production/src/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open(AppConfig.DBPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(
		&models.User{},
		&models.Script{},
		&models.Actor{},
		&models.ActorSchedule{},
		&models.Prop{},
		&models.PropRequisition{},
		&models.Theater{},
		&models.TheaterSchedule{},
		&models.Tour{},
	)
}

func GetDB() *gorm.DB {
	return db
}
