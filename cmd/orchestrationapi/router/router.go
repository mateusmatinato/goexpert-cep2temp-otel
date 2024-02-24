package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration"
)

func SetupRouter(handler orchestration.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/orchestrate/{cep}", handler.GetTemperatureByCEP).Methods(http.MethodGet)
	return r
}
