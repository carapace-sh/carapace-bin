package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_deleteCmd = &cobra.Command{
	Use:   "delete <alias name> [flags]",
	Short: "Delete an alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_deleteCmd).Standalone()

	aliasCmd.AddCommand(alias_deleteCmd)

	carapace.Gen(alias_deleteCmd).PositionalCompletion(
		action.ActionAliases(),
	)
}
