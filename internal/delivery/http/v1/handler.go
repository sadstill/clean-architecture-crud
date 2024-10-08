package v1

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	httpHandler "rest-api-crud/internal/delivery/http"
	"rest-api-crud/pkg/logging"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) httpHandler.Handler {
	return &handler{logger}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, httpHandler.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, httpHandler.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, httpHandler.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, httpHandler.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, httpHandler.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, httpHandler.Middleware(h.DeleteUser))
}
