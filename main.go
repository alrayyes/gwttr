// The gwttr command is a utility that reports the current weather.
package main

import (
	"context"
	"io"
	"os"

	"github.com/alrayyes/gwttr/weather"
	"github.com/alrayyes/gwttr/wttrclient"
	"github.com/spf13/cobra"
)

const defaultLocation = "honolulu"

// version is set at build time by goreleaser's default ldflags, which pass
// -X main.version=<tag>. A build that isn't a release says so rather than
// claiming a version it hasn't got.
var version = "dev"

func main() {
	cmd := &cobra.Command{
		Use:   "gwttr [location]",
		Short: "Report the current weather",
		Long: "Report the current weather for a location, as coloured ASCII " +
			"art from wttr.in.\n\nWith no location, it reports on " +
			defaultLocation + ".",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		Version:      version,
		RunE: func(cmd *cobra.Command, args []string) error {
			location := defaultLocation
			if len(args) == 1 {
				location = args[0]
			}

			return report(cmd.Context(), location, cmd.OutOrStdout())
		},
	}

	// cobra prints the error and this sets the exit code, so nothing below
	// needs to call os.Exit itself.
	err := cmd.ExecuteContext(context.Background())
	if err != nil {
		os.Exit(1)
	}
}

func report(ctx context.Context, location string, out io.Writer) error {
	client := wttrclient.NewWTTRClient(wttrclient.URLFor(location))

	return weather.Report(ctx, &client, out)
}
