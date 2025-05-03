package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ramiroschettino/Go-Market-Microservices/orders-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Obtenemos todas las variables de entorno primero
	dbConfig := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_PORT":     os.Getenv("DB_PORT"),
	}

	// Establecemos valores por defecto si no están en las variables de entorno
	if dbConfig["DB_HOST"] == "" {
		dbConfig["DB_HOST"] = "pg-products"
	}
	if dbConfig["DB_USER"] == "" {
		dbConfig["DB_USER"] = "postgres"
	}
	if dbConfig["DB_PASSWORD"] == "" {
		dbConfig["DB_PASSWORD"] = "postgres"
	}
	if dbConfig["DB_NAME"] == "" {
		dbConfig["DB_NAME"] = "products_db"
	}
	if dbConfig["DB_PORT"] == "" {
		dbConfig["DB_PORT"] = "5432"
	}

	// Configuración de reintentos (con valores por defecto)
	retries := 5
	if val := os.Getenv("DB_CONN_RETRIES"); val != "" {
		if r, err := strconv.Atoi(val); err == nil {
			retries = r
		}
	}

	retryDelay := 2
	if val := os.Getenv("DB_CONN_DELAY"); val != "" {
		if rd, err := strconv.Atoi(val); err == nil {
			retryDelay = rd
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbConfig["DB_HOST"],
		dbConfig["DB_USER"],
		dbConfig["DB_PASSWORD"],
		dbConfig["DB_NAME"],
		dbConfig["DB_PORT"],
	)

	var err error
	var attempts int

	for attempts = 1; attempts <= retries; attempts++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			if migrateErr := DB.AutoMigrate(&domain.Order{}); migrateErr != nil {
				log.Fatalf("❌ Error en AutoMigrate de Order: %v", migrateErr)
			}

			log.Println("✅ Conectado a PostgreSQL con GORM y migración completada")
			return
		}

		log.Printf("⌛ Intento %d/%d: Error conectando a PostgreSQL: %v", attempts, retries, err)
		if attempts < retries {
			time.Sleep(time.Duration(retryDelay) * time.Second)
		}
	}

	log.Fatalf("❌ No se pudo conectar a PostgreSQL después de %d intentos: %v", retries, err)
}
