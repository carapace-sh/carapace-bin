package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configSetCmd).Standalone()

	configCmd.AddCommand(configSetCmd)

	carapace.Gen(configSetCmd).PositionalCompletion(
		actionConfigKeys(),
	)
}
