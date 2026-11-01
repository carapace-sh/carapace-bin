package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_plugins_claude_approveEnterWorktreeCmd = &cobra.Command{
	Use:    "approve-enter-worktree",
	Short:  "Internal: the plugin's PermissionRequest hook, reading its payload from stdin",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_plugins_claude_approveEnterWorktreeCmd).Standalone()

	config_plugins_claude_approveEnterWorktreeCmd.Flags().BoolP("help", "h", false, "Print help")
	config_plugins_claudeCmd.AddCommand(config_plugins_claude_approveEnterWorktreeCmd)
}
