package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_createCmd).Standalone()

	workspace_createCmd.Flags().String("cwd", "", "")
	workspace_createCmd.Flags().StringSlice("env", nil, "Set an environment variable for the launched process")
	workspace_createCmd.Flags().Bool("focus", false, "")
	workspace_createCmd.Flags().String("label", "", "")
	workspace_createCmd.Flags().Bool("no-focus", false, "")
	workspaceCmd.AddCommand(workspace_createCmd)

	carapace.Gen(workspace_createCmd).FlagCompletion(carapace.ActionMap{
		"cwd": carapace.ActionFiles(),
	})
}
