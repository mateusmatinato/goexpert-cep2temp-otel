package main

import (
	"fmt"
	"net/http"

	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/inputapi/config"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/cmd/inputapi/router"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input/orchestration"
	platHttp "github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/http"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		panic(fmt.Sprintf("error starting configs: %s", err.Error()))
	}

	httpClient := platHttp.NewDefaultClient()
	orchClient := orchestration.NewClient(httpClient, cfg.OrchClientConfig())

	service := input.NewService(orchClient)
	handler := input.NewHandler(service)

	r := router.SetupRouter(handler)
	fmt.Println("started http server - input")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(fmt.Sprintf("error on listen and serve: %s", err.Error()))
	}
}
