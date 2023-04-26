package request

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type AccountConfirmPayload struct {
	Token string `json:"token"`
}
