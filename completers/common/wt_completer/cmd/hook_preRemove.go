package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_preRemoveCmd = &cobra.Command{
	Use:   "pre-remove",
	Short: "Run pre-remove hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_preRemoveCmd).Standalone()

	hook_preRemoveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_preRemoveCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_preRemoveCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_preRemoveCmd)

	// TODO complete names
}
