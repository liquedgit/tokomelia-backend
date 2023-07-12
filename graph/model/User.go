package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null;"`
	Role     string `json:"role" gorm:"not null;"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
