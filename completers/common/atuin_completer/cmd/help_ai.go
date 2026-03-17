package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Run the AI assistant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_aiCmd).Standalone()

	helpCmd.AddCommand(help_aiCmd)
}
