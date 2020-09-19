package handler2

import (
	"github.com/incrypt0/cokut-server/brokers"
	"github.com/incrypt0/cokut-server/store"
	"github.com/labstack/echo/v4"
)

//Handler is main api handler object which controlls every network interactions
type Handler struct {
	store         brokers.CokutBroker
	fireAuthMWare echo.MiddlewareFunc
}

//NewHandler is creates a new Handler object
func NewHandler(store *store.Store, fireAuthMWare echo.MiddlewareFunc) *Handler {
	return &Handler{
		store:         store,
		fireAuthMWare: fireAuthMWare,
	}
}