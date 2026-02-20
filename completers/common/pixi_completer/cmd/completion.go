package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates a completion script for a shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().StringP("shell", "s", "", "The shell to generate a completion script for")
	completionCmd.MarkFlagRequired("shell")
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "elvish", "fish", "nushell", "powershell", "zsh"),
	})
}
