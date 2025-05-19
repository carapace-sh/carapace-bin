package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reportingCmd = &cobra.Command{
	Use:   "reporting",
	Short: "Commands for reporting data into New Relic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reportingCmd).Standalone()
	rootCmd.AddCommand(reportingCmd)
}
