package dto

type CreateUserResponse struct {
	ID string `json:"id"`
}

type CreateUserRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
	Email        string `json:"email"`
}
