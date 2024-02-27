package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/orchestrationapi/config"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/orchestrationapi/router"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/cep"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/weather"
	platHttp "github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/http"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/metrics"
	"go.opentelemetry.io/otel"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		panic(fmt.Sprintf("error starting configs: %s", err.Error()))
	}

	shutdown, err := metrics.InitProvider(cfg.ServiceName, cfg.OTELExporterEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	tracer := otel.Tracer(cfg.ServiceName)

	httpClient := platHttp.NewDefaultClient()
	cepService := cep.NewService(tracer, httpClient, cfg.CepAPIConfig())
	weatherService := weather.NewService(tracer, httpClient, cfg.WeatherAPIConfig())

	service := orchestration.NewService(cepService, weatherService)
	handler := orchestration.NewHandler(service)

	r := router.SetupRouter(cfg.ServiceName, handler)
	fmt.Println("started http server - orchestration")
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		panic(fmt.Sprintf("error on listen and serve: %s", err.Error()))
	}
}
