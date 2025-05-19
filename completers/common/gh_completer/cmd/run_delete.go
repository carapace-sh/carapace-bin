package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_deleteCmd = &cobra.Command{
	Use:   "delete [<run-id>]",
	Short: "Delete a workflow run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_deleteCmd).Standalone()

	runCmd.AddCommand(run_deleteCmd)

	carapace.Gen(run_deleteCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_deleteCmd, action.RunOpts{All: true}),
	)
}
