package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Shell tab-completion scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).PositionalCompletion(
		carapace.ActionValues("bash", "fish", "powershell", "zsh"),
	)
}
