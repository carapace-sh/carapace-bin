package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_currentContextCmd = &cobra.Command{
	Use:   "current-context",
	Short: "Display the current-context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_currentContextCmd).Standalone()

	configCmd.AddCommand(config_currentContextCmd)
}
