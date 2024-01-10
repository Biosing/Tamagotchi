package repository

import (
	"Tamagotchi/internal/model"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetUsers(ctx context.Context) {
}
func (r *userRepository) CreateUser(ctx context.Context) {
}
func (r *userRepository) GetUser(ctx context.Context) {
}
func (r *userRepository) UpdateUser(ctx context.Context) {
}
func (r *userRepository) DeleteUser(ctx context.Context) {
}
