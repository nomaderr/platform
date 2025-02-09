package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global variable to hold DB connection
var DB *gorm.DB

// ConnectDB initializes the MySQL connection with retry logic
func ConnectDB() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// MySQL DSN connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbName)

	var err error
	maxRetries := 10 // Number of times to retry before giving up

	for retries := 0; retries < maxRetries; retries++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err == nil {
			break // If connection succeeds, exit the loop
		}

		log.Printf("❌ Failed to connect to MySQL (attempt %d/%d): %v\n", retries+1, maxRetries, err)
		time.Sleep(5 * time.Second) // Wait 5 seconds before retrying
	}

	if err != nil {
		log.Fatalf("❌ Could not connect to MySQL after multiple retries: %v", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get SQL database instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("✅ Successfully connected to MySQL!")
}
