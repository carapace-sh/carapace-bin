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

	drCmd.Flags().BoolP("audit-debug", "D", false, "Enable debugging and profiling of audit methods.")
	drCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	drCmd.Flags().BoolP("help", "h", false, "Show this message.")
	drCmd.Flags().Bool("list-checks", false, "List all audit methods, which can be run individually if provided as arguments.")
	drCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	drCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(drCmd)
}
