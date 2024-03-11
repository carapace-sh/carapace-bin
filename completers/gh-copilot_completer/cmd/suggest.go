package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Suggest a command based on a natural language description of the desired output effect",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suggestCmd).Standalone()

	suggestCmd.Flags().BoolP("help", "h", false, "help for suggest")
	suggestCmd.Flags().StringP("target", "t", "", "Target for suggestion; must be shell, gh, git")
	rootCmd.AddCommand(suggestCmd)

	carapace.Gen(suggestCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionValues("shell", "gh", "git"),
	})
}
