package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var namespace_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove existing namespace",
	Aliases: []string{"rm", "delete"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespace_removeCmd).Standalone()

	namespaceCmd.AddCommand(namespace_removeCmd)

	carapace.Gen(namespace_removeCmd).PositionalCompletion(
		action.ActionNamespaces(),
	)
}
