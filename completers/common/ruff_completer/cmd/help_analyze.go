package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Run analysis over Python source code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_analyzeCmd).Standalone()

	helpCmd.AddCommand(help_analyzeCmd)
}
