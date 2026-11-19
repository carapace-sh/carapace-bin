package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion SHELL",
	Short: "Output shell completion for the given shell (bash, fish or zsh)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	applyCommonFlagCompletion(completionCmd)
	rootCmd.AddCommand(completionCmd)
}
