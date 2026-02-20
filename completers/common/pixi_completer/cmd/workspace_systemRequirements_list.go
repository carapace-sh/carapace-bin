package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var workspace_systemRequirements_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the environments in the manifest file",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_systemRequirements_listCmd).Standalone()

	workspace_systemRequirements_listCmd.Flags().StringP("environment", "e", "", "The environment to list the system requirements for")
	workspace_systemRequirements_listCmd.Flags().Bool("json", false, "List the system requirements in JSON format")
	workspace_systemRequirementsCmd.AddCommand(workspace_systemRequirements_listCmd)

	carapace.Gen(workspace_systemRequirements_listCmd).FlagCompletion(carapace.ActionMap{
		"environment": pixi.ActionEnvironments(),
	})
}
