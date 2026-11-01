package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batteryreportCmd = &cobra.Command{
	Use:   "batteryreport",
	Short: "generate a report of battery usage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batteryreportCmd).Standalone()
	rootCmd.AddCommand(batteryreportCmd)
}
