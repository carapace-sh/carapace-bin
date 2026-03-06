package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_postRemoveCmd = &cobra.Command{
	Use:   "post-remove",
	Short: "Run post-remove hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_postRemoveCmd).Standalone()

	hook_postRemoveCmd.Flags().Bool("foreground", false, "Run in foreground (block until complete)")
	hook_postRemoveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_postRemoveCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_postRemoveCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_postRemoveCmd)

	// TODO complete names
}
