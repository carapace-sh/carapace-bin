package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show firewall report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	carapace.Gen(showCmd).PositionalCompletion(
		carapace.ActionValues(
			"added",
			"after-rules",
			"before-rules",
			"builtins",
			"listening",
			"logging-rules",
			"raw",
			"user-rules",
		))

	rootCmd.AddCommand(showCmd)
}
