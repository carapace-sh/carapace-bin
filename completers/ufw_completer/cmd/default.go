package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "default",
	Short: "set default policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultCmd).Standalone()

	carapace.Gen(loggingCmd).PositionalCompletion(
		carapace.ActionValues(
			"allow",
			"deny",
			"reject",
		))
	rootCmd.AddCommand(defaultCmd)
}
