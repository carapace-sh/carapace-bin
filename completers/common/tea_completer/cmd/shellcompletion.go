package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellcompletionCmd = &cobra.Command{
	Use:     "shellcompletion",
	Short:   "Install shell completion for tea",
	GroupID: "SETUP",
	Aliases: []string{"autocomplete"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellcompletionCmd).Standalone()

	shellcompletionCmd.Flags().Bool("install", false, "Persist in shell config instead of printing commands")
	rootCmd.AddCommand(shellcompletionCmd)

	carapace.Gen(shellcompletionCmd).PositionalCompletion(
		carapace.ActionValues("bash", "zsh", "powershell", "fish"),
	)
}
