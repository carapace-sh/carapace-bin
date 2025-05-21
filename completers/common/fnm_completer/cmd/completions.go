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
	rootCmd.AddCommand(completionsCmd)

	addCommonFlags(completionsCmd)

	completionsCmd.Flags().String("shell", "", "The shell syntax to use. Infers when missing")

	carapace.Gen(completionsCmd).Standalone()

	carapace.Gen(completionsCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "zsh", "fish", "powershell"),
	})
}
