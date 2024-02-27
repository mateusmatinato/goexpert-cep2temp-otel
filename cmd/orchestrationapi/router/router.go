package router

import (
	"github.com/gorilla/mux"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/http"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

func SetupRouter(serviceName string, handler orchestration.Handler) *mux.Router {
	r := mux.NewRouter()

	r.Use(otelmux.Middleware(serviceName))
	r.Use(http.RequestIDMiddleware)

	r.HandleFunc("/orchestrate/{cep}", handler.GetTemperatureByCEP)
	return r
}
