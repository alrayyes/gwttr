// Package weather reports a forecast to somewhere a person can read it.
package weather

import (
	"context"
	"fmt"
	"io"
)

// Source is a place a forecast can be fetched from. It is declared here, in the
// package that consumes it, so nothing about wttr.in or HTTP reaches this far
// in — swapping the client for another service means writing another Source,
// not touching Report.
type Source interface {
	CurrentWeather(ctx context.Context) (string, error)
}

// Report fetches the current forecast from src and writes it to out.
func Report(ctx context.Context, src Source, out io.Writer) error {
	forecast, err := src.CurrentWeather(ctx)
	if err != nil {
		return fmt.Errorf("could not get the current weather: %w", err)
	}

	_, err = fmt.Fprintln(out, forecast)
	if err != nil {
		return fmt.Errorf("could not write the current weather: %w", err)
	}

	return nil
}
