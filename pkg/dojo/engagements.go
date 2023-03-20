package dojo

type EngagementList struct {
	Count       int          `json:"count"`
	Next        string       `json:"next"`
	Previous    string       `json:"previous"`
	Engagements []Engagement `json:"results"`
}

type Engagement struct {
	Id                          int      `json:"id"`
	Tags                        []string `json:"tags"`
	Name                        string   `json:"name"`
	Description                 string   `json:"description"`
	Version                     string   `json:"version"`
	FirstContacted              string   `json:"first_contacted"`
	TargetStart                 string   `json:"target_start"`
	TargetEnd                   string   `json:"target_end"`
	Reason                      string   `json:"reason"`
	Updated                     string   `json:"updated"`
	Created                     string   `json:"created"`
	IsActive                    bool     `json:"active"`
	Tracker                     string   `json:"tracker"`
	TestStrategy                string   `json:"test_strategy"`
	ThreatModel                 string   `json:"threat_model"`
	IsApiTest                   bool     `json:"api_test"`
	IsPentest                   bool     `json:"pen_test"`
	IsChecklist                 bool     `json:"check_list"`
	Status                      string   `json:"status"`
	Progress                    string   `json:"progress"`
	ThreatModelPath             string   `json:"tmodel_path"`
	IsDoneTesting               bool     `json:"done_testing"`
	EngagementType              string   `json:"engagement_type"`
	BuildId                     string   `json:"build_id"`
	CommitHash                  string   `json:"commit_hash"`
	BranchTag                   string   `json:"branch_tag"`
	SourceCodeManagementURI     string   `json:"source_code_management_uri"`
	IsDeduplicationOnEngagement bool     `json:"deduplication_on_engagement"`
	Lead                        int      `json:"lead"`        // Matches a user id
	Requester                   int      `json:"requester"`   // Matches a user id
	Preset                      int      `json:"preset"`      // Matches a preset id
	ReportType                  int      `json:"report_type"` // Matches report type id
	Product                     int      `json:"product"`     // Matches a product id
	BuildServer                 int      `json:"build_server"`
	SourceCodeManagementServer  int      `json:"source_code_management_server"`
	OrchestrationEngine         int      `json:"orchestration_engine"`
	Notes                       []Note   `json:"notes"`
	Files                       []File   `json:"files"`
	RiskAcceptance              []int    `json:"risk_acceptance"`
}
