package model

import (
	"time"
)

type Session struct {
	Id        string    `gorm:"primaryKey"` // UUID в качестве первичного ключа
	UserId    uint      // Ссылка на ID пользователя, предполагая что у вас есть структура User
	CreatedAt time.Time // Время создания сессии
	ExpiresAt time.Time // Время истечения сессии
	UserAgent string    // Информация о браузере пользователя
	IPAddress string    // IP-адрес пользователя
}

func (m *Session) TableName() string {
	return "Sessions"
}
