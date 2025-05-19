package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drCmd = &cobra.Command{
	Use:   "dr",
	Short: "Check your system for potential problems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drCmd).Standalone()

	drCmd.Flags().Bool("audit-debug", false, "Enable debugging and profiling of audit methods.")
	drCmd.Flags().Bool("debug", false, "Display any debugging information.")
	drCmd.Flags().Bool("help", false, "Show this message.")
	drCmd.Flags().Bool("list-checks", false, "List all audit methods, which can be run individually if provided as arguments.")
	drCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	drCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(drCmd)
}
