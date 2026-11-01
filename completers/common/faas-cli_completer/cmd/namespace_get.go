package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var namespace_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespace_getCmd).Standalone()

	namespaceCmd.AddCommand(namespace_getCmd)

	carapace.Gen(namespace_getCmd).PositionalCompletion(
		action.ActionNamespaces(),
	)
}
