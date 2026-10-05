package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the value of one key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configGetCmd).Standalone()

	configCmd.AddCommand(configGetCmd)

	carapace.Gen(configGetCmd).PositionalCompletion(
		actionConfigKeys(),
	)
}
