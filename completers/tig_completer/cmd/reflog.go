package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var reflogCmd = &cobra.Command{
	Use:   "reflog",
	Short: "Manage reflog information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflogCmd).Standalone()

	rootCmd.AddCommand(reflogCmd)

	carapace.Gen(reflogCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRefs(git.RefOptionDefault).Chdir(grepCmd.Root().Flag("C").Value.String())
		}),
	)
}
