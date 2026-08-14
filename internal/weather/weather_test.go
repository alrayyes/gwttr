package weather_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alrayyes/gwttr/internal/weather"
	"github.com/alrayyes/gwttr/internal/wttrclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errUnavailable = errors.New("no forecast today")

// fixedSource is an in-memory Source rather than a mock: it really does return
// a forecast, it just always returns the same one. Nothing here records how it
// was called.
type fixedSource struct {
	forecast string
	err      error
}

func (s fixedSource) CurrentWeather(context.Context) (string, error) {
	return s.forecast, s.err
}

func TestReport(t *testing.T) {
	t.Parallel()

	t.Run("writes the forecast", func(t *testing.T) {
		t.Parallel()

		src := fixedSource{forecast: "Weather report: honolulu", err: nil}

		var out bytes.Buffer

		require.NoError(t, weather.Report(t.Context(), src, &out))

		assert.Equal(t, "Weather report: honolulu\n", out.String())
	})

	t.Run("passes the source's failure back to the caller", func(t *testing.T) {
		t.Parallel()

		src := fixedSource{forecast: "", err: errUnavailable}

		err := weather.Report(t.Context(), src, &bytes.Buffer{})

		require.ErrorIs(t, err, errUnavailable)
	})

	t.Run("writes nothing when the source fails", func(t *testing.T) {
		t.Parallel()

		src := fixedSource{forecast: "", err: errUnavailable}

		var out bytes.Buffer

		_ = weather.Report(t.Context(), src, &out)

		assert.Empty(t, out.String())
	})
}

// The same specification as TestReport's happy path, driven through the real
// wttrclient over a real socket. This is what keeps fixedSource honest — a fake
// that has drifted from the adapter it stands in for is worse than no fake.
func TestReportAgainstTheRealClient(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		_, _ = w.Write([]byte("Weather report: honolulu"))
	}))
	defer srv.Close()

	client := wttrclient.NewWTTRClient(srv.URL)

	var out bytes.Buffer

	require.NoError(t, weather.Report(t.Context(), &client, &out))

	assert.Equal(t, "Weather report: honolulu\n", out.String())
}
