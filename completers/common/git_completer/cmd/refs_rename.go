package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var refs_renameCmd = &cobra.Command{
	Use:   "rename <old-ref> <new-ref>",
	Short: "Rename a reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_renameCmd).Standalone()

	refs_renameCmd.Flags().String("message", "", "reason of the update")
	refsCmd.AddCommand(refs_renameCmd)

	carapace.Gen(refs_renameCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
