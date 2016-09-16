package theforeman

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Config struct {
	// Base URL for TheForeman API
	BaseURL string

	HttpClient *http.Client

	Username string

	Password string
}

type Client struct {
	EnvironmentsService *EnvironmentsService
	HostsService        *HostsService
}

// NewClient returns a new Client
func NewClient(conf *Config, httpClient *http.Client) *Client {
	s := sling.New().Client(conf.HttpClient).Base(conf.BaseURL).SetBasicAuth(conf.Username, conf.Password)
	return &Client{
		EnvironmentsService: NewEnvironmentsService(s),
		HostsService:        NewHostsService(s),
	}
}
