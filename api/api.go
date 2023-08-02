package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"

	"1inch-test/models"
)

type api struct {
	core core
}

type core interface {
	GetAmountOut(ctx context.Context, in models.In) (*models.Out, error)
}

func New(
	core core,
) *api {
	return &api{core: core}
}

func (s *api) Route(r *mux.Router) {
	r.HandleFunc("/get-amount-out", s.getAmountOut).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
