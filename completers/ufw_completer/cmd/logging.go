package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loggingCmd = &cobra.Command{
	Use:   "logging",
	Short: "set logging to LEVEL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loggingCmd)

	carapace.Gen(loggingCmd).PositionalCompletion(
		carapace.ActionValues(
			"full",
			"high",
			"low",
			"medium",
			"off",
			"on",
		))

	rootCmd.AddCommand(loggingCmd)
}
