package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_gcCmd = &cobra.Command{
	Use:   "gc",
	Short: "Find and optionally delete repo-level config directories whose repo path no longer exists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_gcCmd).Standalone()

	config_gcCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_gcCmd)
}
