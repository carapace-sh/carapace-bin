package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ci_artifactCmd = &cobra.Command{
	Use:     "artifact <refName> <jobName> [flags]",
	Short:   "Download all artifacts from the last pipeline",
	Aliases: []string{"push"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_artifactCmd).Standalone()
	ci_artifactCmd.Flags().StringP("path", "p", "./", "Path to download the artifact files")
	ciCmd.AddCommand(ci_artifactCmd)

	carapace.Gen(ci_artifactCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
	})

	// TODO positional completion
}
