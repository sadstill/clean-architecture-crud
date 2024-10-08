package http

import (
	"errors"
	"net/http"
	"rest-api-crud/internal/apperror"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var appError *apperror.Error
		err := h(writer, request)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appError) {
				if errors.Is(err, apperror.NotFound) {
					writer.WriteHeader(http.StatusNotFound)
					writer.Write(apperror.NotFound.Marshal())
					return
				}
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write(apperror.BadRequest.Marshal())
				return
			}

			writer.WriteHeader(http.StatusTeapot)
			writer.Write(apperror.Wrap(err).Marshal())
		}
	}
}
