package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/leetcode-golang-classroom/go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		//middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
