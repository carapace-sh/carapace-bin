package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dotfiles_alias_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_alias_setCmd).Standalone()

	dotfiles_alias_setCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_aliasCmd.AddCommand(dotfiles_alias_setCmd)

	carapace.Gen(dotfiles_alias_setCmd).PositionalCompletion(
		atuin.ActionAliases(),
		bridge.ActionCarapaceBin().SplitP(),
	)
}
