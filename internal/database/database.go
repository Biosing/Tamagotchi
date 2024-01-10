package database

import (
	"Tamagotchi/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataSources struct {
	DB *gorm.DB
}

func Initialize() (*DataSources, error) {
	log.Printf("Initializing data sources\n")

	dsn := "host=db user=postgres password=123123123 dbname=Tamagotchi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Выполнение миграции для всех моделей
	err = db.AutoMigrate(model.User{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	return &DataSources{
		DB: db,
	}, nil
}

func (ds *DataSources) Close() error {
	return nil
}
