package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ci_triggerCmd = &cobra.Command{
	Use:   "trigger <job-id>",
	Short: "Trigger a manual CI/CD job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_triggerCmd).Standalone()

	ci_triggerCmd.Flags().StringP("branch", "b", "", "The branch to search for the job. (default current branch)")
	ci_triggerCmd.Flags().StringP("pipeline-id", "p", "", "The pipeline ID to search for the job.")
	ciCmd.AddCommand(ci_triggerCmd)

	carapace.Gen(ci_triggerCmd).FlagCompletion(carapace.ActionMap{
		"branch":      action.ActionBranches(ci_triggerCmd),
		"pipeline-id": action.ActionPipelines(ci_triggerCmd, ""),
	})

	// TODO complete job ids
}
