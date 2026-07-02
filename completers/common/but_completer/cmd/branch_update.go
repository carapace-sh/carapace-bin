package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var branch_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update your local branch with the content of its remote counterpart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_updateCmd).Standalone()

	branch_updateCmd.Flags().Bool("dry-run", false, "Preview the resulting branch state without persisting changes")
	branch_updateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_updateCmd.Flags().BoolP("interactive", "i", false, "Open the generated integration script in an editor")
	branch_updateCmd.Flags().StringP("strategy", "s", "", "Strategy to use for the integration. If no strategy is specified, we default to pull-rebase")
	branch_updateCmd.Flags().BoolP("verbose", "v", false, "Show additional dry-run details like the current divergence")
	branchCmd.AddCommand(branch_updateCmd)

	carapace.Gen(branch_updateCmd).FlagCompletion(carapace.ActionMap{
		"strategy": carapace.ActionValuesDescribed(
			"pull-rebase", "Rebuilds the branch picking first the commits on the remote, and then the commits on the local branch",
			"smart-squash", "Tries to fold matching remote work into related local commits. This is done through matching Change IDs, and falling back to pull-rebase ordering otherwise",
			"merge", "Keeps your local history and merges the remote tip into it",
			"pick-remote", "Rebuilds the branch picking only the commits on the remote",
		),
	})

	carapace.Gen(branch_updateCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
