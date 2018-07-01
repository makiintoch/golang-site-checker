package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Config struct {
	CodeTime  int   `json:"getCodeTime"`
	EmailTime int   `json:"sendEmailTime"`
	Urls      []Url `json:"urls"`
}

type Url struct {
	UrlName string `json:"url"`
}

type Email struct {
	From string
	To   string
}

func getStatusCode(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	return resp.StatusCode
}

func startPolling(cfg Config) {
	for {
		time.Sleep(time.Duration(cfg.CodeTime) * time.Minute)

		urls := cfg.Urls

		for _, url := range urls {
			code := getStatusCode(url.UrlName)

			if code < 200 || code > 299 {
				fmt.Println("Argh! Broken: ", code)
			} else {
				fmt.Println("All ok: ", code)
			}
		}
	}
}

func main() {
	configFile, err := ioutil.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err)
	}

	var cfg Config

	json.Unmarshal(configFile, &cfg)

	startPolling(cfg)
}
