package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configHelpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display the config help menu",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configHelpCmd).Standalone()

	configCmd.AddCommand(configHelpCmd)
}
