package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("do something every 2 minutes")
}

func startPolling() {
	for {
		time.Sleep(2 * time.Minute)

		doSomething()
	}
}

func main() {
	startPolling()
}
