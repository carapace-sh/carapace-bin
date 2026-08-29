package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginValidateCmd).Standalone()
	pluginCmd.AddCommand(pluginValidateCmd)

	carapace.Gen(pluginValidateCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
