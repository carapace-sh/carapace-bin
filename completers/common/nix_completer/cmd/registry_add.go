package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registry_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add/replace flake in user flake registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_addCmd).Standalone()

	registry_addCmd.Flags().String("registry", "", "The registry to operate on")
	registryCmd.AddCommand(registry_addCmd)

	addEvaluationFlags(registry_addCmd)
	addLoggingFlags(registry_addCmd)

	carapace.Gen(registry_addCmd).FlagCompletion(carapace.ActionMap{
		"registry": carapace.ActionFiles(),
	})
	// TODO positional completion
}
