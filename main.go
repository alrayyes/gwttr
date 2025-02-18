package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	const timeout = 10 * time.Second

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       timeout,
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://wttr.in/honolulu?0A",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(body))
}
