package main

import "github.com/go-playground/log/v8"

func SayHello() string {
	return "test"
}

func main() {
	log.Info("Testing")
	log.Alert("Okayy")
}
