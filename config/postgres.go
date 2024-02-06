package config

import (
	"fmt"
	"os"

	"github.com/mnzsss/go-jobs/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Get environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Build connection string
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, password, dbName, port)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	// verify connection
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)

		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Error migrating database: %v", err)

		return nil, err
	}

	return db, nil
}
