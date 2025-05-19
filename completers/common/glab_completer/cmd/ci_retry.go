package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_retryCmd = &cobra.Command{
	Use:   "retry <job-id>",
	Short: "Retry a CI/CD job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_retryCmd).Standalone()

	ci_retryCmd.Flags().StringP("branch", "b", "", "The branch to search for the job. (default current branch)")
	ci_retryCmd.Flags().StringP("pipeline-id", "p", "", "The pipeline ID to search for the job.")
	ciCmd.AddCommand(ci_retryCmd)

	carapace.Gen(ci_retryCmd).FlagCompletion(carapace.ActionMap{
		"branch":      action.ActionBranches(ci_retryCmd),
		"pipeline-id": action.ActionPipelines(ci_retryCmd, ""),
	})

	// TODO positional completion
}
