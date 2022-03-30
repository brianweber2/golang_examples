package models

// AssetMetrics - stores Asset metrics
type AssetMetrics struct {
	AssetID string   `json:"AssetID"`
	Metrics []Metric `json:"metrics"`
}

// Metric - store metric data
type Metric struct {
	MetricName   string ` json:"metricName"`
	LastReviewed string `json:"lastReviewed"`
	LastUpdated  string ` json:"lastUpdated"`
}
