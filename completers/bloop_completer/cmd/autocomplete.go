package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bloop"
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
	autocompleteCmd.Flags().String("format", "", "output format")
	autocompleteCmd.Flags().String("mode", "", "completion mode")
	autocompleteCmd.Flags().String("project", "", "")
	rootCmd.AddCommand(autocompleteCmd)

	carapace.Gen(autocompleteCmd).FlagCompletion(carapace.ActionMap{
		"format":  carapace.ActionValues("bash", "fish", "zsh"),
		"mode":    carapace.ActionValues("commands", "mainsfqcn", "projects", "project-commands", "protocols", "reporters", "testsfqcn"),
		"project": bloop.ActionProjects(),
	})
}
