package object

type ScanResponse struct {
	Errors      []interface{} `json:"errors"`
	GeneratedAt string        `json:"generated_at"`
	Metrics     map[string]Metric   `json:"metrics"`
    Results     []ScanResult  `json:"results"`
}
