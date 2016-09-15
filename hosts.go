package theforeman

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Host struct {
	ArchitectureID      int    `json:"architecture_id"`
	ArchitectureName    string `json:"architecture_name"`
	Build               bool   `json:"build"`
	CertName            string `json:"certname"`
	Comment             string `json:"comment"`
	ComputeProfileName  string `json:"compute_profile_name"`
	ComputeResourceID   int    `json:"compute_resource_id"`
	ComputeResourceName string `json:"compute_resource_name"`
	ComputerProfileID   int    `json:"compute_profile_id"`
	DomainID            int    `json:"domain_id"`
	DomainName          string `json:"domain_name"`
	Enabled             bool   `json:"enabled"`
	EnvironmentID       int    `json:"environment_id"`
	EnvironmentName     string `json:"environment_name"`
	GlobalStatus        int    `json:"global_status"`
	GlobalStatusLabel   string `json:"global_status_label"`
	HostgroupID         int    `json:"hostgroup_id"`
	HostgroupName       string `json:"hostgroup_name"`
	HostgroupTitle      string `json:"hostgroup_title"`
	ID                  int    `json:"id"`
	IP                  string `json:"ip"`
	ImageFile           string `json:"image_file"`
	ImageID             int    `json:"image_id"`
	ImageName           string `json:"image_name"`
	LastReport          string `json:"last_report"`
	LocationID          int    `json:"location_id"`
	LocationName        string `json:"location_name"`
	MAC                 string `json:"mac"`
	Managed             bool   `json:"managed"`
	MediumID            int    `json:"medium_id"`
	MediumName          string `json:"medium_name"`
	ModelID             int    `json:"model_id"`
	ModelName           string `json:"model_name"`
	Name                string `json:"name"`
	OperatingSystemID   int    `json:"operatingsystem_id"`
	OperatingSystemName string `json:"operatingsystem_name"`
	OrganizationID      int    `json:"organization_id"`
	OrganizationName    string `json:"organization_name"`
	OwnerID             int    `json:"owner_id"`
	OwnerType           string `json:"owner_type"`
	PTableID            int    `json:"ptable_id"`
	PTableName          string `json:"ptable_name"`
	ProvisionMethod     string `json:"provision_method"`
	PuppetCAProxyID     int    `json:"puppet_ca_proxy_id"`
	PuppetProxyID       int    `json:"puppet_proxy_id"`
	PuppetStatus        int    `json:"puppet_status"`
	RealmID             int    `json:"realm_id"`
	RealmName           string `json:"realm_name"`
	SPIP                string `json:"sp_ip"`
	SPMAC               string `json:"sp_mac"`
	SPName              string `json:"sp_name"`
	SPSUbnetID          int    `json:"sp_subnet_id"`
	SubnetID            int    `json:"subnet_id"`
	SubnetName          string `json:"subnet_name"`
	UUID                string `json:"uuid"`
}

type Hosts struct {
	Total    int    `json:"total"`
	SubTotal int    `json:"subtotal"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Search   string `json:"search"`
	Results  []Host `json:"results"`
}

type HostsParams struct {
	EnvironmentID  string `url:"environment_id,omitempty"`
	HostgroupID    string `url:"hostgroup_id,omitempty"`
	LocationID     string `url:"location_id,omitempty"`
	Order          string `url:"order,omitempty"`
	OrganizationID string `url:"organization_id,omitempty"`
	Page           string `url:"page,omitempty"`
	PerPage        string `url:"per_page,omitempty"`
	Search         string `url:"search,omitempty"`
}

type HostsService struct {
	sling *sling.Sling
}

func NewHostsService(s *sling.Sling) *HostsService {
	return &HostsService{
		sling: s,
	}
}

func (s *HostsService) List(params *HostsParams) (Hosts, *http.Response, error) {
	hosts := new(Hosts)
	resp, err := s.sling.New().Path("hosts").QueryStruct(params).ReceiveSuccess(hosts)
	return *hosts, resp, err
}

func (s *HostsService) ListHostNames(params *HostsParams) ([]string, *http.Response, error) {
	hosts := new(Hosts)
	resp, err := s.sling.New().Path("hosts").QueryStruct(params).ReceiveSuccess(hosts)
	names := make([]string, 0, len(hosts.Results))
	for _, host := range hosts.Results {
		names = append(names, host.Name)
	}
	return names, resp, err
}
