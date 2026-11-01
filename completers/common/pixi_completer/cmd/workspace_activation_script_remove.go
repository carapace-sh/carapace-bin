package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_activation_script_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove activation scripts from the list",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_script_removeCmd).Standalone()

	workspace_activation_script_removeCmd.Flags().StringP("environment", "e", "", "The environment to select or modify")
	workspace_activation_script_removeCmd.Flags().StringP("feature", "f", "", "The feature to select or modify")
	workspace_activation_script_removeCmd.Flags().StringP("target", "t", "", "The target selector, such as linux-64, unix or a glob")
	workspace_activation_scriptCmd.AddCommand(workspace_activation_script_removeCmd)

	carapace.Gen(workspace_activation_script_removeCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"feature":     pixi.ActionFeatures(),
		"target":      carapace.Batch(pixi.ActionPlatforms(), pixi.ActionKnownPlatforms(), carapace.ActionValues("linux", "osx", "unix", "win")).ToA(),
	})

	carapace.Gen(workspace_activation_script_removeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
