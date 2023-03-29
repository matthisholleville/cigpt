package cmd

import (
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cigpt",
	Long:  `All software has versions. This is cigpt's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("cigpt version %s", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
