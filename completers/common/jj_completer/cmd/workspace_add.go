package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var workspace_addCmd = &cobra.Command{
	Use:   "add [OPTIONS] <DESTINATION>",
	Short: "Add a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_addCmd).Standalone()

	workspace_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspace_addCmd.Flags().String("name", "", "A name for the workspace")
	workspace_addCmd.Flags().StringSliceP("revision", "r", nil, "A list of parent revisions for the working-copy commit of the newly created workspace. You may specify nothing, or any number of parents")
	workspace_addCmd.Flags().String("sparse-patterns", "", "How to handle sparse patterns when creating a new workspace")
	workspaceCmd.AddCommand(workspace_addCmd)

	carapace.Gen(workspace_addCmd).FlagCompletion(carapace.ActionMap{
		"name":     jj.ActionWorkspaces(),
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"sparse-patterns": carapace.ActionValuesDescribed(
			"copy", "Copy all sparse patterns from the current workspace",
			"full", "Include all files in the new workspace",
			"empty", "Clear all files from the workspace (it will be empty)",
		),
	})

	carapace.Gen(workspace_addCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
