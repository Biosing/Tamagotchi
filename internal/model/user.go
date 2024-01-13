package model

import (
	"time"
)

// User структура, представляющая пользователя.
// swagger:model User
type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Login     string    `json:"login" gorm:"size:150;not null"`
	Email     string    `json:"email" gorm:"size:150;not null"`
	Password  string    `json:"password" gorm:"size:1000;not null"`
	FullName  string    `json:"fullName" gorm:"size:255"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *User) TableName() string {
	return "Users"
}
