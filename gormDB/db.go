package gormdb

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	conn := "postgres://postgres:1234@localhost:5432/gogin?sslmode=disable"

	data, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	DB = data
}
