package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Update applied branches and worktrees onto the latest target",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().BoolP("check", "c", false, "Only check whether the update would apply cleanly, without updating")
	pullCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pullCmd.Flags().StringSlice("update", []string{"worktrees", "workspace"}, "What to rebase onto the target, comma-separated")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"update": carapace.ActionValuesDescribed(
			"workspace", "The applied branches, along with the linked worktrees based on them",
			"worktrees", "The linked worktrees that are based on the target rather than on an applied branch",
		),
	})
}
