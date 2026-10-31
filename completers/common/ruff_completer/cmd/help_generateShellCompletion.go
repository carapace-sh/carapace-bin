package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_generateShellCompletionCmd = &cobra.Command{
	Use:    "generate-shell-completion",
	Short:  "Generate shell completion",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_generateShellCompletionCmd).Standalone()

	helpCmd.AddCommand(help_generateShellCompletionCmd)
}
