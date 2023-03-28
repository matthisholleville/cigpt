package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gitlabci-gpt",
	Long:  `All software has versions. This is gitlabci-gpt's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("gitlabci-gpt version %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
