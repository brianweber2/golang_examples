package models

type defectOverSla struct {
	DaysOverSla float64 `dynamo:"daysOverSla" json:"daysOverSla"`
	IssueKey    string  `dynamo:"issueKey" json:"issueKey"`
}

type sponsorList struct {
	DevManagerEmail        string `dynamo:"email" json:"devManagerEmail"`
	DevManagerName         string `dynamo:"name" json:"devManager"`
	LevelTwoManager        string `dynamo:"levelTwoManager" json"levelTwoManager"`
	LevelTwoManagerEmail   string `dynamo:"levelTwoManagerEmail" json"levelTwoManagerEmail"`
	LevelThreeManager      string `dynamo:"levelThreeManager" json"levelThreeManager"`
	LevelThreeManagerEmail string `dynamo:"levelThreeManagerEmail" json"levelThreeManagerEmail"`
	LevelFourManager       string `dynamo:"levelFourManager" json"levelFourManager"`
	LevelFourManagerEmail  string `dynamo:"levelFourManagerEmail" json"levelFourManagerEmail"`
	LevelFiveManager       string `dynamo:"levelFiveManager" json"levelFiveManager"`
	LevelFiveManagerEmail  string `dynamo:"levelFiveManagerEmail" json"levelFiveManagerEmail"`
	LevelSixManager        string `dynamo:"levelSixManager" json"levelSixManager"`
	LevelSixManagerEmail   string `dynamo:"levelSixManagerEmail" json"levelSixManagerEmail"`
	LevelSevenManager      string `dynamo:"levelSevenManager" json"levelSevenManager"`
	LevelSevenManagerEmail string `dynamo:"levelSevenManagerEmail" json"levelSevenManagerEmail"`
	// Six, Seven, Eight
}

type defectsOverSla struct {
	P0 []defectOverSla `dynamo:"p0" json:"p0"`
	P1 []defectOverSla `dynamo:"p1" json:"p1"`
	P2 []defectOverSla `dynamo:"p2" json:"p2"`
	P3 []defectOverSla `dynamo:"p3" json:"p3"`
}

type githubBranchProtection struct {
	hasTestCoverage          bool `dynamo:"hasTestcoverage" json:"hasTestCoverage"`
	hasContinuousIntegration bool `dynamo:"hasContinuousIntegration" json:"hasContinuousIntegration"`
	mergeCodeCoverage        bool `dynamo:"mergeCodeCoverage" json:"mergeCodeCoverage"`
	branchName               bool `dynamo:"branchName" json:"branchName"`
	requiresCodeReviews      bool `dynamo:"requiresCodeReviews" json:"requiresCodeReviews"`
}

type githubRepo struct {
	branchProtection            githubBranchProtection `dynamo:"githubBranchProtection" json:"githubBranchProtection"`
	invalidOrInaccessibleBranch bool                   `dynamo:"invalidOrInaccessibleBranch" json:"invalidOrInaccessibleBranch"`
	repoName                    string                 `dynamo:"repoName" json:"repoName"`
	repoOrg                     string                 `dynamo:"repoOrg" json:"repoOrg"`
}

type metrics struct {
	hasCPUMetrics    bool `dynamo:hasCPUMetrics" json:"hasCPUMetrics"`
	hasMemoryMetrics bool `dynamo:hasMemoryMetrics json:"hasMemoryMetrics"`
}

type outageResponse struct {
	year                 int    `dynamo:"metricYear" json:"year"`
	metricBasis          string `dynamo:"metricBasis" json:"metricBasis"`
	month                int    `dynamo:"metricMonth" json:"month"`
	MttrMins             string `dynamo:"MTTRmins" json:"mttrMins"`
	MttdMins             string `dynamo:"MTTDmins" json:"mttdMins"`
	timeToDetectMins     int    `dynamo:"timeToDetectMins" json:"timeToDetectMins"`
	timeToRepairMins     int    `dynamo:"timeToRepairMins" json:"timeToRepairMins"`
	TotalNumberOfOutages int    `dynamo:"totalNumberOfOutages" json:"totalNumberOfOutages"`
}

type Asset struct {
	AssetID                string         `dynamo:"AssetID" json:"assetID"`
	ServiceName            string         `dynamo:"Service Name" json:"assetName"`
	SnowCI                 string         `dynamo:"SnowCI"`
	Sponsor                sponsorList    `dynamo:"Sponsor" json:"sponsor"`
	Micrometer             string         `dynamo:"Micrometer Enabled" json:"micrometer"`
	SnowID                 string         `dynamo:"SnowID" json:"snowID"`
	DRStrategy             string         `dynamo:"DRStrategy" json:"drStrategy"`
	AssetVerified          string         `dynamo:"Asset Verified" json:"assetVerified"`
	AssetAliasTagWF        string         `dynamo:"ASSET ALIAS TAG IN WF METRICS" json:"assetAliasTagWF"`
	DefectsOverSla         defectsOverSla `dynamo:"defectsOverSla" json:"defectsOverSla"`
	SecurityDefectsOverSla defectsOverSla `dynamo:"securityDefectsOverSla" json:"securityDefectsOverSla"`
	AssetAlias             string         `dynamo:"Asset Alias"`
	GoldenSignals          []string       `dynamo:"GoldenSignals" json:"goldenSignals"`
	Lifecycle              string         `dynamo:"Lifecycle" json:"lifecycle"`
	JiraProject            string         `dynamo:"Jira project" json:"jiraProject"`
	BU                     string         `dynamo:"Business Unit" json:"bu"`
	Metrics                metrics        `dynamo:"metrics" json:"metrics"`
	Correto                string         `dynamo:"Correto" json:"correto"`
	Tier                   string         `dynamo:"Tier" json:"tier"`
	EscalationId           string         `dynamo:"PagerDuty escalationId" json:"escalationID"`
	GithubRepo             githubRepo     `dynamo:"githubRepo" json:"githubRepo"`
	Outages                outageResponse `dynamo:"OutageResponse" json:"outageResponse"`
	SystemUpTime           string         `dynamo:"systemUpTime" json:"systemUpTime"`
}
