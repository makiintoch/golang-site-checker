package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getStatusCode() int {
	resp, err := http.Get("https://google.com/")

	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode
}

func startPolling() {
	for {
		time.Sleep(1 * time.Minute)

		code := getStatusCode()

		if code < 200 || code > 299 {
			fmt.Println("Argh! Broken: ", code)
		} else {
			fmt.Println("All ok: ", code)
		}
	}
}

func main() {
	defer startPolling()
}
