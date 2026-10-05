package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configUnsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Restore a key to its default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configUnsetCmd).Standalone()

	configCmd.AddCommand(configUnsetCmd)

	carapace.Gen(configUnsetCmd).PositionalCompletion(
		actionConfigKeys(),
	)
}
