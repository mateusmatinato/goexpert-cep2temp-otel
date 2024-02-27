package router

import (
	"github.com/gorilla/mux"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/http"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func SetupRouter(serviceName string, handler input.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(otelmux.Middleware(serviceName))
	r.Use(http.RequestIDMiddleware)

	r.HandleFunc("/input", handler.GetTemperatureByCEP)
	return r
}
