package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()
	completionCmd.Flags().String("shell", "", "Output completion for the specified shell.  (bash, powershell, zsh)")
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "powershell", "zsh"),
	})
}
