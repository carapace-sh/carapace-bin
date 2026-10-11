package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pixi_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_activation_env_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove activation environment variables",
	Aliases: []string{"rm", "unset"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_env_removeCmd).Standalone()

	workspace_activation_env_removeCmd.Flags().StringP("environment", "e", "", "The environment to select or modify")
	workspace_activation_env_removeCmd.Flags().StringP("feature", "f", "", "The feature to select or modify")
	workspace_activation_env_removeCmd.Flags().StringP("target", "t", "", "The target selector, such as linux-64, unix or a glob")
	workspace_activation_envCmd.AddCommand(workspace_activation_env_removeCmd)

	carapace.Gen(workspace_activation_env_removeCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"feature":     pixi.ActionFeatures(),
		"target":      carapace.Batch(pixi.ActionPlatforms(), pixi.ActionKnownPlatforms(), carapace.ActionValues("unix", "linux", "win", "osx")).ToA(),
	})

	carapace.Gen(workspace_activation_env_removeCmd).PositionalAnyCompletion(
		action.ActionActivationVariables(workspace_activation_env_removeCmd),
	)
}
