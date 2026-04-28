package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var hook_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a list of hooks which will be run on <hook-name> event",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_listCmd).Standalone()

	hook_listCmd.Flags().Bool("allow-unknown-hook-name", false, "allow running a hook with a non-native hook name")
	hook_listCmd.Flags().Bool("no-allow-unknown-hook-name", false, "do not allow running a hook with a non-native hook name")
	hook_listCmd.Flags().Bool("no-show-scope", false, "do not show the config scope that defined each hook")
	hook_listCmd.Flags().Bool("show-scope", false, "show the config scope that defined each hook")
	hook_listCmd.Flags().BoolS("z", "z", false, "use NUL as line terminator")
	hookCmd.AddCommand(hook_listCmd)

	carapace.Gen(hook_listCmd).PositionalCompletion(
		git.ActionHookEvents(),
	)
}
