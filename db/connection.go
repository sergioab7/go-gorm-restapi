package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost user=fazt password=mysecretpassword dbname=gorm port=5432"

var (
	DB *gorm.DB
)

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("[+] DB conectado correctamente")
	}
}
