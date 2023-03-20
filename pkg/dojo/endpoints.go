package dojo

type EndpointsList struct {
	Count     int        `json:"count"`
	Next      string     `json:"next"`
	Previous  string     `json:"previous"`
	Endpoints []Endpoint `json:"results"`
}

type Endpoint struct {
	Id             int      `json:"id"`
	Tags           []string `json:"tags"`
	Protocol       string   `json:"protocol"`
	UserInfo       string   `json:"userinfo"`
	Host           string   `json:"host"`
	Port           int      `json:"port"`
	Path           string   `json:"path"`
	Query          string   `json:"query"`
	Fragment       string   `json:"fragment"`
	Product        string   `json:"product"`
	EndpointParams []int    `json:"endpoint_params"`
	Findings       []int    `json:"findings"`
}
