package model

import "time"

type ChatHeader struct {
	ChatID string `gorm:"primaryKey"`
}

type ChatDetails struct {
	ChatID    string      `gorm:"primaryKey;foreignKey:ChatID"`
	SenderID  string      `gorm:"primaryKey;foreignKey:UserID"`
	Message   string      `json:"message"`
	Chat      *ChatHeader `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Sender    *User       `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE"`
	CreatedAt time.Time   `json:"created_at" gorm:"primaryKey;date"`
}
