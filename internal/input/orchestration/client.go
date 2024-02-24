package orchestration

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	internalErrors "github.com/mateusmatinato/goexpert-cep2temp-otel/internal/platform/errors"
)

const (
	FailedOrchestration = "ERR_EXECUTING_ORCHESTRATION"
	FailedUnmarshall    = "ERR_FAILED_UNMARSHALL_ORCHESTRATION"
)

type Client interface {
	GetTemperatureByCEP(ctx context.Context, cep string) (Response, error)
}

type clientHandler struct {
	client    http.Client
	apiConfig APIConfig
}

func (c clientHandler) GetTemperatureByCEP(ctx context.Context, cep string) (Response, error) {
	url := fmt.Sprintf("%s/%s", c.apiConfig.OrchestrationURL, cep)
	log.Printf("calling url: %s\n", url)
	res, err := c.client.Get(url)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedOrchestration, err)
	}
	if res.StatusCode != http.StatusOK {
		return Response{}, internalErrors.NewApplicationError(FailedOrchestration,
			errors.New(fmt.Sprintf("status_code:%d", res.StatusCode)))
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, internalErrors.NewApplicationError(FailedUnmarshall, err)
	}

	return resp, nil
}

func NewClient(client http.Client, apiConfig APIConfig) Client {
	return &clientHandler{client: client, apiConfig: apiConfig}
}
