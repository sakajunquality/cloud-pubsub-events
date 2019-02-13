package gcrevent

const (
	ActionInsert = "INSERT"
	ActionDelete = "DELETE"
)

// Event is a Cloud Build event published to the Cloud Pub/Sub topic "gcr"
type Event struct {
	Action string  `json:"action"`
	Digest *string `json:"digest"`
	Tag    *string `json:"tag"`
}
