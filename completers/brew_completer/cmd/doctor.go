package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:     "doctor",
	Short:   "Check your system for potential problems",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().Bool("audit-debug", false, "Enable debugging and profiling of audit methods.")
	doctorCmd.Flags().Bool("debug", false, "Display any debugging information.")
	doctorCmd.Flags().Bool("help", false, "Show this message.")
	doctorCmd.Flags().Bool("list-checks", false, "List all audit methods, which can be run individually if provided as arguments.")
	doctorCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	doctorCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(doctorCmd)
}
