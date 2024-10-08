package v1

import (
	"fmt"
	"net/http"
	"rest-api-crud/internal/apperror"
)

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.NotFound
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	w.Write([]byte("this is get user by uuid"))
	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("this is api error")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))
	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update user"))
	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))
	return nil
}
