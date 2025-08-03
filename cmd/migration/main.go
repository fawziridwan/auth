package main

import (
	"log"
	"os"

	"github.com/fawziridwan/auth_module/internal/config"
	"github.com/fawziridwan/auth_module/internal/models"
	"github.com/fawziridwan/auth_module/internal/utils/database"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Load configuration
	cfg := config.Config{
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBSource:   os.Getenv("DB_SOURCE"),
		ServerPort: os.Getenv("SERVER_PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}

	// 3. Validate required config
	if cfg.DBDriver == "" {
		log.Fatal("DB_DRIVER is not set in environment variables")
	}
	if cfg.DBSource == "" {
		log.Fatal("DB_SOURCE is not set in environment variables")
	}

	// 4. Initialize database
	var db database.Database
	switch cfg.DBDriver {
	case "mysql":
		db, err = database.NewMySQLDatabase(database.DBConfig{
			Driver: cfg.DBDriver,
			DSN:    cfg.DBSource,
		})
	case "postgres":
		db, err = database.NewPostgresDatabase(database.DBConfig{
			Driver: cfg.DBDriver,
			DSN:    cfg.DBSource,
		})
	default:
		log.Fatalf("Unsupported database driver: %s", cfg.DBDriver)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 5. Run migrations
	if err := db.Migrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migration completed successfully")
}
