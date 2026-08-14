// Package wttrclient provides a client to access https://wttr.in.
package wttrclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	timeout = 5 * time.Second

	// baseURL is wttr.in. The ?0A query asks for the current conditions only,
	// in ANSI colour rather than the terminal-sniffed default.
	baseURL = "https://wttr.in"
	query   = "?0A"
)

// URLFor returns the wttr.in URL reporting on location. The location is
// escaped, so a place with a space in its name works without the caller
// thinking about it.
func URLFor(location string) string {
	return baseURL + "/" + url.PathEscape(location) + query
}

// WTTRClient provides a client to access https://wttr.in
type WTTRClient struct {
	client http.Client
	url    string
}

// NewWTTRClient creates a new WTTRClient instance.
func NewWTTRClient(url string) WTTRClient {
	return WTTRClient{
		client: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       timeout,
		},
		url: url,
	}
}

// CurrentWeather returns the current weather for Honolulu.
func (w *WTTRClient) CurrentWeather(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		w.url,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("could not create request: %w", err)
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not do request: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response: %w", err)
	}

	return string(bytes), nil
}
