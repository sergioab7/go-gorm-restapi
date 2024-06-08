package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Está en local, si no,

var DSN string

var (
	DB *gorm.DB
)

func DBConnection() {
	CargarENV()
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("[+] DB conectado correctamente")
	}

}

func CargarENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[-] Error al cargar .env archivo")
	}

	DSN = os.Getenv("DB_BASE")
	if DSN == "" {
		fmt.Println("[-] La variable DB_BASE no está establecida")
	}

}
