package main

import (
	"context"
	"fmt"
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
	health.CheckAll(ctx, config.Services)
}
