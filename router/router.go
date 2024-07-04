package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {

	r := chi.NewMux()
	r.Use(middleware.Logger)

	return r
}
