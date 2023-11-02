package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_addCmd).Standalone()

	workspace_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspace_addCmd.Flags().String("name", "", "A name for the workspace")
	workspace_addCmd.Flags().StringP("revision", "r", "", "The revision that the workspace should be created at; a new working copy commit will be created on top of it")
	workspaceCmd.AddCommand(workspace_addCmd)
}
