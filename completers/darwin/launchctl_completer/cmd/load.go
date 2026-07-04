package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a service for execution (deprecated, use bootstrap)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadCmd).Standalone()
	loadCmd.Flags().BoolP("all", "a", false, "Load all services in the directory")
	loadCmd.Flags().BoolP("write", "w", false, "Override the Disabled key and persist the change")
	rootCmd.AddCommand(loadCmd)
	carapace.Gen(loadCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
