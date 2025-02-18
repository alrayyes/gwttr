package apiclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const timeout = 10 * time.Second
const url = "https://wttr.in/honolulu?0A"

type APIClient struct {
	client http.Client
}

func (a *APIClient) GetWeather(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not do request: %w", err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response: %w", err)
	}

	return string(bytes), nil
}

func NewAPIClient() APIClient {
	client := APIClient{
		client: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       timeout,
		},
	}

	return client
}
