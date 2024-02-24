package input

import (
	"context"
	"errors"
	"log"

	"github.com/mateusmatinato/goexpert-cep2temp-otel/internal/input/orchestration"
)

var (
	ErrOrchestrationCEP = errors.New("ERR_ORCHESTRATION_CEP")
)

type Service interface {
	GetTemperatureByCEP(ctx context.Context, request Request) (Response, error)
}

type service struct {
	orchClient orchestration.Client
}

func (s service) GetTemperatureByCEP(ctx context.Context, request Request) (Response, error) {
	if err := request.Validate(); err != nil {
		log.Printf("error validating request: %s\n", err.Error())
		return Response{}, err
	}

	resp, err := s.orchClient.GetTemperatureByCEP(ctx, request.CEP)
	if err != nil {
		log.Printf("error getting orchestration cep2temp: %s\n", err.Error())
		return Response{}, ErrOrchestrationCEP
	}

	return NewResponse(resp), nil
}

func NewService(orchClient orchestration.Client) Service {
	return &service{
		orchClient: orchClient,
	}
}
