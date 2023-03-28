package gitlab

import (
	"bytes"

	"github.com/xanzy/go-gitlab"
)

type IGitlab interface {
	Configure(apiURL string, token string) error
	GetPipeline(projectID string, pipelineID int) (*gitlab.Pipeline, error)
	ListPipelineFailedJobs(projectID string, pipelineID int) ([]*gitlab.Job, error)
	GetTraceFile(projectID string, jobID int) (*bytes.Reader, error)
}
