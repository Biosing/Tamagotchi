package main

import (
	"database/sql"
	"log"

	"Tamagotchi/internal/database"

	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// @title Tamagotchi API
// @version 1
// @description This is the API for the Tamagotchi application.
// @termsOfService http://localhost:8080/terms/
// @BasePath /api/v1
// @schemes http
// @host localhost:8080
// @contact name: Vladislav email: Biosing@yandex.ru

// @securityDefinitions.cookie AuthCookie
// @name Auth
// @in cookie
func main() {
	initViper()

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbSslmode := viper.GetString("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	ds, err := database.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}

	router, err := inject(ds)
	if err != nil {
		log.Fatalf("Injection error: %v", err)
	}

	// Applying migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration error: %v", err)
	}

	log.Println("Migration successfully applied")

	runPort := viper.GetString("APP_PORT")
	router.Run(runPort)
}

func initViper() {
	viper.SetConfigFile("config/.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}
}
