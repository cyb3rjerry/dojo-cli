package dojo

type ReportOptions struct {
	IncludeFindingNotes     bool `json:"include_finding_notes"`
	IncludeFindingImages    bool `json:"include_finding_images"`
	IncludeExecutiveSummary bool `json:"include_executive_summary"`
	IncludeTableOfContents  bool `json:"include_table_of_contents"`
}

type Report struct {
	ExecutiveSummary ExecutiveSummary `json:"executive_summary"`
	ProductType      ProductType      `json:"product_type"`
	Product          Product          `json:"product"`
	Engagement       Engagement       `json:"engagement"`
	ReportName       string           `json:"report_name"`
	ReportInfo       string           `json:"report_info"`
	Test             Test             `json:"test"`
	Endpoint         Endpoint         `json:"endpoint"`
	EndPointList     []Endpoint       `json:"endpoints"`
	Findings         []Finding        `json:"findings"`
	User             UserStub         `json:"user"`
	TeamName         string           `json:"team_name"`
	Title            string           `json:"title"`
	UserId           int              `json:"user_id"`
	Host             string           `json:"host"`
	FindingNotes     []FindingToNotes `json:"finding_notes"`
}
