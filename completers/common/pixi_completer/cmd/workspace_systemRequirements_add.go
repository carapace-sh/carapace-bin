package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_systemRequirements_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a system requirement to the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_systemRequirements_addCmd).Standalone()

	workspace_systemRequirements_addCmd.Flags().String("family", "", "The family of the system requirement")
	workspace_systemRequirements_addCmd.Flags().StringP("feature", "f", "", "The feature for which the dependency should be modified")
	workspace_systemRequirements_addCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_systemRequirementsCmd.AddCommand(workspace_systemRequirements_addCmd)

	carapace.Gen(workspace_systemRequirements_addCmd).FlagCompletion(carapace.ActionMap{
		"feature":       pixi.ActionFeatures(),
		"manifest-path": carapace.ActionFiles(),
	})
}
