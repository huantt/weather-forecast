package weatherapi_com

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log/slog"
	"net/http"
	"time"
)

type Service struct {
	httpClient *resty.Client
	key        string
}

type Option func(s *Service)

func WithEndpoint(endpoint string) Option {
	return func(s *Service) {
		s.httpClient.SetBaseURL(endpoint)
	}
}

func NewService(key string, opts ...Option) *Service {
	httpClient := resty.New()
	httpClient.
		SetRetryCount(12).
		SetRetryWaitTime(5 * time.Second).
		SetBaseURL(APIEndpoint).AddRetryCondition(func(response *resty.Response, err error) bool {
		if err != nil {
			return true
		}
		if response.StatusCode() == http.StatusInternalServerError ||
			response.StatusCode() == http.StatusBadGateway ||
			response.StatusCode() == http.StatusGatewayTimeout ||
			response.StatusCode() == http.StatusServiceUnavailable {
			slog.Warn(fmt.Sprintf("Response status code is %d - Request: %s - Body: %s - Retrying...", response.StatusCode(), response.Request.URL, response.Body()))
			return true
		}

		return false
	})
	service := &Service{httpClient, key}
	for _, opt := range opts {
		opt(service)
	}
	return service
}

func (s *Service) Forecast(ctx context.Context, city string, days int) (*Forecast, error) {
	var forecast *Forecast
	resp, err := s.httpClient.R().SetContext(ctx).
		SetQueryParam("q", city).
		SetQueryParam("days", fmt.Sprintf("%d", days)).
		SetQueryParam("key", s.key).
		SetResult(&forecast).
		Get("/v1/forecast.json")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(fmt.Sprintf("Request: %s - Response code: %d - Response body: %s", resp.Request.URL, resp.StatusCode(), resp.Body()))
	}

	return forecast, nil
}
