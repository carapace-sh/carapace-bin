package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Generate tab-completion scripts for your shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_completionsCmd).Standalone()

	helpCmd.AddCommand(help_completionsCmd)
}
