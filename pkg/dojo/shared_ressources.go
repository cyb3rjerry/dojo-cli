package dojo

type Severity int

const (
	INFO Severity = iota
	LOW
	MEDIUM
	HIGH
	CRITICAL
)

type FindingsGroupBy string

const (
	COMPONENT_NAME             FindingsGroupBy = "component_name"
	COMPONENT_NAME_AND_VERSION FindingsGroupBy = "component_name+component_version"
	FILE_PATH                  FindingsGroupBy = "file_path"
	FINDING_TITLE              FindingsGroupBy = "finding_title"
)

type ExecutiveSummary struct {
	EngagementName            string `json:"engagement_name"`
	EngagementTargetStartDate string `json:"engagement_target_start"`
	EngagementTargetEndDate   string `json:"engagement_target_end"`
	TestTypeName              string `json:"test_type_name"`
	TestTargetStart           string `json:"test_target_start"`
	TestTargetEnd             string `json:"test_target_end"`
	TestEnvironmentName       string `json:"test_environment_name"`
	TestStrategyRef           string `json:"test_strategy_ref"`
	TotalFindings             int    `json:"total_findings"`
}

type UserStub struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type NoteType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsSingle    bool   `json:"is_single"`
	IsActive    bool   `json:"is_active"`
	IsMandatory bool   `json:"is_mandatory"`
}

type NoteHistory struct {
	Id            int      `json:"id"`
	CurrentEditor UserStub `json:"current_editor"`
	NoteType      NoteType `json:"note_type"`
	Data          string   `json:"data"`
	Time          string   `json:"time"`
}

type Note struct {
	Id        int           `json:"id"`
	Author    UserStub      `json:"author"`
	Editor    UserStub      `json:"Editor"`
	History   []NoteHistory `json:"history"`
	NoteType  NoteType      `json:"note_type"`
	Entry     string        `json:"entry"`
	Date      string        `json:"date"`
	IsPrivate bool          `json:"private"`
	IsEdited  bool          `json:"edited"`
	EditTime  string        `json:"edit_time"`
}

type File struct {
	Id    int    `json:"id"`
	File  string `json:"file"`
	Title string `json:"title"`
}

type JiraIssue struct {
	Id               int    `json:"id"`
	URL              string `json:"url"`
	JiraId           string `json:"jira_id"`
	JiraKey          string `json:"jira_key"`
	JiraCreationDate string `json:"jira_creation"`
	JiraChangedDate  string `json:"jira_change"`
	Finding          int    `json:"finding"`
	Engagement       int    `json:"engagement"`
	FindingGroup     int    `json:"finding_group"`
}

type FindingGroup struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Test      int       `json:"test"`
	JiraIssue JiraIssue `json:"jira_issue"`
}

type Test struct {
	Id                   int            `json:"id"`
	Tags                 []string       `json:"tags"`
	TestTypeName         string         `json:"test_type_name"`
	FindingGroups        []FindingGroup `json:"finding_groups"`
	ScanType             string         `json:"scan_type"`
	Title                string         `json:"title"`
	Description          string         `json:"description"`
	TargetStart          string         `json:"target_start"`
	TargetEnd            string         `json:"target_end"`
	EstimatedTime        string         `json:"estimated_time"`
	ActualTime           string         `json:"actual_time"`
	PercentComplete      int            `json:"percent_complete"`
	UpdatedOn            string         `json:"updated"`
	CreatedOn            string         `json:"created"`
	Version              string         `json:"version"`
	BuildId              string         `json:"build_id"`
	CommitHash           string         `json:"commit_hash"`
	BranchTag            string         `json:"branch_tag"`
	Engagement           int            `json:"engagement"`
	Lead                 int            `json:"lead"`
	TestType             int            `json:"test_type"`
	Environment          int            `json:"environment"`
	ApiScanConfiguration int            `json:"api_scan_configuration"`
	Notes                []Note         `json:"notes"`
	Files                []File         `json:"files"`
}

type RiskAcceptance struct {
	Id                    int    `json:"id"`
	Recommendation        string `json:"recommendation"`
	Decision              string `json:"decision"`
	Path                  string `json:"path"`
	Name                  string `json:"name"`
	RecommendationDetails string `json:"recommendation_details"`
	DecisionDetails       string `json:"decision_detals"`
	AcceptedBy            string `json:"accepted_by"`
	ExpirationDate        string `json:"expiration_date"`
	ExpirationDateWarned  string `json:"expiration_date_warned"`
	ExpirationDateHandled string `json:"expiration_date_handled"`
	ReactivateExpired     bool   `json:"reactivate_expired"`
	RestartSLAExpired     bool   `json:"restart_sla_expired"`
	CreationDate          string `json:"created"`
	UpdatedDate           string `json:"updated"`
	Owner                 int    `json:"owner"` // Maps to a user id
	AcceptedFindings      []int  `json:"AcceptedFindings"`
	Notes                 []int  `json:"notes"`
}

