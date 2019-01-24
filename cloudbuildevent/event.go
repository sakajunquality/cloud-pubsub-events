package cloudbuildevent

import (
	"time"
)

const (
	StatusQueued  = "QUEUED"
	StatusWorking = "WORKING"
	StatusSuccess = "SUCCESS"
	StatusFailure = "FAILURE"
)

// Event is a Cloud Build event published to the Cloud Pub/Sub topic "cloud-builds"
type Event struct {
	ID         string     `json:"id"`
	ProjectID  string     `json:"projectId"`
	Status     string     `json:"status"`
	Timeout    string     `json:"timeout"`
	LogURL     string     `json:"logUrl"`
	StartTime  *time.Time `json:"startTime"`
	FinishTime *time.Time `json:"finishTime"`
	TriggerID  *string    `json:"buildTriggerId"`
	Steps      []Step     `json:"steps"`
	Source     `json:"source"`
	Artifacts  `json:"artifacts"`
}

// Source is event source
type Source struct {
	RepoSource `json:"repoSource"`
}

// RepoSource is git source...
type RepoSource struct {
	RepoName   *string `json:"repoName"`
	TagName    *string `json:"tagName"`
	BranchName *string `json:"branchName"`
}

// Artifacts includes images pushed to GCR
type Artifacts struct {
	Images []string `json:"images"`
}

// Step is the steps of build
type Step struct {
	Name   string     `json:"name"`
	Args   []string   `json:"args"`
	Timing StepTiming `json:"timing"`
	Status string     `json:"status"`
}

// StepTiming is timing inside Step
type StepTiming struct {
	StartTime  *time.Time `json:"startTime"`
	FinishTime *time.Time `json:"finishTime"`
}

// IsFinished checks if the build is finished
func (e Event) IsFinished() bool {
	return (e.FinishTime != nil)
}

// IsSuuccess checks if the build is successed
func (e Event) IsSuuccess() bool {
	return (e.Status == StatusSuccess)
}

// IsTriggerdBuild checks if the build is triggerd one
func (e Event) IsTriggerdBuild() bool {
	return (e.TriggerID != nil)
}
