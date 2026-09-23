package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bugReportCmd = &cobra.Command{
	Use:   "bug-report",
	Short: "Request bug report generation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugReportCmd).Standalone()

	bugReportCmd.Flags().Bool("progress", false, "launch a notification right away to show its progress")
	bugReportCmd.Flags().Bool("telephony", false, "dump only telephony sections")

	rootCmd.AddCommand(bugReportCmd)

}
