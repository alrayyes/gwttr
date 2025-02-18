package wttrclient_test

import (
	"github.com/alrayyes/gwttr/wttrclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPIClient_CanGetTheWeather(t *testing.T) {
	client := wttrclient.NewWTTRClient()
	got, err := client.GetWeather(t.Context())

	require.NoError(t, err)

	want := "Weather report: honolulu"

	assert.Contains(t, got, want)
}
