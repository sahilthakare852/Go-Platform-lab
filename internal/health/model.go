package health

type Service struct {
	Name string
	URL  string
}

type ServiceConfig struct {
	Services []Service `json:"services"`
}

type HealthCheckResult struct {
	Service Service
	Error   error
}
