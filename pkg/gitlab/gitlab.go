package gitlab

import (
	"bytes"
	"errors"

	"github.com/xanzy/go-gitlab"
)

type GitlabClient struct {
	client *gitlab.Client
}

func (g *GitlabClient) Configure(apiURL string, token string) error {
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(apiURL))
	if err != nil {
		return errors.New("error creating gitlab client.")
	}
	g.client = client
	return nil
}

func (g *GitlabClient) GetPipeline(projectID string, pipelineID int) (*gitlab.Pipeline, error) {
	pipeline, _, err := g.client.Pipelines.GetPipeline(projectID, pipelineID)
	if err != nil {
		return &gitlab.Pipeline{}, errors.New("cannot get pipeline details.")
	}
	return pipeline, nil
}

func (g *GitlabClient) ListPipelineFailedJobs(projectID string, pipelineID int) ([]*gitlab.Job, error) {
	options := &gitlab.ListJobsOptions{
		Scope: &[]gitlab.BuildStateValue{
			"failed",
		},
	}
	jobs, _, err := g.client.Jobs.ListPipelineJobs(projectID, pipelineID, options)
	if err != nil {
		return []*gitlab.Job{}, errors.New("cannot list pipeline jobs.")
	}
	return jobs, nil
}

func (g *GitlabClient) GetTraceFile(projectID string, jobID int) (*bytes.Reader, error) {
	logs, _, err := g.client.Jobs.GetTraceFile(projectID, jobID)
	if err != nil {
		return &bytes.Reader{}, errors.New("cannot get trace file.")
	}
	return logs, nil
}
