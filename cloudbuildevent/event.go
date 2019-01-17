package cloudbuildevent

import (
	"time"
)

const (
	statusQueued  = "QUEUED"
	statusWorking = "WORKING"
	statusSuccess = "SUCCESS"
	statusFailure = "FAILURE"
)

// Event is Cloud Build events published to Cloud Pub/Sub the topic "cloud-builds"
type Event struct {
	ID         string     `json:"id"`
	ProjectID  string     `json:"projectId"`
	Status     string     `json:"status"`
	Timeout    string     `json:"timeout"`
	LogURL     string     `json:"logUrl"`
	StartTime  *time.Time `json:"startTime"`
	FinishTime *time.Time `json:"finishTime"`

	TriggerID string `json:"buildTriggerId"`

	Source `json:"source"`
	Steps  []Step `json:"steps"`

	Artifacts `json:"artifacts"`
}

type Source struct {
	RepoSource `json:"repoSource"`
}

type RepoSource struct {
	RepoName   *string `json:"repoName"`
	TagName    *string `json:"tagName"`
	BranchName *string `json:"branchName"`
}

type Artifacts struct {
	Images []string `json:"images"`
}

type Step struct {
	Name   string   `json:"name"`
	Args   []string `json:"args"`
	Timing []string `json:"timing"`
	Status string   `json:"status"`
}

func (e Event) isFinished() bool {
	return (e.FinishTime != nil)
}

func (e Event) isSuuccess() bool {
	return (e.Status == statusSuccess)
}

func (e Event) isTriggerdBuild() bool {
	return (e.TriggerID != "")
}
