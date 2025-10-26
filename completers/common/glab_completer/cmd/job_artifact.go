package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var job_artifactCmd = &cobra.Command{
	Use:     "artifact <refName> <jobName> [flags]",
	Short:   "Download all artifacts from the last pipeline.",
	Aliases: []string{"push"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(job_artifactCmd).Standalone()

	job_artifactCmd.Flags().BoolP("list-paths", "l", false, "Print the paths of downloaded artifacts.")
	job_artifactCmd.Flags().StringP("path", "p", "", "Path to download the artifact files.")
	jobCmd.AddCommand(job_artifactCmd)

	carapace.Gen(job_artifactCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
