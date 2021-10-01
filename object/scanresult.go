package object

type ScanResult struct {
	Code            string  `json:"code"`
	Filename        string  `json:"filename"`
	IssueConfidence string  `json:"issue_confidence"`
	IssueSeverity   string  `json:"issue_severity"`
	IssueText       string  `json:"issue_text"`
	LineNumber      int64   `json:"line_number"`
	LineRange       []int64 `json:"line_range"`
	MoreInfo        string  `json:"more_info"`
	TestID          string  `json:"test_id"`
	TestName        string  `json:"test_name"`
}
