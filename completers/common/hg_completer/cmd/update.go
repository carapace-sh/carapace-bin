package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "update working directory (or switch revisions)",
	Aliases: []string{"up", "checkout", "co"},
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().BoolP("check", "c", false, "require clean working directory")
	updateCmd.Flags().BoolP("clean", "C", false, "discard uncommitted changes (no backup)")
	updateCmd.Flags().StringP("date", "d", "", "tipmost revision matching date")
	updateCmd.Flags().BoolP("merge", "m", false, "merge uncommitted changes")
	updateCmd.Flags().StringP("rev", "r", "", "revision")
	updateCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"rev":  hg.ActionRevisions(),
		"tool": hg.ActionMergeTools(),
	})

	carapace.Gen(updateCmd).PositionalCompletion(
		hg.ActionRevisions(),
	)
}
