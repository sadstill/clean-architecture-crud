package http

import (
	"errors"
	"net/http"
	errors2 "rest-api-crud/internal/apperror"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var appError *errors2.Error
		if err := h(writer, request); err != nil {
			writer.Header().Set("Content-Type", "application/json")

			if errors.As(err, &appError) {
				switch {
				case errors.Is(err, errors2.NotFound):
					writer.WriteHeader(http.StatusNotFound)
					writer.Write(errors2.NotFound.Marshal())
					return
				default:
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write(errors2.BadRequest.Marshal())
					return
				}
			}

			writer.WriteHeader(http.StatusTeapot)
			writer.Write(errors2.Wrap(err).Marshal())
		}
	}
}
