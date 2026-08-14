package wttrclient_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/alrayyes/gwttr/internal/wttrclient"
)

// A local test server stands in for wttr.in here. Pointing the example at the
// real service would make its expected output depend on today's weather, and
// on the network being up when the suite runs.
func ExampleWTTRClient_CurrentWeather() {
	srv := httptest.NewServer(http.HandlerFunc(func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		_, _ = w.Write([]byte("Weather report: honolulu"))
	}))
	defer srv.Close()

	client := wttrclient.NewWTTRClient(srv.URL)

	forecast, err := client.CurrentWeather(context.Background())
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(forecast)
	// Output: Weather report: honolulu
}
