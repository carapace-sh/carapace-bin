package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_activation_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List activation scripts and environment variables",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_listCmd).Standalone()

	workspace_activation_listCmd.Flags().StringP("environment", "e", "", "The environment to select or modify")
	workspace_activation_listCmd.Flags().StringP("feature", "f", "", "The feature to select or modify")
	workspace_activation_listCmd.Flags().StringP("target", "t", "", "The target selector, such as linux-64, unix or a glob")
	workspace_activationCmd.AddCommand(workspace_activation_listCmd)

	carapace.Gen(workspace_activation_listCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
		"feature":     pixi.ActionFeatures(),
		"target":      carapace.Batch(pixi.ActionPlatforms(), pixi.ActionKnownPlatforms(), carapace.ActionValues("unix", "linux", "win", "osx")).ToA(),
	})
}
