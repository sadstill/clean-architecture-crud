package model

type CreateUserResponse struct {
	ID string `json:"id"`
}

type CreateUserRequest struct {
	Username     string `db:"username"`
	PasswordHash string `db:"password"`
	Email        string `db:"email"`
}
