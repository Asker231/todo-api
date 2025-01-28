package auth

type (
	RegisterRequest struct{
		Name string `json:"name"`
		Email string `json:"email" validate:"email"`
		Password string `json:"password"`
	}
	RegisterResponse struct{
		Token string `json:"token"`
	}
)