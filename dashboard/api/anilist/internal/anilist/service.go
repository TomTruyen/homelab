package anilist

import (
	"net/http"
	"time"
)

type Service struct {
	httpClient   *http.Client
	clientID     string
	clientSecret string
	username     string
}

func NewService(clientID, clientSecret, username string) *Service {
	return &Service{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		clientID:     clientID,
		clientSecret: clientSecret,
		username:     username,
	}
}
