package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var workspace_removeCmd = &cobra.Command{
	Use:   "remove [OPTIONS] <WORKSPACES>...",
	Short: "Remove a workspace and its working-copy files from disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_removeCmd).Standalone()

	workspace_removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspaceCmd.AddCommand(workspace_removeCmd)

	carapace.Gen(workspace_removeCmd).PositionalAnyCompletion(
		jj.ActionWorkspaces().FilterArgs(),
	)
}
