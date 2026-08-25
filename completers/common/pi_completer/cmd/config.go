package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Open TUI to enable/disable package resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()
	configCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this command")
	configCmd.Flags().BoolP("local", "l", false, "Edit project overrides (.pi/settings.json)")
	configCmd.Flags().Bool("no-approve", false, "Ignore project-local files for this command")
	rootCmd.AddCommand(configCmd)
}