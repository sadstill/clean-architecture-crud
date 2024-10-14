package http

import (
	"errors"
	"net/http"
	"rest-api-crud/internal/model"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var appError *model.Error
		if err := h(writer, request); err != nil {
			writer.Header().Set("Content-Type", "application/json")

			if errors.As(err, &appError) {
				switch {
				case errors.Is(err, model.NotFound):
					writer.WriteHeader(http.StatusNotFound)
					writer.Write(model.NotFound.Marshal())
					return
				default:
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write(model.BadRequest.Marshal())
					return
				}
			}

			writer.WriteHeader(http.StatusTeapot)
			writer.Write(model.Wrap(err).Marshal())
		}
	}
}
