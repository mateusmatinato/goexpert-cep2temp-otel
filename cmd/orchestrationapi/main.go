package main

import (
	"fmt"
	"net/http"

	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/orchestrationapi/config"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/orchestrationapi/router"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/cep"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/weather"
	platHttp "github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/http"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		panic(fmt.Sprintf("error starting configs: %s", err.Error()))
	}

	httpClient := platHttp.NewDefaultClient()
	cepService := cep.NewService(httpClient, cfg.CepAPIConfig())
	weatherService := weather.NewService(httpClient, cfg.WeatherAPIConfig())
	service := orchestration.NewService(cepService, weatherService)

	handler := orchestration.NewHandler(service)

	r := router.SetupRouter(handler)
	fmt.Println("started http server - orchestration")
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		panic(fmt.Sprintf("error on listen and serve: %s", err.Error()))
	}
}
