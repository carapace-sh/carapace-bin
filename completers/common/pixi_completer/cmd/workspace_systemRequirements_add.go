package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_systemRequirements_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds an environment to the manifest file",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_systemRequirements_addCmd).Standalone()

	workspace_systemRequirements_addCmd.Flags().String("family", "", "The Libc family, this can only be specified for requirement `other-libc`")
	workspace_systemRequirements_addCmd.Flags().StringP("feature", "f", "", "The name of the feature to modify")
	workspace_systemRequirementsCmd.AddCommand(workspace_systemRequirements_addCmd)

	carapace.Gen(workspace_systemRequirements_addCmd).FlagCompletion(carapace.ActionMap{
		"feature": pixi.ActionFeatures(),
	})
}
