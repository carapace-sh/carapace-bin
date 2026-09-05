package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get or set project configuration (meta:* keys)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().BoolP("help", "h", false, "Print help")
	configCmd.Flags().Bool("list", false, "List all config values")
	configCmd.Flags().Bool("unset", false, "Remove a config key")
	rootCmd.AddCommand(configCmd)
}
