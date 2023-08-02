package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/irbgeo/1inch-test/controller"
)

type api struct {
	controller iController
}

type iController interface {
	GetAmountOut(ctx context.Context, in controller.In) (*controller.Out, error)
}

func New(
	controller iController,
) *api {
	return &api{controller: controller}
}

func (s *api) Route(r *mux.Router) {
	r.HandleFunc("/get-amount-out", s.getAmountOut).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
