package theforeman

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Environment struct {
	Name      string `json:"name"`
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Environments struct {
	Total    int           `json:"total"`
	SubTotal int           `json:"subtotal"`
	Page     int           `json:"page"`
	PerPage  int           `json:"per_page"`
	Search   string        `json:"search"`
	Results  []Environment `json:"results"`
}

type EnvironmentsParams struct {
	PuppetClassID  string `url:"puppetclass_id,omitempty"`
	LocationID     string `url:"location_id,omitempty"`
	Order          string `url:"order,omitempty"`
	OrganizationID string `url:"organization_id,omitempty"`
	Page           string `url:"page,omitempty"`
	PerPage        string `url:"per_page,omitempty"`
	Search         string `url:"search,omitempty"`
}

type EnvironmentsService struct {
	sling *sling.Sling
}

func NewEnvironmentsService(s *sling.Sling) *EnvironmentsService {
	return &EnvironmentsService{
		sling: s,
	}
}

func (s *EnvironmentsService) List(params *EnvironmentsParams) (Environments, *http.Response, error) {
	environments := new(Environments)
	resp, err := s.sling.New().Path("environments").QueryStruct(params).ReceiveSuccess(environments)
	return *environments, resp, err
}

func (s *EnvironmentsService) ListEnvironmentNames(params *EnvironmentsParams) ([]string, *http.Response, error) {
	environments := new(Environments)
	resp, err := s.sling.New().Path("environments").QueryStruct(params).ReceiveSuccess(environments)
	names := make([]string, 0, len(environments.Results))
	for _, environment := range environments.Results {
		names = append(names, environment.Name)
	}
	return names, resp, err
}
