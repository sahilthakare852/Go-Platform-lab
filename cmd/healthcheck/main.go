package main

import (
	"fmt"

	"github.com/sahilthakare852/go-platform-lab/internal/health"
)

func main() {
	config, err := health.LoadConfig("services.json")
	if err != nil {
		fmt.Println(err)
	} else {
		health.CheckAll(config.Services)
	}
}
