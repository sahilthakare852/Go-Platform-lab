package health

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func Check(service Service) error {
	resp, err := client.Get(service.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

func CheckAll(services []Service) {
	var wg sync.WaitGroup
	results := make(chan HealthCheckResult)
	for _, service := range services {
		wg.Add(1)
		go func(service Service) {
			defer wg.Done()
			//fmt.Printf("Checking %s...\n", service.Name)
			err := Check(service)
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
