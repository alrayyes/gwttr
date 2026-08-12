// The gwttr command is a utility that returns the current weather for Honolulu.
package main

import (
	"context"
	"log"
	"os"

	"github.com/alrayyes/gwttr/weather"
	"github.com/alrayyes/gwttr/wttrclient"
)

const url = "https://wttr.in/honolulu?0A"

func main() {
	client := wttrclient.NewWTTRClient(url)

	err := weather.Report(context.Background(), &client, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
