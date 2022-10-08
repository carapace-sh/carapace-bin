package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion code for the specified shell (bash, zsh, fish, or powershell)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).PositionalCompletion(
		carapace.ActionValues("bash", "zsh"),
	)
}
