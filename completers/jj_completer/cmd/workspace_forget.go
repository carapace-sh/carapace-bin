package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var workspace_forgetCmd = &cobra.Command{
	Use:   "forget [OPTIONS] [WORKSPACES]...",
	Short: "Stop tracking a workspace's working-copy commit in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_forgetCmd).Standalone()

	workspace_forgetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	workspaceCmd.AddCommand(workspace_forgetCmd)

	carapace.Gen(workspace_forgetCmd).PositionalAnyCompletion(
		jj.ActionWorkspaces().FilterArgs(),
	)
}
