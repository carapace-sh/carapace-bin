package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_cancelCmd = &cobra.Command{
	Use:   "cancel [<run-id>]",
	Short: "Cancel a workflow run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_cancelCmd).Standalone()

	run_cancelCmd.Flags().Bool("force", false, "Force cancel a workflow run")
	runCmd.AddCommand(run_cancelCmd)

	carapace.Gen(run_cancelCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_cancelCmd, action.RunOpts{InProgress: true}),
	)
}
