package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/goschool/crud/api"
	"github.com/goschool/crud/middleware"
)

func SetupRoutes(userHandler api.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/echo", api.HandleEchoUser)
	r.Post("/register", userHandler.HandlerRegisterUser)
	r.Post("/login", userHandler.HandlerLoginUser)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.UserAuthentication)
		r.Get("/test", api.HandleTest)
	})
	return r
}
