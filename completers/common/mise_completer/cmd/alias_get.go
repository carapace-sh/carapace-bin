package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show a configured version alias for a tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_getCmd).Standalone()

	aliasCmd.AddCommand(alias_getCmd)

	carapace.Gen(alias_getCmd).PositionalCompletion(
		action.ActionTools(),
		action.ActionAliases(),
	)
}
