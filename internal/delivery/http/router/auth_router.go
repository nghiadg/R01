package router

import (
	"r01/internal/delivery/http/handler"
	"r01/internal/infrastructure/persistence"
	"r01/internal/usecase/auth"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() *chi.Mux {
	r := chi.NewRouter()

	// create dependencies
	userPostgresRepository := persistence.NewUserPostgresRepository()
	loginUsecase := auth.NewLoginUsecase(userPostgresRepository)
	authHandler := handler.NewAuthHandler(*loginUsecase)

	r.Post("/login", authHandler.Login)

	return r
}
