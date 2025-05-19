package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate and display various kinds of reports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reportCmd).Standalone()

	reportCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(reportCmd)
}
