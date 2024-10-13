package model

type Book struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Author Author `json:"author"`
}

type CreateBookResponse struct {
	ID string `json:"id"`
}

type CreateBookRequest struct {
	Username     string `db:"username"`
	PasswordHash string `db:"password"`
	Email        string `db:"email"`
}
