package routes

import (
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	serverMux *chi.Mux
}

func NewRoutes(serverMux *chi.Mux,
) *Routes {
	return &Routes{
		serverMux: serverMux,
	}
}
func (r *Routes) RegisterRoutes() {}
