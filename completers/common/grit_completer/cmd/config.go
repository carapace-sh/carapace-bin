package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read, set, or list configuration values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("global", false, "Use the global (per-user) config file instead of this repository's")
	configCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.Flags().BoolP("list", "l", false, "List all configuration values")
	configCmd.Flags().Bool("unset", false, "Remove the key instead of reading or setting it")
	rootCmd.AddCommand(configCmd)
}
