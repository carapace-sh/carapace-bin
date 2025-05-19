package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Show Homebrew and system configuration info useful for debugging",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("debug", false, "Display any debugging information.")
	configCmd.Flags().Bool("help", false, "Show this message.")
	configCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	configCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(configCmd)
}
