package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interact with the global Skaffold config file (defaults to `$HOME/.skaffold/config`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	applyCommonFlagCompletion(configCmd)
	rootCmd.AddCommand(configCmd)
}
