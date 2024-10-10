package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_renameCmd = &cobra.Command{
	Use:   "rename [OPTIONS] <NEW_WORKSPACE_NAME>",
	Short: "Renames the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_renameCmd).Standalone()

	workspace_renameCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspaceCmd.AddCommand(workspace_renameCmd)
}
