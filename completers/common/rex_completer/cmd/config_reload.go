package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Make the server read the config file again",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configReloadCmd).Standalone()

	configCmd.AddCommand(configReloadCmd)
}
