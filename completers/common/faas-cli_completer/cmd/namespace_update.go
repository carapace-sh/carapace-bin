package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var namespace_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespace_updateCmd).Standalone()

	namespace_updateCmd.Flags().StringArray("annotation", nil, "Set one or more annotation (ANNOTATION=VALUE)")
	namespace_updateCmd.Flags().StringArrayP("label", "l", nil, "Set one or more label (LABEL=VALUE)")
	namespaceCmd.AddCommand(namespace_updateCmd)

	carapace.Gen(namespace_updateCmd).PositionalCompletion(
		action.ActionNamespaces(),
	)
}
