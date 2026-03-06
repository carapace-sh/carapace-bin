package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_preSwitchCmd = &cobra.Command{
	Use:   "pre-switch",
	Short: "Run pre-switch hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_preSwitchCmd).Standalone()

	hook_preSwitchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_preSwitchCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_preSwitchCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_preSwitchCmd)

	// TODO complete names
}
