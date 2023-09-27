package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Nik        string `json:"nik"`
	Token      string `json:"token"`
	Role       string `json:"role"`
	AvatarPath string `json:"avatar_path"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Nik:        user.Nik,
		Token:      token,
		Role:       user.Role,
		AvatarPath: user.AvatarFileName,
	}

	return formatter
}

func FormatUsers(users []User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		userFormatter := FormatUser(user, "")
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
