package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var refs_existsCmd = &cobra.Command{
	Use:   "exists",
	Short: "Check if a reference exists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(refs_existsCmd).Standalone()

	refsCmd.AddCommand(refs_existsCmd)

	carapace.Gen(refs_existsCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
