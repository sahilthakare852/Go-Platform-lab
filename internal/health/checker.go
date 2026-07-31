package health

import "net/http"

func Check(service Service) error{
	_, err := http.Get(service.URL)
	if err != nil {
		return err
	}
	return nil
}

