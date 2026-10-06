package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var dropCmd = &cobra.Command{
	Use:   "drop <commit>",
	Short: "Remove <commit>, replaying its descendants onto its parent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dropCmd).Standalone()

	dropCmd.Flags().BoolP("dry-run", "n", false, "perform a dry-run without updating any refs")
	dropCmd.Flags().String("empty", "", "how to handle descendants that become empty")
	dropCmd.Flags().Bool("no-dry-run", false, "do not perform a dry-run without updating any refs")
	dropCmd.Flags().String("update-refs", "", "control which refs should be updated")
	historyCmd.AddCommand(dropCmd)

	carapace.Gen(dropCmd).FlagCompletion(carapace.ActionMap{
		"empty": carapace.ActionValuesDescribed(
			"drop", "drop commits that become empty",
			"keep", "keep commits that become empty",
			"abort", "abort if any commit becomes empty",
		),
		"update-refs": git.ActionUpdateRefsModes(),
	})

	carapace.Gen(dropCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(dropCmd).DashAnyCompletion(
		carapace.ActionPositional(dropCmd),
	)
}
