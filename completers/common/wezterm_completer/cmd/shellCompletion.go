package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellCompletionCmd = &cobra.Command{
	Use:   "shell-completion --shell <SHELL>",
	Short: "Generate shell completion information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCompletionCmd).Standalone()

	shellCompletionCmd.Flags().BoolP("help", "h", false, "Print help")
	shellCompletionCmd.Flags().String("shell", "", "Which shell to generate for")
	rootCmd.AddCommand(shellCompletionCmd)

	carapace.Gen(shellCompletionCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "elvish", "fish", "power-shell", "zsh", "fig"),
	})
}
