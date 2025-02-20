// The gwttr command is a utility that returns the current weather for Honolulu.
package main

import (
	"context"
	"fmt"
	"github.com/alrayyes/gwttr/wttrclient"
	"log"
)

func main() {
	client := wttrclient.NewWTTRClient()
	weather, err := client.CurrentWeather(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(weather)
}
