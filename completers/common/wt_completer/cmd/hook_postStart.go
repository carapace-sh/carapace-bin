package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_postStartCmd = &cobra.Command{
	Use:   "post-start",
	Short: "Run post-start hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_postStartCmd).Standalone()

	hook_postStartCmd.Flags().Bool("foreground", false, "Run in foreground (block until complete)")
	hook_postStartCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_postStartCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_postStartCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_postStartCmd)

	// TODO complete names
}
