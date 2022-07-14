package validations

type RegisterAuthValidation struct {
	Name     string `json:"name" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type LoginAuthValidation struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

type RefreshAuthValidation struct {
	UserID       string `json:"user_id" validate:"required,uuid"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
