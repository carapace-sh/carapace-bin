package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateReportCmd = &cobra.Command{
	Use:     "update-report",
	Short:   "The Ruby implementation of `brew update`",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateReportCmd).Standalone()

	updateReportCmd.Flags().Bool("auto-update", false, "Run in 'auto-update' mode (faster, less output).")
	updateReportCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	updateReportCmd.Flags().BoolP("force", "f", false, "Treat installed and updated formulae as if they are from the same taps and migrate them anyway.")
	updateReportCmd.Flags().BoolP("help", "h", false, "Show this message.")
	updateReportCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	updateReportCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateReportCmd)
}
