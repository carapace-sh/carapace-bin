package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var run_rerunCmd = &cobra.Command{
	Use:   "rerun",
	Short: "Rerun a failed run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(run_rerunCmd).Standalone()
	runCmd.AddCommand(run_rerunCmd)

	carapace.Gen(run_rerunCmd).PositionalCompletion(
		action.ActionWorkflowRuns(run_rerunCmd, action.RunOpts{Failed: true}),
	)
}
