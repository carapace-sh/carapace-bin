package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configClearCmd = &cobra.Command{
	Use:   "config:clear",
	Short: "Remove the configuration cache file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configClearCmd).Standalone()

	rootCmd.AddCommand(configClearCmd)
}
