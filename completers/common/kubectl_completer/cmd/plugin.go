package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin [flags]",
	Short: "Provides utilities for interacting with plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	rootCmd.AddCommand(pluginCmd)
}
