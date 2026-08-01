package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sahilthakare852/go-platform-lab/internal/health"
)

func main() {
	config, err := health.LoadConfig("./configs/services.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	health.CheckAll(ctx, config.Services, client)
}
