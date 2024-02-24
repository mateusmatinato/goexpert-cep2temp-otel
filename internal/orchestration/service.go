package orchestration

import (
	"context"
	"log"

	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/cep"
	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/orchestration/weather"
)

type Service interface {
	GetTemperatureByCEP(ctx context.Context, request Request) (Response, error)
}

type service struct {
	weatherService weather.Service
	cepService     cep.Service
}

func (s service) GetTemperatureByCEP(ctx context.Context, request Request) (Response, error) {
	log.Printf("starting cep2temp for cep %s\n", request.CEP)

	cepRequest := request.BuildCEPRequest()
	cepResponse, err := s.cepService.GetInfo(ctx, cepRequest)
	if err != nil {
		log.Printf("error getting cep info: %v\n", err)
		return Response{}, err
	}

	weatherRequest := NewWeatherRequest(cepResponse)
	weatherResponse, err := s.weatherService.GetInfo(ctx, weatherRequest)
	if err != nil {
		log.Printf("error getting weather info: %v\n", err)
		return Response{}, err
	}

	resp := NewResponse(cepResponse, weatherResponse)
	log.Printf("finished cep2temp for cep %s | temp_c: %.2f, | temp_f: %.2f, | temp_k: %.2f | city: %s\n",
		request.CEP, resp.TempCelsius, resp.TempFahrenheit, resp.TempKelvin, resp.City)
	return resp, nil
}

func NewService(cepService cep.Service, weatherService weather.Service) Service {
	return &service{
		weatherService: weatherService,
		cepService:     cepService,
	}
}
