package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Run the Ruff formatter on the given files or directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_formatCmd).Standalone()

	helpCmd.AddCommand(help_formatCmd)
}
