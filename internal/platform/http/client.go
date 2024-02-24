package http

import (
	"net/http"
	"time"
)

func NewDefaultClient() http.Client {
	return http.Client{
		Timeout: 3 * time.Second,
	}
}
