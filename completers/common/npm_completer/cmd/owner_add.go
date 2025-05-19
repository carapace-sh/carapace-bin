package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var owner_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new user as a maintainer of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(owner_addCmd).Standalone()

	ownerCmd.AddCommand(owner_addCmd)

	carapace.Gen(owner_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionPackageNames(owner_addCmd),
	)
}
