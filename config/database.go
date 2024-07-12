package config

import (
	"fmt"
	"tokyokitten/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnetion() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
