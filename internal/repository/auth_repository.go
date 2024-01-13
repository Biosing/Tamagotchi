package repository

import (
	"Tamagotchi/internal/model"
	"context"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) model.AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (a *authRepository) Registration(c context.Context, user *model.User) error {
	result := a.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *authRepository) GetUserByEmailOrLogin(c context.Context, loginOrEmail string) (*model.User, error) {
	var user model.User
	result := a.DB.Where("login = ? OR email = ?", loginOrEmail, loginOrEmail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (a *authRepository) SaveSession(c context.Context, session *model.Session) error {
	result := a.DB.Create(session)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a *authRepository) CheckSession(c context.Context, sessionID string) (*model.Session, error) {
	var session model.Session
	result := a.DB.Where("id = ?", sessionID).First(&session)
	if result.Error != nil {
		return nil, result.Error
	}
	return &session, nil
}
