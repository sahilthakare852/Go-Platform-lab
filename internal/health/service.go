package health

func NewService(name string, url string) Service {
	return Service{
		Name: name,
		URL:  url,
	}
}
