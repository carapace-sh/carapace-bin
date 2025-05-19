package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_renameWorkspaceCmd = &cobra.Command{
	Use:   "rename-workspace [OPTIONS] <NEW_WORKSPACE>",
	Short: "Rename a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_renameWorkspaceCmd).Standalone()

	cli_renameWorkspaceCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cli_renameWorkspaceCmd.Flags().String("pane-id", "", "Specify the current pane")
	cli_renameWorkspaceCmd.Flags().String("workspace", "", "Specify the workspace to rename")
	cliCmd.AddCommand(cli_renameWorkspaceCmd)

	carapace.Gen(cli_renameWorkspaceCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
		// TODO workspace
	})
}
