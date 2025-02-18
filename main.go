package main

import (
	"context"
	"fmt"
	"github.com/alrayyes/gwttr/wttrclient"
	"log"
)

func main() {
	client := wttrclient.NewWTTRClient()
	weather, err := client.GetWeather(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(weather)
}
