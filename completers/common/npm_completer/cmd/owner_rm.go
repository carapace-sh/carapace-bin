package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var owner_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a user from the package owner list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(owner_rmCmd).Standalone()

	ownerCmd.AddCommand(owner_rmCmd)

	carapace.Gen(owner_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionPackageNames(owner_rmCmd),
	)
}
