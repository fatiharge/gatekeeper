package gatekeeper

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	ExternalURL string `json:"externalURL,omitempty"`
	AuthHeader  string `json:"authHeader,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		AuthHeader: "Authorization", 
	}
}

type Gatekeeper struct {
	next        http.Handler
	name        string
	externalURL string
	authHeader  string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Gatekeeper{
		next:        next,
		name:        name,
		externalURL: config.ExternalURL,
		authHeader:  config.AuthHeader,
	}, nil
}

func (g *Gatekeeper) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Get the Authorization header
	authHeader := req.Header.Get(g.authHeader)

	if authHeader == "" {
		g.next.ServeHTTP(rw, req)
		return
	}

	if g.externalURL != "" {
		client := &http.Client{Timeout: 3 * time.Second}
		request, err := http.NewRequest("GET", g.externalURL, nil)
		if err != nil {
			http.Error(rw, "Failed to create request", http.StatusInternalServerError)
			return
		}

		request.Header.Add(g.authHeader, authHeader)

		resp, err := client.Do(request)
		if err != nil {
			http.Error(rw, "Failed to reach external validation service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(rw, "Unauthorized", http.StatusForbidden)
			return
		}
	}

	g.next.ServeHTTP(rw, req)
}
