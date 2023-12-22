package user

type RegisterUserInput struct {
	Nik      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" `
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	Type     string `json:"type" binding:"required"`
	Token    string `json:"token"`
}

type EmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type GetUserDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateUserInput struct {
	Nik        string `json:"nik" `
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password"`
	Role       string `json:"role" binding:"required"`
	AvatarPath string `json:"avatar_path"`
	Token      string `json:"token"`
}
