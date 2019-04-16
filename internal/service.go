package internal

const (
	Status     = "status.json"
	Components = "components.json"
)

type Service struct {
	// ID string
	Name     string
	Endpoint string
}

func NewService(service map[string]string) Service {
	return Service{
		Name:     service["name"],
		Endpoint: service["endpoint"],
	}
}
