package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_listCmd).Standalone()

	workspace_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspaceCmd.AddCommand(workspace_listCmd)
}
