package model

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username" gorm:"unique;not null"`
	Email      string    `json:"email" gorm:"unique;not null"`
	Password   string    `json:"password" gorm:"not null;"`
	Role       string    `json:"role" gorm:"not null;"`
	IsVerified bool      `json:"verified" gorm:"not null;default:false"`
	CreatedAt  time.Time `json:"created_At" gorm:"not null;date"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
