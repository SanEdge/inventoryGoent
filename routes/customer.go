package routes

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/inventorygoent/ent"
	"github.com/renaldyhidayatt/inventorygoent/handler"
	"github.com/renaldyhidayatt/inventorygoent/middlewares"
	"github.com/renaldyhidayatt/inventorygoent/repository"
	"github.com/renaldyhidayatt/inventorygoent/services"
)

func NewCustomerRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewCustomerRepository(db, context)
	service := services.NewCustomerService(repository)
	handler := handler.NewCustomerHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", handler.Results)
		r.Get("/{id:[0-9]+}", handler.Result)
		r.Post("/", handler.Create)
		r.Put("/{id:[0-9]+}", handler.UpdateById)
		r.Delete("/{id:[0-9]+}", handler.DeleteById)
	})
}
