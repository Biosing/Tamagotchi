package model

import (
	"time"
)

// Сreature структура, представляющая существ.
// swagger:model Сreature
type Сreature struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"fullName" gorm:"size:255"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *Сreature) TableName() string {
	return "Сreatures"
}
