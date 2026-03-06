package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_postSwitchCmd = &cobra.Command{
	Use:   "post-switch",
	Short: "Run post-switch hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_postSwitchCmd).Standalone()

	hook_postSwitchCmd.Flags().Bool("foreground", false, "Run in foreground (block until complete)")
	hook_postSwitchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_postSwitchCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_postSwitchCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_postSwitchCmd)

	// TODO complete names
}
