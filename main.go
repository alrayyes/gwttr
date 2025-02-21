// The gwttr command is a utility that returns the current weather for Honolulu.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alrayyes/gwttr/wttrclient"
)

func main() {
	client := wttrclient.NewWTTRClient()
	weather, err := client.CurrentWeather(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(weather)
}
