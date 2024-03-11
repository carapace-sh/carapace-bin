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
	updateReportCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updateReportCmd.Flags().Bool("force", false, "Treat installed and updated formulae as if they are from the same taps and migrate them anyway.")
	updateReportCmd.Flags().Bool("help", false, "Show this message.")
	updateReportCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updateReportCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(updateReportCmd)
}
