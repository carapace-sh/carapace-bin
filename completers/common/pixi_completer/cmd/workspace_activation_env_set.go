package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_activation_env_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set activation environment variables",
	Aliases: []string{"add"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_env_setCmd).Standalone()

	workspace_activation_env_setCmd.Flags().StringP("environment", "e", "", "The environment to select or modify")
	workspace_activation_env_setCmd.Flags().StringP("feature", "f", "", "The feature to select or modify")
	workspace_activation_env_setCmd.Flags().StringP("target", "t", "", "The target selector, such as linux-64, unix or a glob")
	workspace_activation_envCmd.AddCommand(workspace_activation_env_setCmd)

	carapace.Gen(workspace_activation_env_setCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"feature":     pixi.ActionFeatures(),
		"target":      carapace.Batch(pixi.ActionPlatforms(), pixi.ActionKnownPlatforms(), carapace.ActionValues("linux", "osx", "unix", "win")).ToA(),
	})

	carapace.Gen(workspace_activation_env_setCmd).PositionalAnyCompletion(
		carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return os.ActionEnvironmentVariables().Suffix("=")
			}
			return carapace.ActionValues()
		}),
	)
}
