package main

import (
	"github.com/sahilthakare852/go-platform-lab/internal/health"
)

var services = []health.Service{
	health.NewService("Google", "https://google.com"),
	health.NewService("GitHub", "https://github.com"),
	health.NewService("OpenAI", "https://OpenAI.com"),
}

func main() {
	health.CheckAll(services)
}
