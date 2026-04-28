package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Interactively split up <commit> into two commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().Bool("dry-run", false, "perform a dry-run without updating any refs")
	splitCmd.Flags().Bool("no-dry-run", false, "do not perform a dry-run without updating any refs")
	splitCmd.Flags().String("update-refs", "", "control which refs should be updated")
	historyCmd.AddCommand(splitCmd)

	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"update-refs": carapace.ActionValuesDescribed(
			"branches", "all local branches will be rewritten",
			"head", "only the current HEAD reference will be rewritten",
		),
	})

	carapace.Gen(splitCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(splitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefChanges(c.Args[0]).FilterArgs().Shift(1)
		}),
	)

	carapace.Gen(splitCmd).DashAnyCompletion(
		carapace.ActionPositional(splitCmd),
	)
}
