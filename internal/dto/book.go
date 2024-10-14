package dto

type CreateBookResponse struct {
	ID string `json:"id"`
}

type CreateBookRequest struct {
	Username     string `db:"username"`
	PasswordHash string `db:"password"`
	Email        string `db:"email"`
}
