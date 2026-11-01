package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Short: "Manage runtimes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_runtimeCmd).Standalone()

	helpCmd.AddCommand(help_runtimeCmd)
}
