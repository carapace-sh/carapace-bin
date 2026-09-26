package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Display the path of a workspace member",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dirCmd).Standalone()

	workspace_dirCmd.Flags().String("package", "", "Display the path to a specific package in the workspace")
	workspaceCmd.AddCommand(workspace_dirCmd)
}
