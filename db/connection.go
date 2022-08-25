package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSM = "host=localhost user=postgres password=123456 dbname=gorm port=5432"
var db *gorm.DB

func DBConnect() {
	var error error
	db, error = gorm.Open(postgres.Open(DSM), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("db connectada correctamente")
	}
}
