package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Updates all applied branches to be up to date with the target branch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolP("check", "c", false, "Only check the status without updating (equivalent to the old but base check)")
	pullCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pullCmd.Flags().StringP("strategy", "s", "pull-rebase", "Strategy to use for the integration. If no strategy is specified, we default to pull-rebase")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"strategy": carapace.ActionValuesDescribed(
			"pull-rebase", "Rebuilds the branch picking first the commits on the remote, and then the commits on the local branch",
			"smart-squash", "Tries to fold matching remote work into related local commits. This is done through matching Change IDs, and falling back to pull-rebase ordering otherwise",
			"merge", "Keeps your local history and merges the remote tip into it",
			"pick-remote", "Rebuilds the branch picking only the commits on the remote",
		),
	})
}
