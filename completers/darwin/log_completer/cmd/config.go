package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View/change logging system settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().String("category", "", "Get/set settings for a given category")
	configCmd.Flags().String("mode", "", "Enable given modes")
	configCmd.Flags().String("process", "", "Get/set settings for a given process")
	configCmd.Flags().Bool("reset", false, "Reset settings to defaults")
	configCmd.Flags().Bool("status", false, "Show current settings")
	configCmd.Flags().String("subsystem", "", "Get/set settings for a given subsystem")
	rootCmd.AddCommand(configCmd)
}
