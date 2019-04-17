package internal

import (
	"encoding/json"
	"log"
)

type Service struct {
	// ID string
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	History  `json:"history,omitempty"`
}

func NewService(service map[string]string) *Service {
	return &Service{
		Name:     service["name"],
		Endpoint: service["endpoint"],
		History:  make(History),
	}
}

func (s Service) String() string {
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
