package object

type Metrics struct {
	Totals map[string]Metric `json:"_totals"`
}

type Metric struct {
	ConfidenceHigh      float64 `json:"CONFIDENCE.HIGH"`
	ConfidenceLow       float64 `json:"CONFIDENCE.LOW"`
	ConfidenceMedium    float64 `json:"CONFIDENCE.MEDIUM"`
	ConfidenceUndefined float64 `json:"CONFIDENCE.UNDEFINED"`
	SeverityHigh        float64 `json:"SEVERITY.HIGH"`
	SeverityLow         float64 `json:"SEVERITY.LOW"`
	SeverityMedium      float64 `json:"SEVERITY.MEDIUM"`
	SeverityUndefined   float64 `json:"SEVERITY.UNDEFINED"`
	LOC                 float64 `json:"loc"`
	Nosec               float64 `json:"nosec"`
}
