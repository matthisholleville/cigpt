package analyze

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/cigpt-ai/cigpt/pkg/ai"
	"github.com/cigpt-ai/cigpt/pkg/gitlab"
	"github.com/cigpt-ai/cigpt/pkg/util"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	backend    string
	language   string
	projectID  string
	pipelineID int
)

var AnalyzeCmd = &cobra.Command{
	Use:     "analyze",
	Aliases: []string{"analyze"},
	Short:   "This command will analyze the error logs of a GitlabCI pipeline.",
	Long:    "This command will analyze the error logs of a GitlabCI pipeline and provide you with a list of issues.",
	Run: func(cmd *cobra.Command, args []string) {
		// get backend from file
		backendType := viper.GetString("backend_type")
		if backendType == "" {
			color.Red("No backend set. Please run cigpt auth")
			os.Exit(1)
		}
		// override the default backend if a flag is provided
		if backend != "" {
			backendType = backend
		}
		// get the token with viper
		token := viper.GetString(fmt.Sprintf("%s_key", backendType))
		// check if nil
		if token == "" {
			color.Red("No %s key set. Please run cigpt auth", backendType)
			os.Exit(1)
		}

		//gitlab
		gToken := viper.GetString("gitlab_token")
		if token == "" {
			color.Red("No gitlab token set. Please run cigpt auth")
			os.Exit(1)
		}
		gApiURL := viper.GetString("gitlab_api_url")
		if token == "" {
			color.Red("No gitlab api url set. Please run cigpt auth")
			os.Exit(1)
		}

		//aiClient
		var aiClient ai.IAI
		switch backendType {
		case "openai":
			aiClient = &ai.OpenAIClient{}
			if err := aiClient.Configure(token, language); err != nil {
				color.Red("Error: %v", err)
				os.Exit(1)
			}
		default:
			color.Red("Backend not supported")
			os.Exit(1)
		}

		//gitlabClient
		var gitlabClient gitlab.IGitlab
		gitlabClient = &gitlab.GitlabClient{}
		if err := gitlabClient.Configure(gApiURL, gToken); err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		if projectID == "" || pipelineID == 0 {
			color.Red("Error: pipeline ID and project ID are required")
			os.Exit(1)
		}

		pipeline, err := gitlabClient.GetPipeline(projectID, pipelineID)
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		color.Blue("Pipeline URL : " + pipeline.WebURL)

		jobs, err := gitlabClient.ListPipelineFailedJobs(projectID, pipelineID)
		if err != nil {
			color.Red("Error: %v", err)
			os.Exit(1)
		}

		if len(jobs) == 0 {
			color.Yellow("No job errors found for this pipeline.")
			os.Exit(0)
		}

		for _, job := range jobs {
			logs, err := gitlabClient.GetTraceFile(projectID, job.ID)
			if err != nil {
				color.Red("Error: %v", err)
				os.Exit(1)
			}
			color.Yellow("Analyze error job : %s", job.Name)

			content, err := ioutil.ReadAll(logs)
			if err != nil {
				color.Red("Error: %v", err)
				os.Exit(1)
			}

			inputText := util.FilterLogs(string(content))
			inputText = util.TrimLogs(inputText)

			ctx := context.Background()
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Suffix = " Loading..."
			s.Start()
			explanation, err := aiClient.GetCompletion(ctx, inputText)
			if err != nil {
				color.Red("Error: %v", err)
				os.Exit(1)
			}
			s.Stop()

			color.Green("Solution : %s", explanation)
		}

	},
}

func init() {
	// add flag for pipeline ID
	AnalyzeCmd.Flags().IntVar(&pipelineID, "pipeline-id", 0, "Gitlab PipelineID")
	// add flag for backend
	AnalyzeCmd.Flags().StringVarP(&backend, "backend", "b", "openai", "Backend AI provider")
	// gitlab project ID
	AnalyzeCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Gitlab ProjectID")
	// add language options for output
	AnalyzeCmd.Flags().StringVarP(&language, "language", "l", "english", "Languages to use for AI (e.g. 'English', 'Spanish', 'French', 'German', 'Italian', 'Portuguese', 'Dutch', 'Russian', 'Chinese', 'Japanese', 'Korean')")
}
