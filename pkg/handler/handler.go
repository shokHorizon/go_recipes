package handler

import (
	"shokHorizon/go_recipes/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	Register(router *httprouter.Router)
	NewHandler(logger *logger.Logger) Handler
}
