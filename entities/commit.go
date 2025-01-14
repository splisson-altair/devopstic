package entities

import (
	"time"
)

type Commit struct {
	Model
	State              string
	CommitTime         time.Time
	SubmitTime         time.Time
	ApprovalTime       time.Time
	DeploymentTime     time.Time
	ReviewLeadTime     int64
	DeploymentLeadTime int64
	TotalLeadTime      int64  // ReviewLeadTime + DeploymentLeadTime
	PipelineId         string // external unique id of the component/pipeline tracked
	CommitId           string // external unique id of the commit
	PullRequestId      int64
}

const (
	COMMIT_STATE_COMMITTED = "committed"
	COMMIT_STATE_SUBMITTED = "submitted"
	COMMIT_STATE_APPROVED  = "approved"
	COMMIT_STATE_DEPLOYED  = "deployed"
)
