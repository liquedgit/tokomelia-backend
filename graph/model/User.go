package model

type User struct {
	ID       string `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
