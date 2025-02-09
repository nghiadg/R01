package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/auth", AuthRouter())
	r.Mount("/manage-car", ManageCarRouter(db))

	return r
}
