package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var owner_rmCmd = &cobra.Command{
	Use:     "rm",
	Short:   "Remove a user from the package owner list",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(owner_rmCmd).Standalone()

	ownerCmd.AddCommand(owner_rmCmd)

	carapace.Gen(owner_rmCmd).PositionalCompletion(
		action.ActionPackageNames(owner_rmCmd),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return npm.ActionOwners(c.Args[0])
		}),
	)
}
