package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// @title Tamagotchi API
// @version 1
// @description Это API для приложения Tamagotchi.
// @termsOfService http://localhost:8080/terms/
// @BasePath /api/v1
// @schemes http
// @host localhost:8080
// @contact name: Vladislav email: Biosing@yandex.ru

// @securityDefinitions.cookie AuthCookie
// @name Auth
// @in cookie
func main() {
	// Строка подключения к базе данных
	connStr := "postgres://postgres:123123123@db:5432/Tamagotchi?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	// Настройка миграции
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Не удалось создать драйвер базы данных: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Путь к файлам миграции
		"postgres", driver)
	if err != nil {
		log.Fatalf("Не удалось создать экземпляр миграции: %v", err)
	}

	// Накатывание миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("Миграция успешно применена")

	router := SetupRouter()
	router.Run(":8080") // запускаем сервер на 8080 порту
}
