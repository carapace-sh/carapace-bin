package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var alias_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove an existing alias",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_removeCmd).Standalone()

	alias_removeCmd.Flags().BoolP("global", "g", false, "Remove from global config (in ~/.gitconfig) instead of local")
	alias_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	aliasCmd.AddCommand(alias_removeCmd)

	carapace.Gen(alias_removeCmd).PositionalCompletion(
		but.ActionAliases(but.AliasOpts{User: true}),
	)
}
