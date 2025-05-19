package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crashReportCmd = &cobra.Command{
	Use:   "+crash-report",
	Short: "inspect and send crash report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crashReportCmd).Standalone()

	crashReportCmd.Flags().Bool("help", false, "show help")
	rootCmd.AddCommand(crashReportCmd)
}
