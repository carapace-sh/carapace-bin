package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var analyze_help_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generate a map of Python file dependencies or dependents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyze_help_graphCmd).Standalone()

	analyze_helpCmd.AddCommand(analyze_help_graphCmd)
}
