package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_analyze_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generate a map of Python file dependencies or dependents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_analyze_graphCmd).Standalone()

	help_analyzeCmd.AddCommand(help_analyze_graphCmd)
}
