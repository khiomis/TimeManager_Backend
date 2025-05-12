package dto

type SignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUserDto struct {
	AuthToken          string  `json:"authToken"`
	RefreshToken       string  `json:"refreshToken"`
	NeedSetNewPassword bool    `json:"need_set_new_password"`
	User               UserDTO `json:"user"`
}
