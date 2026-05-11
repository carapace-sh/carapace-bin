package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "show"},
	Short:   "Show the value of a config parameter",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configListCmd).Standalone()

	configCmd.AddCommand(configListCmd)

	carapace.Gen(configListCmd).PositionalCompletion(
		actionConfigListParameters(),
	)
}
