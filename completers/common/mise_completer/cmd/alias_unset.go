package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Clear an alias for a tool/backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_unsetCmd).Standalone()

	aliasCmd.AddCommand(alias_unsetCmd)

	carapace.Gen(alias_unsetCmd).PositionalCompletion(
		action.ActionTools(),
		action.ActionAliases(),
	)
}
