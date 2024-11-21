package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "generate autocompletion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autocompleteCmd).Standalone()

	autocompleteCmd.Flags().String("command", "", "")
	autocompleteCmd.Flags().String("format", "", "")
	autocompleteCmd.Flags().String("mode", "", "")
	autocompleteCmd.Flags().String("project", "", "")
	rootCmd.AddCommand(autocompleteCmd)

	// TODO flag completion
	carapace.Gen(autocompleteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("bash", "fish", "zsh"),
	})
}
