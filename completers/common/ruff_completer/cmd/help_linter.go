package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_linterCmd = &cobra.Command{
	Use:   "linter",
	Short: "List all supported upstream linters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_linterCmd).Standalone()

	helpCmd.AddCommand(help_linterCmd)
}
