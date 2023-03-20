package dojo

type ProductsList struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Products []Product `json:"results"`
}

type ProductMeta struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductType struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	KeyProduct          bool   `json:"key_product"`
	UpdatedOn           string `json:"updated"`
	CreatedOn           string `json:"created"`
	Members             []int  `json:"members"`
	AuthorizationGroups []int  `json:"authorization_groups"`
}

type Product struct {
	Id                             int           `json:"int"`
	FindingsCount                  int           `json:"findings_count"`
	FindingsList                   []int         `json:"findings_list"`
	Tags                           []string      `json:"tags"`
	ProductMetas                   []ProductMeta `json:"product_meta"`
	Name                           string        `json:"name"`
	Description                    string        `json:"description"`
	CreationDate                   string        `json:"created"`
	ProdNumericGrade               int           `json:"prod_numeric_grade"`
	BusinessCriticality            string        `json:"business_criticality"` // Very High | High | Medium | Low | Very Low | None
	Platform                       string        `json:"platform"`             // Web Service | Desktop | IoT | Mobile | Web
	Lifecycle                      string        `json:"lifecycle"`            // Construction | Production | Retirement
	Origin                         string        `json:"origin"`               // Third-party | Purchased | Contractor | Internal | Open Source | Outsourced
	UserRecords                    int           `json:"user_records"`
	Revenue                        string        `json:"revenue"`
	HasExternalAudience            bool          `json:"external_audience"`
	IsInternetAccessible           bool          `json:"internet_accessible"`
	HasSimpleRiskAcceptanceEnabled bool          `json:"enable_simple_risk_acceptance"`
	HasFullRiskAcceptanceEnabled   bool          `json:"enable_full_risk_acceptance"`
	ProductManager                 int           `json:"product_manager"`   // Maps to a user id
	TechnicalContact               int           `json:"technical_contact"` // Maps to a user id
	TeamManager                    int           `json:"team_manager"`
	ProdType                       int           `json:"prod_type"`
	SLA_configuration              int           `json:"sla_configuration"`
	Members                        []int         `json:"members"`
	AuthorizationGroups            []int         `json:"authorization_groups"`
	Regulations                    []int         `json:"regulations"`
}
