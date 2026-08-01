package health

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var client HTTPClient = &http.Client{
	Timeout: 5 * time.Second,
}

func Check(ctx context.Context, service Service, client HTTPClient) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, service.URL, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func CheckAll(ctx context.Context, services []Service, client HTTPClient) {
	var wg sync.WaitGroup
	results := make(chan HealthCheckResult)
	for _, service := range services {
		wg.Add(1)
		go func(service Service) {
			defer wg.Done()
			//fmt.Printf("Checking %s...\n", service.Name)
			err := Check(ctx, service, client)
			//if err == nil {
			//	fmt.Printf("PASSED\n")
			//} else {
			//	fmt.Printf("FAILED:  %v\n", err)
			//}
			results <- HealthCheckResult{Service: service, Error: err}
		}(service)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		if result.Error != nil {
			fmt.Printf("%s Failed: %v\n", result.Service.Name, result.Error)
		} else {
			fmt.Printf("%s Passed\n", result.Service.Name)
		}
	}
}
