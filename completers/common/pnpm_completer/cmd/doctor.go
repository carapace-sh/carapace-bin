package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run diagnostics on the pnpm installation and environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().Bool("benchmark", false, "Also time filesystem and install operations")
	doctorCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	doctorCmd.Flags().Bool("json", false, "Report the results as JSON")
	doctorCmd.Flags().Bool("offline", false, "Skip checks that need network access")
	rootCmd.AddCommand(doctorCmd)
}
