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

	dbHost := getEnv("DB_HOST", "pg-products")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "products_db")

	retries := getEnvAsInt("DB_CONN_RETRIES", 5)
	retryDelay := getEnvAsInt("DB_CONN_DELAY", 2)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName,
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

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
