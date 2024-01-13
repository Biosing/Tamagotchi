package database

import (
	"Tamagotchi/internal/model"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataSources struct {
	DB *gorm.DB
}

func Initialize() (*DataSources, error) {
	log.Printf("Initializing data sources\n")

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbSslmode := viper.GetString("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Выполнение миграции для всех моделей
	err = db.AutoMigrate(&model.User{}, &model.Session{}, &model.Сreature{}, &model.UserCreature{})
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
