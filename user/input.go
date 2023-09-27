package user

type RegisterUserInput struct {
	Nik      string `json:"nik" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type EmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
