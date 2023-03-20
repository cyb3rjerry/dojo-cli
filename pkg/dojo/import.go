package dojo

type StatusStatistics struct {
	Active        int `json:"active"`
	Verified      int `json:"verified"`
	Duplicate     int `json:"duplicate"`
	FalsePositive int `json:"false_p"`
	OutOfScope    int `json:"out_of_scope"`
	IsMitigated   int `json:"is_mitigated"`
	RiskAccepted  int `json:"risk_accepted"`
	Total         int `json:"total"`
}

type SeverityStatusStatistics struct {
	Info     StatusStatistics `json:"info"`
	Low      StatusStatistics `json:"low"`
	Medium   StatusStatistics `json:"medium"`
	High     StatusStatistics `json:"High"`
	Critical StatusStatistics `json:"critical"`
	Total    StatusStatistics `json:"total"`
}

type DeltaStatistics struct {
	Created       SeverityStatusStatistics `json:"created"`
	Closed        SeverityStatusStatistics `json:"closed"`
	Reactivated   SeverityStatusStatistics `json:"reactivated"`
	LeftUntouched SeverityStatusStatistics `json:"left untouched"`
}

type ImportStatistics struct {
	Before SeverityStatusStatistics `json:"before"`
	Delta  DeltaStatistics          `json:"delta"`
	After  SeverityStatusStatistics `json:"after"`
}

type Import struct {
	Date                               string           `json:"scan_date"`
	MinimumSeverity                    string           `json:"minimum_severity"`
	Active                             Severity         `json:"active"`
	IsVerified                         bool             `json:"verified"`
	ScanType                           string           `json:"scan_type"`
	EndpointsToAdd                     int              `json:"endpoint_to_add"` // Maps to an endpoint id
	File                               string           `json:"file"`
	ProductTypeName                    string           `json:"product_type_name"`
	ProductName                        string           `json:"product_name"`
	EngagementName                     string           `json:"engagement_name"`
	EngagementEndDate                  string           `json:"engagement_end_date"`
	SourceCodeManagementURI            string           `json:"source_code_management_uri"`
	Engagement                         int              `json:"engagement"` // Maps to an engagement id
	TestTitle                          string           `json:"test_title"`
	AutoCreateContext                  bool             `json:"auto_create_context"`
	DeduplicationOnEngagement          bool             `json:"deduplication_on_engagement"`
	Lead                               int              `json:"lead"` // Maps to a user id
	Tags                               []string         `json:"tags"`
	CloseOldFindings                   bool             `json:"close_old_findings"`
	CloseOldFindingsProductScope       bool             `json:"close_old_findings_product_scope"`
	PushToJira                         bool             `json:"push_to_jira"`
	Environment                        string           `json:"environment"`
	Version                            string           `json:"version"`
	BuildId                            string           `json:"build_id"`
	BranchTag                          string           `json:"branch_tag"`
	CommitHash                         string           `json:"commit_hash"`
	ApiScanConfig                      int              `json:"api_scan_configuration"` // Maps to an API Scan Config id
	Service                            string           `json:"service"`
	GroupBy                            FindingsGroupBy  `json:"group_by"`
	CreateFindingsGroupsForAllFindings bool             `json:"create_finding_groups_for_all_findings"`
	Test                               int              `json:"test"`
	TestId                             int              `json:"test_id"`
	EngagementId                       int              `json:"engagement_id"`
	ProductId                          int              `json:"product_id"`
	ProductTypeId                      int              `json:"product_type_id"`
	Statistics                         ImportStatistics `json:"statistics"`
}
