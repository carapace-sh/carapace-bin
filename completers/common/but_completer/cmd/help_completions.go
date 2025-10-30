package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Generate shell completion scripts for the specified or inferred shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_completionsCmd).Standalone()

	helpCmd.AddCommand(help_completionsCmd)
}
