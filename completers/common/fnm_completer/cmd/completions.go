package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Print shell completions to stdout",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	completionsCmd.Flags().String("shell", "", "The shell syntax to use. Infers when missing")
	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "zsh", "fish", "powershell"),
	})
}
