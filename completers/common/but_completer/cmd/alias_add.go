package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var alias_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_addCmd).Standalone()

	alias_addCmd.Flags().BoolP("global", "g", false, "Store the alias globally (in ~/.gitconfig) instead of locally")
	alias_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	aliasCmd.AddCommand(alias_addCmd)

	carapace.Gen(alias_addCmd).PositionalCompletion(
		but.ActionAliases(but.AliasOpts{}.Default()),
		bridge.ActionCarapaceBin("but").Split(),
	)
}
