package recipe

import (
	"net/http"
	"shokHorizon/go_recipes/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

var recipeURL = "/recipe"

type recipeHandler struct {
	l *logger.Logger
}

func NewHandler(l *logger.Logger) recipeHandler {
	return recipeHandler{l}
}

func (h *recipeHandler) Register(r *httprouter.Router) {
	r.GET(recipeURL, h.GetRecipe)
}

func (h *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("This is the user"))
}