type FindingMeta struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type FindingTestType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FindingEnvironment struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type FindingTest struct {
	Id          int                `json:"id"`
	Title       string             `json:"title"`
	TestType    FindingTestType    `json:"test_type"`
	Engagement  Engagement         `json:"engagement"`
	Environment FindingEnvironment `json:"environement"`
	BranchTag   string             `json:"branch_tag"`
	BuildId     string             `json:"build_id"`
	CommitHash  string             `json:"commit_hash"`
	Version     string             `json:"version"`
}

type FindingRelatedFields struct {
	Test FindingTest `json:"test"`
	Jira JiraIssue   `json:"jira"`
}

type VulnerabilityId struct {
	Id string `json:"vulnerability_id"`
}

type Finding struct {
	Id              int      `json:"id"`
	Tags            []string `json:"tags"`
	RequestResponse struct {
		any
	} `json:"request_response"`
	AcceptedRisks          []RiskAcceptance     `json:"accepted_risks"`
	PushToJira             bool                 `json:"push_to_jira"`
	Age                    int                  `json:"age"`
	SLADaysRemaining       int                  `json:"sla_days_remaining"`
	FindingMeta            []FindingMeta        `json:"finding_meta"`
	RelatedFields          FindingRelatedFields `json:"related_fields"`
	JiraCreation           string               `json:"jira_creation"`
	JiraChange             string               `json:"jira_change"`
	DisplayStatus          string               `json:"display_status"`
	FindingGroups          []FindingGroup       `json:"finding_groups"`
	VulnerabilityIds       []VulnerabilityId    `json:"vulnerability_ids"`
	Title                  string               `json:"title"`
	Date                   string               `json:"date"`
	SLAStartDate           string               `json:"sla_start_date"`
	CWE                    int                  `json:"cwe"`
	CVSSv3                 string               `json:"cvssv3"`
	CVSSv3Score            int                  `json:"cvssv3_score"`
	URL                    string               `json:"url"`
	Severity               string               `json:"severity"`
	Description            string               `json:"description"`
	Mitigation             string               `json:"mitigation"`
	Impact                 string               `json:"impact"`
	StepsToReproduce       string               `json:"steps_to_reproduce"`
	SeverityJustification  string               `json:"severity_justification"`
	References             string               `json:"references"`
	IsActive               bool                 `json:"active"`
	IsVerified             bool                 `json:"verified"`
	IsFalsePositive        bool                 `json:"false_p"`
	IsDuplicate            bool                 `json:"duplicate"`
	IsOutOfScope           bool                 `json:"out_of_scope"`
	IsRiskAccepted         bool                 `json:"risk_accepted"`
	IsUnderReview          bool                 `json:"under_review"`
	LastStatusUpdate       string               `json:"last_status_update"`
	IsUnderDefectReview    bool                 `json:"under_defect_review"`
	IsMitigiated           bool                 `json:"is_mitigated"`
	ThreadId               int                  `json:"thread_id"`
	MitigatedDate          string               `json:"mitigated"`
	NumericalSeverity      string               `json:"numerical_severity"`
	LastReviewed           string               `json:"last_reviewed"`
	Param                  string               `json:"param"`
	Payload                string               `json:"payload"`
	HashCode               string               `json:"hash_code"`
	Line                   int                  `json:"line"`
	FilePath               string               `json:"file_path"`
	ComponentName          string               `json:"component_name"`
	ComponentVersion       string               `json:"component_version"`
	IsStaticFinding        bool                 `json:"static_finding"`
	IsDynamicFinding       bool                 `json:"dynamic_finding"`
	CreatedOn              string               `json:"created"`
	ScannerConfidence      int                  `json:"scanner_confidence"`
	UniqueIdFromTool       string               `json:"unique_id_from_tool"`
	VulnIdFromTool         string               `json:"vuln_id_from_tool"`
	SASTSourceObject       string               `json:"sast_source_object"`
	SASTSinkObject         string               `json:"sast_sink_object"`
	SASTSourceLine         int                  `json:"sast_source_line"`
	SASTSourceFilePath     string               `json:"sast_source_file_path"`
	NumberOfOccurences     int                  `json:"nb_occurences"`
	PublishDate            string               `json:"publish_date"`
	Service                string               `json:"service"`
	PlannedRemediationDate string               `json:"planned_remediation_date"`
	Test                   int                  `json:"test"`
	DuplicateFinding       int                  `json:"duplicate_finding"`
	ReviewRequestedby      int                  `json:"review_requested_by"`        // Maps to a user id
	DefectReviewRequestBy  int                  `json:"defect_review_requested_by"` // Maps to a user id
	MitigatedBy            int                  `json:"mitigated_by"`               // Maps to a user id
	Reporter               int                  `json:"reporter"`                   // Maps to a user id
	LastReviewedBy         int                  `json:"last_reviewed_by"`           // Maps to a user id
	SonarqubeIssue         int                  `json:"sonarqube_issue"`
	Endpoints              []int                `json:"endpoints"`
	Reviewers              []int                `json:"reviewers"` // Maps to user ids
	Notes                  []Note               `json:"notes"`
	Files                  []int                `json:"files"`    // Maps to a file id
	FoundBy                []int                `json:"found_by"` // Maps to user ids
}

type FindingToNotes struct {
	FindingId int    `json:"finding_id"`
	Notes     []Note `json:"notes"`
}

type FindingNotes struct {
	NoteId int `json:"note_id"`
}
