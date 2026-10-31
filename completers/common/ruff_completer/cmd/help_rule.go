package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Explain a rule (or all rules)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ruleCmd).Standalone()

	helpCmd.AddCommand(help_ruleCmd)
}
