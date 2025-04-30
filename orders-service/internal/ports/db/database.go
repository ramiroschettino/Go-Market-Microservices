package db

import (
	"log"

	"github.com/ramiroschettino/Go-Market-Microservices/orders-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=products_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a la base de datos: %v", err)
	}

	if err := DB.AutoMigrate(&domain.Order{}); err != nil {
		log.Fatalf("❌ Error en AutoMigrate de Order: %v", err)
	}

	log.Println("✅ Conectado a PostgreSQL con GORM y migración completada")
}
