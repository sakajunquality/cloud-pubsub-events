package gcranalysisevent

const (
	EventVulnerbility = "VULNERABILITY"
)

// Event is a Cloud Build event published to the Cloud Pub/Sub topic "container-analysis-occurrences-v1beta1"
type Event struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}
