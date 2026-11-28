package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configReportCmd = &cobra.Command{
	Use:   "config-report",
	Short: "Generate a JSON file with the project configuration report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configReportCmd).Standalone()
	rootCmd.AddCommand(configReportCmd)
}
