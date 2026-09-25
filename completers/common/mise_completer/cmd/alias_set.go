package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Add/update an alias for a tool/backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_setCmd).Standalone()

	aliasCmd.AddCommand(alias_setCmd)

	carapace.Gen(alias_setCmd).PositionalCompletion(
		action.ActionTools(),
		carapace.ActionValues(),
	)
}
