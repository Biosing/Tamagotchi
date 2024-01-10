package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Login     string    `gorm:"size:50;not null"`
	Password  string    `gorm:"size:1000;not null"`
	FullName  string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
}

func (m *User) TableName() string {
	return "Users"
}
