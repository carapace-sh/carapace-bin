package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rewordCmd = &cobra.Command{
	Use:   "reword",
	Short: "Rewrite the commit message of the specified commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewordCmd).Standalone()

	rewordCmd.Flags().Bool("dry-run", false, "perform a dry-run without updating any refs")
	rewordCmd.Flags().Bool("no-dry-run", false, "do not perform a dry-run without updating any refs")
	rewordCmd.Flags().String("update-refs", "", "control which refs should be updated")
	historyCmd.AddCommand(rewordCmd)

	carapace.Gen(rewordCmd).FlagCompletion(carapace.ActionMap{
		"update-refs": carapace.ActionValuesDescribed(
			"branches", "all local branches will be rewritten",
			"head", "only the current HEAD reference will be rewritten",
		),
	})

	carapace.Gen(rewordCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
