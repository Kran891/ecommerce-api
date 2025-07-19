package config

import (
	"ecommerce-api/logger"
	"ecommerce-api/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConncectDB() {
	godotenv.Load()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"db.mhzjnslrwbepyuiebvoi.supabase.co","postgres","3MdDnbdiHbCkOIwa","postgres","5432","require",
	)
    logger.Info(os.Getenv("DB_HOST"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("❌❌❌ failed to connect database", err)
		// Log the error and panic to stop the application
		panic("❌❌❌ failed to connect database")

	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(models.Cart{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.OrderItem{})
	DB = db
	// Log the successful connection
	logger.Info("✅✅✅ Database connection established successfully")

}
