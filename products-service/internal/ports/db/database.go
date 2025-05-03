package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Obtenemos valores de entorno o usamos defaults
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "pg-products" // Valor por defecto
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres" // Valor por defecto
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		dbPass = "postgres" // Valor por defecto
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "products_db" // Valor por defecto
	}

	// Construimos el DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		dbHost, dbUser, dbPass, dbName,
	)

	// Intentamos conectar
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a PostgreSQL: %v", err)
	}

	log.Println("✅ Conexión exitosa con PostgreSQL")
}
