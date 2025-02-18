package main

import (
	"context"
	"fmt"
	"github.com/alrayyes/gwttr/apiclient"
	"log"
)

func main() {
	client := apiclient.NewAPIClient()
	weather, err := client.GetWeather(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(weather)
}
