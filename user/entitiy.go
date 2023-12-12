package user

import "time"

type User struct {
	ID         int
	Nik        string
	Name       string
	Email      string
	Password   string
	AvatarPath string
	Role       string
	Token      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (User) TableName() string {
	return "tb_user"
}
