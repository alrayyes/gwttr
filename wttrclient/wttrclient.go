// Package wttrclient provides a client to access https://wttr.in.
package wttrclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const timeout = 5 * time.Second
const url = "https://wttr.in/honolulu?0A"

// WTTRClient provides a client to access https://wttr.in
type WTTRClient struct {
	client http.Client
}

// NewWTTRClient creates a new WTTRClient instance.
func NewWTTRClient() WTTRClient {
	client := WTTRClient{
		client: http.Client{
			Timeout: timeout,
		},
	}

	return client
}

// CurrentWeather returns the current weather for Honolulu.
func (w *WTTRClient) CurrentWeather(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}

	resp, err := w.client.Do(req)
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
