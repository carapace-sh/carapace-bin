package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var reflog_existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Check whether a ref has a reflog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_existsCmd).Standalone()

	reflogCmd.AddCommand(reflog_existsCmd)

	carapace.Gen(reflogCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
