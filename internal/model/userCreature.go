package model

import (
	"time"
)

// UserCreature структура, представляющая существ.
// swagger:model UserCreature
type UserCreature struct {
	UserId     uint      `json:"user_id"`
	CreatureID uint      `json:"creature_id"`
	CreatedAt  time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (m *UserCreature) TableName() string {
	return "UserCreaturs"
}
