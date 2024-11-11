package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ZipcodeService struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewZipcodeService(baseUrl string, client *http.Client) *ZipcodeService {
	return &ZipcodeService{BaseURL: baseUrl, HTTPClient: client}
}

func (s *ZipcodeService) Send(cep string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	url := fmt.Sprintf("%s/weather?zipcode=%s", s.BaseURL, cep)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
				return "service-a-send-zipcode"
			}),
		),
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ctx_err := ctx.Err()
	if ctx_err != nil {
		<-ctx.Done()
		err := ctx.Err()
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to send zipcode, status code: %d", resp.StatusCode)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}
