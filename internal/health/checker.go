package health

import (
	"fmt"
	"net/http"
)

func Check(service Service) error {
	resp, err := http.Get(service.URL)
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
	for _, service := range services {
		fmt.Printf("Checking %s...\n", service.Name)
		err := Check(service)
		if err == nil {
			fmt.Printf("PASSED\n")
		} else {
			fmt.Printf("FAILED:  %v\n", err)
		}
	}
}
