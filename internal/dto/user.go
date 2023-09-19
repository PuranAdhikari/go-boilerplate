package dto

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email" label:"Email"`
	Password        string `json:"password" binding:"required,min=8" label:"Password"`
	ConfirmPassword string `json:"confirm_password" binding:"eqfield=Password" label:"Confirm Password"`
}

type RegisterResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
