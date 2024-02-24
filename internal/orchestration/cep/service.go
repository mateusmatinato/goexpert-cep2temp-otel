package cep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	internalErrors "github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/errors"
)

const (
	NotFoundCEP      = "ERR_NOT_FOUND_CEP"
	FailedGetInfo    = "ERR_FAILED_GET_CEP"
	FailedUnmarshall = "ERR_FAILED_UNMARSHALL_CEP"
)

type Service interface {
	GetInfo(ctx context.Context, request Request) (Response, error)
}

type service struct {
	client    http.Client
	apiConfig APIConfig
}

func (s service) GetInfo(_ context.Context, request Request) (Response, error) {
	res, err := s.client.Get(fmt.Sprintf(s.apiConfig.URL, request.Cep))
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo, err)
	}

	if res.StatusCode == http.StatusNotFound {
		return Response{}, internalErrors.NewNotFoundError(NotFoundCEP, err)
	}
	if res.StatusCode != http.StatusOK {
		return Response{}, internalErrors.NewApplicationError(FailedGetInfo,
			errors.New(fmt.Sprintf("status_code:%d", res.StatusCode)))
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedUnmarshall, err)
	}

	if resp.City == "" {
		return Response{}, internalErrors.NewNotFoundError(NotFoundCEP, err)
	}

	return resp, nil
}

func NewService(client http.Client, config APIConfig) Service {
	return &service{
		client:    client,
		apiConfig: config,
	}
}
