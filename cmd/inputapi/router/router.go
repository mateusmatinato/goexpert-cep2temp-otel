package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input"
)

func SetupRouter(handler input.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/input", handler.GetTemperatureByCEP).Methods(http.MethodPost)
	return r
}
