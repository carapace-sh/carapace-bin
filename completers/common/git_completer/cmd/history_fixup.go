package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var fixupCmd = &cobra.Command{
	Use:   "fixup",
	Short: "Interactively fix up <commit>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixupCmd).Standalone()

	fixupCmd.Flags().BoolP("dry-run", "n", false, "perform a dry-run without updating any refs")
	fixupCmd.Flags().String("empty", "", "how to handle commits that become empty")
	fixupCmd.Flags().Bool("no-dry-run", false, "do not perform a dry-run without updating any refs")
	fixupCmd.Flags().Bool("reedit-message", false, "edit the commit message of the fixup")
	fixupCmd.Flags().String("update-refs", "", "control which refs should be updated")
	historyCmd.AddCommand(fixupCmd)

	carapace.Gen(fixupCmd).FlagCompletion(carapace.ActionMap{
		"empty": carapace.ActionValuesDescribed(
			"drop", "drop empty commits",
			"keep", "keep empty commits",
			"abort", "abort if any commit becomes empty",
		),
		"update-refs": carapace.ActionValuesDescribed(
			"branches", "all local branches will be rewritten",
			"head", "only the current HEAD reference will be rewritten",
		),
	})

	carapace.Gen(fixupCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(fixupCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefChanges(c.Args[0]).FilterArgs().Shift(1)
		}),
	)

	carapace.Gen(fixupCmd).DashAnyCompletion(
		carapace.ActionPositional(fixupCmd),
	)
}
