package health

type Service struct {
	Name string
	URL  string
}

type ServiceConfig struct {
	Services []Service `json:"services"`
}
