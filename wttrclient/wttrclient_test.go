package wttrclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alrayyes/gwttr/wttrclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIClient_CanGetTheCurrentWeather(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		_, _ = w.Write([]byte("Weather report: honolulu"))
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := wttrclient.NewWTTRClient(srv.URL)
	got, err := client.CurrentWeather(t.Context())

	require.NoError(t, err)

	want := "Weather report: honolulu"

	assert.Contains(t, got, want)
}
