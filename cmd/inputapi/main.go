package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/inputapi/config"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/inputapi/router"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input/orchestration"
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
	orchClient := orchestration.NewClient(tracer, httpClient, cfg.OrchClientConfig())

	service := input.NewService(orchClient)
	handler := input.NewHandler(service)

	r := router.SetupRouter(cfg.ServiceName, handler)
	fmt.Println("started http server - input")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(fmt.Sprintf("error on listen and serve: %s", err.Error()))
	}
}
