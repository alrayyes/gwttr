package wttrclient_test

import (
	"testing"

	"github.com/alrayyes/gwttr/wttrclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIClient_CanGetTheCurrentWeather(t *testing.T) {
	client := wttrclient.NewWTTRClient()
	got, err := client.CurrentWeather(t.Context())

	require.NoError(t, err)

	want := "Weather report: honolulu"

	assert.Contains(t, got, want)
}
