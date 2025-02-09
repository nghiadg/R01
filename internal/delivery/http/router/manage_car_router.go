package router

import (
	"r01/internal/delivery/http/handler"
	"r01/internal/infrastructure/persistence"
	"r01/internal/usecase/manage_car"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ManageCarRouter(db *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	// create dependencies
	carPostgresRepository := persistence.NewCarPostgresRepository(db)
	createCarUsecase := manage_car.NewCreateCarUsecase(carPostgresRepository)
	listCarUsecase := manage_car.NewListCarUsecase(carPostgresRepository)
	manageCarHandler := handler.NewManageCarHandler(*createCarUsecase, *listCarUsecase)

	r.Post("/create", manageCarHandler.CreateCar)
	r.Get("/list", manageCarHandler.ListCar)

	return r
}
