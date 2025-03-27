package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_analyzeProfileCmd = &cobra.Command{
	Use:   "analyze-profile",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_analyzeProfileCmd).Standalone()

	helpCmd.AddCommand(help_analyzeProfileCmd)
}
