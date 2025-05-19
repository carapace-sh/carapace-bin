package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/toit.pkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var registry_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_removeCmd).Standalone()
	registryCmd.AddCommand(registry_removeCmd)

	carapace.Gen(registry_removeCmd).PositionalCompletion(
		action.ActionRegistries(),
	)
}
